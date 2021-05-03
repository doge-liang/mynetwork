from __future__ import (absolute_import, division, print_function, unicode_literals)

# 导入 backtrader
import backtrader as bt
import pandas as pd
import json
# 导入其他包
from datetime import datetime
from datetime import timedelta

from app.data_source.api.tushare_api import Stock
from tqdm import tqdm


# 创建策策略
class RSI(bt.Strategy):
    # 自定义一些参数
    params = (
        ('maperiod', 25),
        ('short_period', 6),
        ('long_period', 9),
        ('printlog', False),
        ('stake', 100),
        ('stock_pool_size', 50),
        ('sell_period', 20),
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
        # self.mas = {stock: bt.ind.MovingAverageSimple(stock.close, period=self.p.maperiod) for stock in self.datas}
        self.rsi_short = {stock: bt.ind.RSI(stock.close, period=self.p.short_period, safediv=True) for stock in
                          self.datas}
        self.rsi_long = {stock: bt.ind.RSI(stock.close, period=self.p.long_period, safediv=True) for stock in
                         self.datas}
        self.rsi_diff = {stock: self.rsi_short[stock] - self.rsi_long[stock] for stock in self.datas}
        # 跟踪挂单
        self.orderlist = []
        self.buyprice = None
        self.comms = []
        self.bar_counter = 0
        self.pool = []

        self.sell_all = False

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

        # 选出初始股票池
        if len(self.pool) == 0:
            print("初始股票集：")
            for i in range(self.p.stock_pool_size):
                max_stock = self.datas[0]
                for stock_i in self.datas:
                    if self.rsi_diff[max_stock][0] < self.rsi_diff[stock_i][0] and stock_i not in self.pool:
                        max_stock = stock_i
                if max_stock not in self.pool:
                    print(max_stock._name, ":", self.rsi_diff[max_stock][0])
                    self.pool.append(max_stock)

        self.bar_counter += 1
        # 到了新一轮周期，修改卖出信号
        if self.bar_counter == self.p.sell_period:
            self.sell_all = True

        # 遍历股票集
        for stock in self.pool:
            #             self.log('%s, Open: %.2f, Close, %.2f' % (stock._name, stock.open[0], stock.close[0]))
            # 到了清仓周期或出现了卖信号都要卖
            # 卖信号:短线 rsi 从上往下穿过 长线 rsi ，则卖出
            sell_signal = self.rsi_diff[stock][0] < 0 and self.rsi_diff[stock][-1] > 0

            # 出现了买信号且在新一轮股票集中就买
            # 买信号:短线 rsi 从下往上穿过 长线 rsi ，则买入
            buy_signal = self.rsi_diff[stock][0] > 0 and self.rsi_diff[stock][-1] < 0
            # 检查我们是否入市
            # 尚未入市，可能会买入
            if not self.getposition(stock):
                # 尚未入市...如果...的话，我们可能会买
                if buy_signal:
                    # 买，买，买！！！ (with all possible default parameters)
                    self.log('%s BUY CREATE, %.2f' % (stock._name, stock.close[0]))
                    # self.buy(data=data, size=self.p.stake)
                    order = self.buy(data=stock, size=self.p.stake, exectype=bt.Order.Market, valid=bt.Order.DAY)
                    self.orderlist.append(order)
            # 入市了，可能会卖
            else:
                if sell_signal:
                    # 卖，卖，卖！！！ (with all possible default parameters)
                    self.log('%s SELL CREATE, %.2f' % (stock._name, stock.close[0]))
                    # order = self.sell(data=stock, size=self.p.stake, exectype=bt.Order.Market, valid=bt.Order.DAY)
                    order = self.close(stock)
                    self.orderlist.append(order)

        # 重新筛选股票池
        if self.bar_counter == self.p.sell_period:
            self.pool = []
            #             print("下一个周期：")
            for i in range(self.p.stock_pool_size):
                max_stock = self.datas[0]
                for stock_i in self.datas:
                    if self.rsi_diff[max_stock][0] < self.rsi_diff[stock_i][0] and stock_i not in self.pool:
                        max_stock = stock_i
                #                 print(max_stock._name, ":", self.rsi_diff[max_stock][0])
                if max_stock not in self.pool:
                    self.pool.append(max_stock)
            self.bar_counter = 0

        self.sell_all = False

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
    _ = cerebro.addstrategy(RSI)

    # 参数调优
    # strats = cerebro.optstrategy(
    #     TestStrategy,
    #     maperiod=range(10, 31))

    data_api = Stock(datetime.now(), 365, 0)
    for code in tqdm(data_api.getHS300()):
        try:
            ohlc_data = data_api.getDailyKVLocal(code)
        except Exception:
            ohlc_data = data_api.getDailyKVOnline(code)

        #         flag = False
        #         if flag:
        #             print("以下股票是近一年上市的新股，数据不足：")
        if (datetime.now() - ohlc_data.index[0]) < data_api.before - timedelta(3):
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

    sharpe_ratio = strat.analyzers.SharpeRatio.get_analysis()['sharperatio']
    max_drawdown = strat.analyzers.DrawDown.get_analysis()['max']['drawdown']
    returns = strat.analyzers.Returns.get_analysis()
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
    print('Aannual Return:', returns['rnorm100'], "%")
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
    transactions['ID'] = transactions.index + 1
    transactions = json.loads(transactions.to_json(orient="records", force_ascii=False))
    print(transactions)
    print('===========================================================================================\n\n\n')

    return round(sharpe_ratio, 2), round(max_drawdown, 2), returns[
        'rnorm100'], transactions, positions_dict, planning_trades


if __name__ == '__main__':
    Stock(datetime.now(), 365, 0).updateLocal()
    # run_strategy()
