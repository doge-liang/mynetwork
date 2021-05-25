from __future__ import (absolute_import, division, print_function, unicode_literals)

# 导入 backtrader
import backtrader as bt
import pandas as pd
import json
import heapq
# 导入其他包
from datetime import datetime
from datetime import timedelta

from app.data_source.api.tushare_api import Stock
from tqdm import tqdm


# 创建策策略
class MA(bt.Strategy):
    # 自定义一些参数
    params = (
        ('pslow', 30),
        ('pfast', 10),
        ('printlog', False),
        ('stake', 100),
    )

    def log(self, txt, dt=None, doprint=False):
        """ Logging function fot this strategy"""
        if self.p.printlog or doprint:
            dt = dt or self.datas[0].datetime.date(0)
            print('%s, %s' % (dt.isoformat(), txt))

    def __init__(self):
        # 从 self.datas 中能访问到 cerebro.adddata(data) 中的数据
        # self.datas[0] 即是加载的第一条价格数据，它被框架默认使用。
        # 引用 data [0] 数据序列中“收盘价”行
        self.dataclose = self.datas[0].close
        self.dataopen = self.datas[0].open
        self.bar_executed = len(self)
        # 跟踪挂单
        self.order = None
        self.buyprice = None
        self.buycomm = None
        self.pool = []

        # 添加移动平均线指标
        self.sma_fast = {stock: bt.ind.SMA(stock.close, period=self.p.pfast) for stock in self.datas}
        self.sma_slow = {stock: bt.ind.SMA(stock.close, period=self.p.pslow) for stock in self.datas}
        # 长线 - 短线
        self.sma_diff = {stock: self.sma_slow[stock] - self.sma_fast[stock] for stock in self.datas}
        self.orderlist = []

    def notify_order(self, order):
        for order in self.orderlist:
            if order.status in [order.Submitted, order.Accepted]:
                # 购买订单已提交给经纪人/接受经纪人的卖单 - 啥也不干
                return

            # 检查订单是否已完成
            # 注意：如果现金不足，经纪人可能会拒绝订单
            if order.status in [order.Completed]:
                if order.isbuy():
                    self.log('BUY EXECUTED, %.2f, Cost: %.2f, Comm: %.2f' %
                             (order.executed.price,
                              order.executed.value,
                              order.executed.comm))

                    self.buyprice = order.executed.price
                    self.buycomm = order.executed.comm
                elif order.issell():
                    self.log('SELL EXECUTED, %.2f, Cost: %.2f, Comm: %.2f' %
                             (order.executed.price,
                              order.executed.value,
                              order.executed.comm))

                self.bar_executed = len(self)

            elif order.status in [order.Canceled, order.Margin, order.Rejected]:
                self.log('Order Canceled/Margin/Rejected')

            # 订单终止
            self.orderlist.remove(order)

    def notify_trade(self, trade):
        if not trade.isclosed:
            return

        self.log('OPERATION PROFIT, GROSS %.2f, NET %.2f' %
                 (trade.pnl, trade.pnlcomm))

    # 当经过一个K线柱的时候 next() 方法就会被调用一次。
    def next(self):
        # Simply log the closing price of the series from the reference
        # 只需从引用中打印该收盘价序列
        self.log('Open: %.2f, Close, %.2f' % (self.dataopen[0], self.dataclose[0]))
        for o in self.orderlist:
            self.cancel(o)  # 取消以往所有订单
            self.orderlist = []  # 置空

        if len(self.pool) == 0:
            self.pool = heapq.nsmallest(50, self.sma_diff, key=lambda d: self.sma_diff.get(d)[0])
            print("初始股票集：")
            print(self.pool)

        for stock in self.pool:
            # 短线上穿长线
            buy_signal = self.sma_diff[stock][0] < 0 and self.sma_diff[stock][-1] > 0
            # 短线下穿长线
            sell_signal = self.sma_diff[stock][0] > 0 and self.sma_diff[stock][-1] < 0
            if not self.getposition(stock):
                if buy_signal:
                    self.log('BUY CREATE, %.2f' % self.dataclose[0])
                    self.orderlist.append(self.buy(stock, size=self.p.stake))
            else:
                if sell_signal:
                    self.log('SELL CREATE, %.2f' % self.dataclose[0])
                    self.orderlist.append(self.sell(stock, size=self.p.stake))

    def stop(self):
        print('==================== Results ====================')
        print('Starting Value - %.2f' % self.broker.startingcash)
        print('Ending   Value - %.2f' % self.broker.getvalue())
        print('=================================================\n\n\n')


