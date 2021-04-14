from flask import Flask
from flask_apscheduler import APScheduler
from config import APSchedulerJobConfig

app = Flask(__name__)

# 定时任务，导入配置
# APSchedulerJobConfig 就是在 config.py文件中的 类 名称。
app.config.from_object(APSchedulerJobConfig)

if __name__ == '__main__':
    # 初始化Flask-APScheduler，定时任务
    scheduler = APScheduler()
    scheduler.init_app(app)
    scheduler.start()
    # response_hello()

    app.run(debug=True)
