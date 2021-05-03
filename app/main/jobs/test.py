from datetime import datetime
from app.data_source.api.tushare_api import Stock


def update_local_data():
    data_api = Stock(datetime.now(), 365)
    data_api.updateLocal()
    print(datetime.now())
