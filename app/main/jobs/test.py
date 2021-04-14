from datetime import datetime
from app.data_source.api.tushare_api import Stock


def test_on_time():
    data_api = Stock(datetime.now(), 365)
    data_api.updateLocal()
    print(datetime.now())
