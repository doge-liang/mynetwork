from flask import jsonify, Blueprint
from app.main.strategy import RSI

analysis = Blueprint('analysis', __name__)


@analysis.route('/RSI')
def RSI_analysis():
    sharpe_ratio, max_drawdown, annual_return, transactions, positions, planning_trades = RSI.run_strategy()
    response = jsonify({
        "sharpeRatio": sharpe_ratio,
        "maxDrawdown": max_drawdown,
        "annualReturn": annual_return,
        "trades": transactions,
        "positions": positions,
        "planningTrades": planning_trades,
    })
    response.status_code = 200
    return response
