package com.graduationProject.dto;

import com.graduationProject.entity.PlanningTrade;
import lombok.Data;

import java.util.List;

/**
 * @author : Niaowuuu
 * @ClassName : PlanningTradeOutput
 * @说明 : 区块链网络返回的交易信号传输实体
 * @创建日期 : 2021/5/13
 * @since : 1.0
 */
@Data
public class PlanningTradeOutput {

    List<PlanningTrade> planningTrades;
    List<PlanningTradeHash> planningTradesHash;

    @Data
    public static class PlanningTradeHash {
        String id;
        String hashcode;
    }
}
