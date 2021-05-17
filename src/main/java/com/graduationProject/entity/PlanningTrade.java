package com.graduationProject.entity;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.graduationProject.dto.State;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.experimental.Accessors;


import java.util.ArrayList;
import java.util.List;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @author : Niaowuuu
 * @ClassName : PlanningTrade
 * @说明 : 计划交易实体类
 * @创建日期 : 2021/5/3
 * @since : 1.0
 */

@Data
@Accessors(chain = true)
@EqualsAndHashCode(callSuper = true)
public class PlanningTrade extends State {

    // 计划交易ID
    private String id;

    // 关联策略ID
    private String strategyID;

    // 交易股票
    private String stockID;

    // 交易份额（买卖用正负来表示）
    private Double amount;

    public static PlanningTrade deserialize(byte[] data) {
        JSONObject json = JSON.parseObject(new String(data, UTF_8));
        return deserialize(json);
    }

    public static PlanningTrade deserialize(JSONObject json) {
        String strategyID = json.getString("strategyID");
        String id = json.getString("id");
        String stockID = json.getString("stockID");
        Double amount = json.getDouble("amount");

        return createInstance(strategyID, id, stockID, amount);
    }

    public static List<PlanningTrade> deserializeList(byte[] data) {
        JSONArray json = JSON.parseArray(new String(data, UTF_8));
        List<PlanningTrade> planningTrades = new ArrayList<>();
        for (int i = 0; i < json.size(); i++) {
            planningTrades.add(deserialize(json.getJSONObject(i)));
        }

        return planningTrades;
    }


    public static PlanningTrade createInstance(String StrategyID, String id, String stockID, Double amount) {
        return new PlanningTrade().setStrategyID(StrategyID).setId(id).setStockID(stockID).setAmount(amount);
    }

}
