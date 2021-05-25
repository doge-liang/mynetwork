from flask import jsonify, Blueprint
import importlib

analysis = Blueprint("analysis", __name__)


@analysis.route('/<strategy>')
def handle_analysis(strategy):
    strategy_path = "app.main.strategy." + strategy
    strategy_package = importlib.import_module(strategy_path)
    run_strategy = getattr(strategy_package, "run_strategy")
    sharpe_ratio, max_drawdown, annual_return, transactions, positions, planning_trades = run_strategy()
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
