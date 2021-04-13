import tushare as ts
from datetime import datetime
from datetime import timedelta
import time
import talib as tb
import pandas as pd
from tqdm import tqdm
from pathlib import Path


class Stock():
    TIME_STR = '%Y%m%d'
    # 股票数据缓存
    stock_datas_cache = {}

    def __init__(self, now, before=0, after=0, token=None):
        token = token or '61d84738dc9144d5570ced6f00f67ce1ec2af4a282ac02938ca416d4'
        print(token)
        self.pro = ts.pro_api(token)
        self.data_local_path = Path.cwd() / '..' / 'data_source' / 'local'
        self.local_data_list = [path.stem for path in self.data_local_path.iterdir() if path.is_file()]
        if type(now) == str:
            try:
                self.now = self.str2date(now)
            except Exception as e:
                print(e)

        elif type(now) == datetime:
            self.now = now
        else:
            raise Exception("now 格式不对")
        self.before = timedelta(days=before)
        self.after = timedelta(days=after)

    def getHS300(self):
        # 获取沪深300当前所有成分股代码
        codes = self.pro.index_weight(index_code='399300.SZ',
                                      start_date=(self.now - self.before + self.after).strftime(self.TIME_STR),
                                      end_date=self.now.strftime(self.TIME_STR))[:300]['con_code']
        # print(list(codes))
        return list(codes)

    def getDailyKV(self, code, from_date, to_date):
        # 如果缓存了就不用下载
        # print(self.local_data_list)
        # print(self.stock_datas_cache.keys())
        if code in self.stock_datas_cache.keys():
            print("找到缓存")
            return self.stock_datas_cache[code]
        elif code in self.local_data_list:
            # print("找到本地文件")
            stock_data = pd.read_csv(self.data_local_path / (code + '.csv'),
                                     # parse_dates=[0],
                                     parse_dates=True,
                                     index_col=0,
                                     )
            # print(stock_data)
            # stock_data.index.name = 'datetime'
            # stock_data.fillna(0)
            # self.save_csv(stock_data, code)
        else:
            stock_data = ts.pro_bar(ts_code=code,
                                    api=self.pro,
                                    adj='qfq',
                                    start_date=from_date,
                                    end_date=to_date)
            stock_data.index = pd.to_datetime(stock_data['trade_date'])
            stock_data.index.name = 'datetime'
            stock_data = stock_data[['open', 'high', 'low', 'close', 'vol']]
            stock_data['openinterest'] = 0
            stock_data.fillna(0)
            time.sleep(1)
            self.stock_datas_cache['code'] = stock_data

            self.save_csv(stock_data, code)
        return stock_data

    def getStockPoolDailyKV(self, codes, from_date, to_date):
        stock_pool = {code: self.getDailyKV(code, from_date, to_date) for code in codes}

        return stock_pool

    def selectStockPoolByRSI(self, stock_codes=None, pool_size=None):
        # 获取备选股票集的数据，若没有就用沪深300
        stock_codes = stock_codes or self.getHS300()
        # 如果传值为空则股票池大小为10
        pool_size = pool_size or 10
        # 如果股票池大小比备选股票集大小大，那么用备选股票集的大小
        pool_size = len(stock_codes) if pool_size > len(stock_codes) else pool_size

        # 计算每支备选股的RSI
        stocks_RSI = {}
        for code in tqdm(stock_codes):
            stocks_RSI[code] = list(
                tb.RSI(self.getDailyKV(code,
                                       self.date2str(self.now - self.before),
                                       self.now.strftime(self.TIME_STR)
                                       )['close'], timeperiod=6)
                -
                tb.RSI(self.getDailyKV(code,
                                       self.date2str(self.now - self.before),
                                       self.now.strftime(self.TIME_STR)
                                       )['close'], timeperiod=12)
            )[-1]

        stock_pool = [selected[0] for selected in
                      sorted(stocks_RSI.items(), key=lambda x: x[1], reverse=True)[:pool_size]]
        return stock_pool

    def getStockName(self, code):
        return self.pro.stock_basic(code)['name']

    def save_csv(self, data: pd.DataFrame, name: str):
        cwd = Path.cwd()
        data.to_csv(cwd / self.data_local_path / (name + ".csv"))

    def date2str(self, date: datetime):
        return date.strftime(self.TIME_STR)

    def str2date(self, date_str: str):
        return datetime.strptime(date_str, self.TIME_STR)
