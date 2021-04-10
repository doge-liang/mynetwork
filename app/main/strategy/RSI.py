import pandas as pd
# import numpy as np
from datetime import datetime
from datetime import timedelta
from tqdm import tqdm
import time
# import matplotlib.pyplot as plt
import talib as tb
import pyfolio as pf
import tushare as ts

pro = ts.pro_api('995808334c794fadeb95ef31b35586073353869506a51b867a4a6305')

##################
#    RSI 策略    #
##################

####
# 短期RSI是指参数相对小的RSI，长期RSI是指参数相对较长的RSI。比如，6日RSI和12日RSI中 ，6日RSI即为短期RSI，12日RSI即为长期RSI。
# 长短期RSI线的交叉情况可以作为我们研判行情的方法。
# 1、当短期RSI>长期RSI时，市场则属于多头市场；
# 2、当短期RSI<长期RSI时，市场则属于空头市场；
# 3、当短期RSI线在低位向上突破长期RSI线,则是市场的买入信号；
# 4、当短期RSI线在高位向下突破长期RSI线，则是市场的卖出信号。
####

###
# 从2021年1月11日开始，周日晚上选出当前沪深300成分股中（短期RSI-长期RSI）最大的10只股票并持有一周。
# 在该周当中，每天晚上运行策略查看这10只持有股票的RSI。若短期RSI<长期RSI,则卖出持仓。
###


# 获取股票日K线数据
def stock_value_get(codes, sdate, edate):
    # 存储股票数据
    stock_data = pd.DataFrame()
    for code in tqdm(codes):
        stock_data = stock_data.append(ts.pro_bar(ts_code=code, adj='qfq', start_date=sdate, end_date=edate))
        time.sleep(0.1)
    return stock_data


# 获取下一周的股票池
def stock_pool_get(data):
    # 记录每只股票的数据
    RSI_data = {}
    # 计算每只股票的 短期RSI-长期RSI
    for stock, value in data.groupby('ts_code'):
        RSI_data[stock] = list(tb.RSI(value['close'], timeperiod=6) - tb.RSI(value['close'], timeperiod=12))[-1]
    # 从大到小排序并选出最大的10个
    return pd.DataFrame(sorted(RSI_data.items(), key=lambda x: x[1], reverse=True)[:10],
                        columns=['stock_code', 'RSI_value'])


# 计算个股每日收益率
def cal_return_daily(df):
    df['return'] = 0
    # 用于记录最后一次买入
    end = 0
    long = list(df[df['position'] == 1].index)
    short = list(df[df['position'] == -1].index)

    for s, e in zip(long, short):
        if e - s < 4:
            df.loc[s:e, 'return'] = (
                    (df.loc[s:e, 'open'] - df.loc[s:e, 'open'].shift(1)) / df.loc[s:e, 'open'].shift(1)).shift(-1)
        else:
            df.loc[s:e, 'return'] = (
                    (df.loc[s:e, 'open'] - df.loc[s:e, 'open'].shift(1)) / df.loc[s:e, 'open'].shift(1)).shift(-1)
            df.loc[e, 'return'] = (df.loc[e, 'close'] - df.loc[e, 'open']) / df.loc[e, 'open']

    return df


# 计算投资组合的每日收益
def cal_portfolio_daily_return(daily_return):
    portfolio_daily_return = {}
    for date, value in daily_return.groupby('trade_date'):
        portfolio_daily_return[date] = sum(value['return']) / 10

    return portfolio_daily_return