def run_strategy():
    pd.set_option('display.max_columns', None)
    pd.set_option('display.width', 5000)

    # 实例化 Cerebro 引擎
    cerebro = bt.Cerebro(tradehistory=True)

    # 添加一个策略
    _ = cerebro.addstrategy(MA)

    # 参数调优
    # strats = cerebro.optstrategy(
    #     TestStrategy,
    #     maperiod=range(10, 31))

    publish_date = datetime(2020, 5, 11)
    now = datetime(2021, 5, 16)
    before = (now - publish_date).days
    data_api = Stock(now, before, 0)
    for code in tqdm(data_api.getHS300()):
        try:
            ohlc_data = data_api.getDailyKVLocal(code)
        except Exception:
            ohlc_data = data_api.getDailyKVOnline(code)

        #         flag = False
        #         if flag:
        #             print("以下股票是近一年上市的新股，数据不足：")
        if (datetime.now() - ohlc_data.index[0]) < data_api.before:
            #             flag = True
            #             print(code)
            continue

        stock = bt.feeds.PandasData(dataname=ohlc_data, nocase=True)
        # 将数据添加到 Cerebro
        _ = cerebro.adddata(stock, name=code)

    print(len(cerebro.datas))
    # 设定我们想要的初始金额
    cerebro.broker.setcash(100000.0)
    # 设定手续费，国内一般是是0.0003
    cerebro.broker.setcommission(0.003)
    # 添加分析器，用于计算各种指标
    cerebro.addanalyzer(bt.analyzers.SharpeRatio, _name="SharpeRatio")
    cerebro.addanalyzer(bt.analyzers.AnnualReturn, _name="AannualReturn")
    cerebro.addanalyzer(bt.analyzers.Returns, _name="Returns")
    cerebro.addanalyzer(bt.analyzers.DrawDown, _name="DrawDown")
    cerebro.addanalyzer(bt.analyzers.TradeAnalyzer, _name="Trade")
    cerebro.addanalyzer(bt.analyzers.PyFolio)
    cerebro.addanalyzer(bt.analyzers.PositionsValue, _name="Positions", headers=True, cash=True)
    strats = cerebro.run()
    strat = strats[0]

    # 获取 pyfolio 分析器内的值，_ 用于丢弃不用的变量
    pyfolio = strat.analyzers.getbyname('pyfolio')
    returns, _, transactions, _ = pyfolio.get_pf_items()

    sharpe_ratio = round(strat.analyzers.SharpeRatio.get_analysis()['sharperatio'], 2)
    max_drawdown = round(strat.analyzers.DrawDown.get_analysis()['max']['drawdown'], 2)
    returns = round(strat.analyzers.Returns.get_analysis()['rnorm100'], 2)
    positions = strat.analyzers.Positions.get_analysis()

    transactions['commission'] = round((abs(transactions['value']) * 0.003), 2)
    #     transactions.index = pd.DatetimeIndex(transactions.index.date)
    planning_trades = [{"stockID": o.data._name, "amount": o.size} for o in strat.orderlist]

    print('============================================================================================')
    print('====================================== Performance =========================================')
    print('============================================================================================')
    print('Sharpe Ratio:', sharpe_ratio)
    print('Max DrawDown:', max_drawdown, '%')

    print('===========================================================================================\n\n\n')

    print('===========================================================================================')
    print("========================================= returns =========================================")
    print('===========================================================================================')
    print('Aannual Return:', returns, "%")
    print('===========================================================================================\n\n\n')

    print('===========================================================================================')
    print("================================== positionsValue =========================================")
    print('===========================================================================================')
    ps = [[k] + v for k, v in iter(positions.items())]
    header = [ps[0][i] for i in range(len(ps[0]))]
    positions = pd.DataFrame(ps[2:], columns=header)
    positions.index = positions['Datetime']
    positions = positions[positions.columns[2:]]
    no_zero_col = positions.columns[(positions[-1:] != 0).any()]
    positions = positions[-1:][no_zero_col]
    positions_dict = []
    for _, i in positions.T.items():
        for k in i.to_dict():
            positions_dict.append({"stockID": k, "amount": i.to_dict()[k]})
    print(positions_dict)
    print('===========================================================================================\n\n\n')

    print('===========================================================================================')
    print("==================================== transactions =========================================")
    print('===========================================================================================')
    transactions = transactions.reset_index()
    transactions = transactions[['date', 'amount', 'price', 'symbol', 'commission']]
    # [['date', 'amount', 'price', 'symbol', 'commision']]
    new_cols = ['dateTime', 'amount', 'price', 'stockID', 'commission']
    transactions.columns = new_cols
    transactions['dateTime'] = transactions['dateTime'].map(lambda x: x.astimezone("+08:00:00").isoformat())
    transactions['id'] = transactions.index + 1
    transactions = json.loads(transactions.to_json(orient="records", force_ascii=False))
    print(transactions)
    print('===========================================================================================\n\n\n')

    return sharpe_ratio, max_drawdown, returns, transactions, positions_dict, planning_trades


if __name__ == '__main__':
    run_strategy()
    # Stock(datetime.now(), 365, 0).updateLocal()
