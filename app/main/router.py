from flask import request
from flask import redirect
from flask import abort
from flask import make_response
from flask import Blueprint
from pathlib import Path

router = Blueprint("router", __name__)


@router.route('/upload', method=['POST'])
def upload():
    f = request.files['file']
    strategy_path = Path.cwd() / '..' / 'data_source' / 'local'
    f.save(strategy_path / f.filename)
    return True


# 简单的路由示例
@router.route('/')
def index():
    user_agent = request.headers.get('User-Agent')
    return '<p>Your browser is %s</p>' % user_agent


# 动态路由 name 作为 url 参数传入
@router.route('/hello/<name>')
def hello(name):
    return '<h1>Hello %s!</h1>' % name


# 带有状态码的返回值
@router.route('/404')
def not_found():
    return '<h1>Not Found</h1>', 404


# 重定向测试
@router.route('/redirect')
def redirect_to():
    return redirect('https://www.baidu.com')


# abort 错误处理函数测试
@router.route('/abort')
def abort_test():
    abort(404)
    # abort() 被调用之后不会再回到这个函数
    # 而是直接将控制权交给 Web 服务器
    return '<h1>Not Found</h1>', 404


# make_response 建立响应测试
@router.route('/test_response')
def response_test():
    response = make_response('<h1>This document carries a cookie!</h1>')
    response.set_cookie('answer', '42')
    return response
