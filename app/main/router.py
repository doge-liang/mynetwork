from flask import request
from flask import redirect
from flask import abort
from flask import make_response
from flask import jsonify

from app.main.main import app
from app.main.strategy import RSI
import json


# 简单的路由示例
@app.route('/')
def index():
    user_agent = request.headers.get('User-Agent')
    return '<p>Your browser is %s</p>' % user_agent


# 动态路由 name 作为 url 参数传入
@app.route('/hello/<name>')
def hello(name):
    return '<h1>Hello %s!</h1>' % name


# 带有状态码的返回值
@app.route('/404')
def not_found():
    return '<h1>Not Found</h1>', 404


# 重定向测试
@app.route('/redirect')
def redirect_to():
    return redirect('https://www.baidu.com')


# abort 错误处理函数测试
@app.route('/abort')
def aborttest():
    abort(404)
    # abort() 被调用之后不会再回到这个函数
    # 而是直接将控制权交给 Web 服务器
    return '<h1>Not Found</h1>', 404


# make_response 建立响应测试
@app.route('/test_response')
def responseTest():
    response = make_response('<h1>This document carries a cookie!</h1>')
    response.set_cookie('answer', '42')
    return response


@app.route('/springboot')
def response_hello():
    sharpe_ratio, max_drawdown, annual_return, transactions, positions, planning_trades = RSI.run_strategy()
    transactions.reset_index(inplace=True)
    transactions_j = json.loads(transactions.to_json(orient="index", force_ascii=False))
    response = jsonify({
        "sharpe_ratio": sharpe_ratio,
        "max_drawdown": max_drawdown,
        "annual_return": annual_return,
        "transaction": transactions_j,
    })
    response.status_code = 200
    return response