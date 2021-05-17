package com.graduationProject.dto;

import lombok.Data;
import lombok.experimental.Accessors;

import java.util.List;

/**
 * AnalyseReturn
 * <p>
 * 回测结果数据传输实体
 * <p>
 * Created : 2021/5/14 0:59
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@Data
@Accessors(chain = true)
// @EqualsAndHashCode(callSuper = true)
public class AnalyseReturn {

    // 策略名
    private String name;
    // 最大回撤
    private Double maxDrawdown;
    // 年化收益率
    private Double annualReturn;
    // 夏普率
    private Double sharpeRatio;

    private List<PlanningTrade> planningTrades;

    private List<Position> positions;

    private List<Trade> trades;

    @Data
    public static class PlanningTrade {
        Double amount;
        String stockID;
    }

    @Data
    public static class Position {
        Double amount;
        String stockID;
    }

    @Data
    public static class Trade {
        String id;
        Double amount;
        Double commission;
        String dateTime;
        Double price;
        String stockID;
    }
}
