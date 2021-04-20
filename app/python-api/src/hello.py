from flask import Flask
from flask import request

app = Flask(__name__)

@app.route('/index')
def index():
    user_agent = request.headers.get('User-Agent')
    return '<p>Your browser is %s</p>' % user_agent

@app.route('/hello')
def hello():
    return "<p>Hello World!</p>"

if __name__ == '__main__':
    app.run(debug=True, host="0.0.0.0", port=8888)
