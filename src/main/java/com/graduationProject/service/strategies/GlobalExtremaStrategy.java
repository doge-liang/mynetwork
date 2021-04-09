package com.graduationProject.service.strategies;

import org.ta4j.core.*;
import org.ta4j.core.analysis.criteria.TotalProfitCriterion;
import org.ta4j.core.indicators.helpers.*;
import org.ta4j.core.trading.rules.OverIndicatorRule;
import org.ta4j.core.trading.rules.UnderIndicatorRule;
import ta4jexamples.loaders.CsvTradesLoader;

public class GlobalExtremaStrategy {

    // We assume that there were at least one trade every 5 minutes during the whole
    // week
    private static final int NB_BARS_PER_WEEK = 12 * 24 * 7;

    /**
     * @param series the bar series
     * @return the global extrema strategy
     */
    public static Strategy buildStrategy(BarSeries series) {
        if (series == null) {
            throw new IllegalArgumentException("Series cannot be null");
        }

        ClosePriceIndicator closePrices = new ClosePriceIndicator(series);

        // Getting the max price over the past week
        HighPriceIndicator maxPrices = new HighPriceIndicator(series);
        HighestValueIndicator weekMaxPrice = new HighestValueIndicator(maxPrices, NB_BARS_PER_WEEK);
        // Getting the min price over the past week
        LowPriceIndicator minPrices = new LowPriceIndicator(series);
        LowestValueIndicator weekMinPrice = new LowestValueIndicator(minPrices, NB_BARS_PER_WEEK);

        // Going long if the close price goes below the min price
        MultiplierIndicator downWeek = new MultiplierIndicator(weekMinPrice, 1.004);
        Rule buyingRule = new UnderIndicatorRule(closePrices, downWeek);

        // Going short if the close price goes above the max price
        MultiplierIndicator upWeek = new MultiplierIndicator(weekMaxPrice, 0.996);
        Rule sellingRule = new OverIndicatorRule(closePrices, upWeek);

        return new BaseStrategy("GlobalExtremaStrategy", buyingRule, sellingRule);
    }

    public static void main(String[] args) {

        // Getting the bar series
        BarSeries series = CsvTradesLoader.loadBitstampSeries();

        // Building the trading strategy
        Strategy strategy = buildStrategy(series);

        // Running the strategy
        BarSeriesManager seriesManager = new BarSeriesManager(series);
        TradingRecord tradingRecord = seriesManager.run(strategy);
        System.out.println("Number of trades for the strategy: " + tradingRecord.getTradeCount());

        // Analysis
        System.out.println(
                "Total profit for the strategy: " + new TotalProfitCriterion().calculate(series, tradingRecord));
    }
}
