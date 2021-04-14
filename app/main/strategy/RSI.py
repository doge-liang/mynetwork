from __future__ import (absolute_import, division, print_function, unicode_literals)

# 导入 backtrader
import backtrader as bt

# 导入其他包
from datetime import datetime
# import os.path
# import sys  # 找出脚本名称 (in argv[0])
from app.data_source.api.tushare_api import Stock
from tqdm import tqdm
# import matplotlib.pylab as plt
# import pandas as pd


# 创建策策略
class TestStrategy(bt.Strategy):
    # 自定义一些参数
    params = (
        ('maperiod', 25),
        ('printlog', False),
        ('stake', 1000),
    )

    def log(self, txt, dt=None, doprint=False):
        """ Logging function fot this strategy"""
        if self.p.printlog or doprint:
            dt = dt or self.datas[0].datetime.date(0)
            print('%s, %s' % (dt.isoformat(), txt))

    def __init__(self):
        # 从 self.data_source 中能访问到 cerebro.adddata(data) 中的数据
        # self.data_source[0] 即是加载的第一条价格数据，它被框架默认使用。
        # 引用 data [0] 数据序列中“收盘价”行
        self.bar_executed = len(self)
        self.mas = {stock: bt.ind.MovingAverageSimple(stock.close, period=self.p.maperiod) for stock in self.datas}
        # 跟踪挂单
        self.orderlist = []
        self.buyprice = None
        self.comms = []

    def notify_order(self, order):
        # print(order.created)
        for order in self.orderlist:
            if order.status in [order.Submitted, order.Accepted]:
                # 购买订单已提交给经纪人/经纪人接受卖单 - 啥也不干
                continue

            # 检查订单是否已完成
            # 注意：如果现金不足，经纪人可能会拒绝订单
            if order.status in [order.Completed]:
                if order.isbuy():
                    self.log('BUY EXECUTED, %.2f, Cost: %.2f, Comm: %.2f' %
                             (order.executed.price,
                              order.executed.value,
                              order.executed.comm),
                             # doprint=True
                             )

                    self.buyprice = order.executed.price
                    self.comms.append(round(order.executed.comm, 2))
                elif order.issell():
                    self.log('SELL EXECUTED, %.2f, Cost: %.2f, Comm: %.2f' %
                             (order.executed.price,
                              order.executed.value,
                              order.executed.comm),
                             # doprint=True
                             )

                    self.comms.append(round(order.executed.comm, 2))

                self.bar_executed = len(self)

            elif order.status in [order.Canceled, order.Margin, order.Rejected]:
                # print([order.Canceled, order.Margin, order.Rejected])
                # 5 7 8
                self.log('Order %s' % str(order.status))

            # 订单终止
            self.orderlist.remove(order)

    def notify_trade(self, trade):
        if not trade.isclosed:
            return

        self.log('OPERATION PROFIT, GROSS %.2f, NET %.2f' %
                 (trade.pnl, trade.pnlcomm))

    # 当经过一个K线柱的时候 next() 方法就会被调用一次。
    def next(self):

        for o in self.orderlist:
            self.cancel(o)  # 取消以往所有订单
            self.orderlist = []  # 置空

        for stock in self.datas:
            self.log('%s, Open: %.2f, Close, %.2f' % (stock._name, stock.open[0], stock.close[0]))
            # 检查我们是否入市
            if not self.getposition(stock):
                # 尚未入市...如果...的话，我们可能会买
                if stock.close[0] > self.mas[stock][0]:
                    # 买，买，买！！！ (with all possible default parameters)
                    self.log('%s BUY CREATE, %.2f' % (stock._name, stock.close[0]))
                    # self.buy(data=data, size=self.p.stake)
                    order = self.buy(data=stock, size=self.p.stake, exectype=bt.Order.Market, valid=bt.Order.DAY)
                    self.orderlist.append(order)
            else:
                # 已经入市...我们可能会出售
                if stock.close[0] < self.mas[stock][0]:
                    # 卖，卖，卖！！！ (with all possible default parameters)
                    self.log('%s SELL CREATE, %.2f' % (stock._name, stock.close[0]))
                    order = self.sell(data=stock, size=self.p.stake, exectype=bt.Order.Market, valid=bt.Order.DAY)
                    self.orderlist.append(order)

    def stop(self):
        print('==================== Results ====================')
        print('Starting Value - %.2f' % self.broker.startingcash)
        print('Ending   Value - %.2f' % self.broker.getvalue())
        print('=================================================')


def run_strategy():

    # 实例化 Cerebro 引擎
    cerebro = bt.Cerebro(tradehistory=True)

    # 添加一个策略
    cerebro.addstrategy(TestStrategy)

    # 参数调优
    # strats = cerebro.optstrategy(
    #     TestStrategy,
    #     maperiod=range(10, 31))

    data_api = Stock(datetime.now(), 365, 0)
    stock_pool = data_api.selectStockPoolByRSI()
    for code in tqdm(stock_pool):
        # print(code)
        now = data_api.now.strftime(data_api.TIME_STR)
        OneYearBefore = (data_api.now - data_api.before).strftime(data_api.TIME_STR)
        ohlc_data = data_api.getDailyKV(code, OneYearBefore, now)
        ohlc_data = ohlc_data.iloc[::-1]
        # print(ohlc_data)
        stock = bt.feeds.PandasData(dataname=ohlc_data, nocase=True)
        # 将数据添加到 Cerebro
        cerebro.adddata(stock, name=code)

    # 设定我们想要的初始金额
    cerebro.broker.setcash(100000.0)
    # 设定手续费，国内一般是是0.0003
    cerebro.broker.setcommission(0.003)
    cerebro.addanalyzer(bt.analyzers.SharpeRatio, _name="SharpeRatio")
    cerebro.addanalyzer(bt.analyzers.AnnualReturn, _name="AannualReturn")
    cerebro.addanalyzer(bt.analyzers.DrawDown, _name="DrawDown")
    cerebro.addanalyzer(bt.analyzers.TradeAnalyzer, _name="Trade")
    cerebro.addanalyzer(bt.analyzers.PyFolio)
    strats = cerebro.run()
    strat = strats[0]

    print(len(strat.comms))
    print(strat.comms)

    pyfolio = strat.analyzers.getbyname('pyfolio')
    returns, positions, transactions, gross_lev = pyfolio.get_pf_items()

    # print('Pyfolio:', strat.analyzers.pyfolio.get_analysis())
    sharpe_ratio = strat.analyzers.SharpeRatio.get_analysis()['sharperatio']
    max_drawdown = strat.analyzers.DrawDown.get_analysis()['max']['drawdown']
    annual_reaturn = strat.analyzers.AannualReturn.get_analysis()
    trade = strat.analyzers.Trade.get_analysis()
    transactions['commision'] = strat.comms

    print('================== Performance ==================')
    print('Sharpe Ratio:', sharpe_ratio)
    print('Max DrawDown:', max_drawdown, '%')
    for k, v in annual_reaturn.items():
        v = round(v, 2)
        print(k, 'Aannual Return:', v)
        annual_reaturn[k] = v
    print('Trade:', trade)
    print('=================================================')

    # print("================== returns ==================")
    # print(returns)
    # print("================== positions ==================")
    # print(positions)
    # print("================== transactions ==================")
    # print(transactions)
    # print("================== gross_lev ==================")

    return round(sharpe_ratio, 2), round(max_drawdown, 2), annual_reaturn, transactions
