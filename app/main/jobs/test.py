from datetime import datetime
from pathlib import Path

from app.data_source.api.tushare_api import Stock


def update_local_data():
    data_api = Stock(datetime.now(), 500)
    data_api.updateLocal()
    data_api.data_local_path = Path.cwd() / '../..' / 'data_source' / 'local'
    print(datetime.now())