if __name__ == '__main__':

    # 记录持仓
    my_position = {}
    # 日期格式
    time_str = '%Y%m%d'
    # 假设当前的时间节点为2020年7月5日，即策略执行的前一天
    now = datetime(2020, 7, 5)

    # 执行2周
    week = 2
    for w in range(week):
        # 获取沪深300当前所有成分股代码
        codes = pro.index_weight(index_code='399300.SZ',
                                 start_date=(now - timedelta(days=180)).strftime(time_str),
                                 end_date=now.strftime(time_str))[:300]['con_code']
        # 获取沪深300当前所有成分股30天内的数据
        stock_data = stock_value_get(codes, sdate=(now - timedelta(days=30)).strftime(time_str),
                                     edate=now.strftime(time_str))
        stock_pool = stock_pool_get(stock_data)
        # 记录当前pool
        pool_code = list(stock_pool['stock_code'])
        # 记录开始时间和卖出时间
        open_close = pd.DataFrame(index=stock_pool['stock_code'], columns=['open', 'close'])
        # 每周第一天以开盘价买入
        open_close['open'] = (now + timedelta(1)).strftime(time_str)
        # 未来5天，每天更新数据，重新计算这10只股票的RSI
        for i in range(5):
            # 模拟增加一天
            now = now + timedelta(1)
            # 获取股票数据
            stock_pool_data = stock_value_get(pool_code, sdate=(now - timedelta(days=30)).strftime(time_str),
                                              edate=now.strftime(time_str))
            # 判断RSI是否小于0
            for stock, value in stock_pool_data.groupby('ts_code'):
                if i == 4:
                    # 每周最后一天以当天收盘价全部清仓
                    open_close.loc[stock, 'close'] = (now).strftime(time_str)
                else:
                    if list(tb.RSI(value['close'], timeperiod=6) - tb.RSI(value['close'], timeperiod=12))[-1] < 0:
                        # 第二天以开盘价卖出
                        open_close.loc[stock, 'close'] = (now + timedelta(1)).strftime(time_str)
                        pool_code.remove(stock)

        # 保存开始和结束时间
        for code in list(open_close.index):
            my_position.setdefault(code, []).append(list(open_close.loc[code, :].values))
        # 过周末
        now = now + timedelta(days=2)
        print('第{}周结束'.format(w + 1))



    # 获取持仓股票在期间内的数据
    stock_buy_data = stock_value_get(list(my_position.keys()),
                                     sdate=datetime(2020, 7, 5).strftime(time_str),
                                     edate=now.strftime(time_str))
    stock_buy_data.reset_index(drop=True, inplace=True)
    # 保存持仓，1为开仓，-1为空仓
    stock_buy_data['position'] = 0
    # 将position转换为1和-1
    for (stock, value) in my_position.items():
        buy = []
        sell = []
        for l in value:
            buy.append(
                list(stock_buy_data[(stock_buy_data['ts_code'] == stock) & (stock_buy_data['trade_date'] == l[0])].index)[
                    0])
            sell.append(
                list(stock_buy_data[(stock_buy_data['ts_code'] == stock) & (stock_buy_data['trade_date'] == l[1])].index)[
                    0])
        stock_buy_data.iloc[buy, -1] = 1
        stock_buy_data.iloc[sell, -1] = -1
    # 记录各股票每日的回报率
    daily_return = pd.DataFrame()
    for stock, value in stock_buy_data.groupby('ts_code'):
        value = value[::-1].copy()
        value.reset_index(drop=True, inplace=True)
        daily_return = daily_return.append(cal_return_daily(value))
    daily_return['return'] = daily_return['return'].fillna(0)
    # 记录投资组合每日的而回报率
    portfolio_daily_return = pd.Series(cal_portfolio_daily_return(daily_return))
    portfolio_daily_return.index = pd.to_datetime(portfolio_daily_return.index)
    # 获取基准（沪深300）数据
    benchmark_data = pro.index_daily(ts_code='399300.SZ',
                                     start_date=datetime(2020, 7, 5).strftime(time_str),
                                     end_date=now.strftime(time_str))
    benchmark_data = benchmark_data[::-1].copy()
    # 计算基准收益
    benchmark_daily_return = ((benchmark_data['open'] - benchmark_data['open'].shift(1))
                              / benchmark_data['open'].shift(1)).shift(-1)
    benchmark_daily_return[list(benchmark_daily_return.index)[-1]] = \
        (list(benchmark_data['close'])[-1] - list(benchmark_data['open'])[-1]) / list(benchmark_data['open'])[-1]
    benchmark_daily_return.index = portfolio_daily_return.index
    pf.create_returns_tear_sheet(portfolio_daily_return, benchmark_rets=benchmark_daily_return)
