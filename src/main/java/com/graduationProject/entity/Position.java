package com.graduationProject.entity;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import lombok.Data;
import lombok.experimental.Accessors;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @author : Niaowuuu
 * @ClassName : Position
 * @说明 : 持仓记录
 * @创建日期 : 2021/4/16
 * @since : 1.0
 */
@Data
@Accessors(chain = true)
public class Position {

    // 仓位 id
    private String id;

    // 关联策略 id
    private String strategyID;

    // 持仓股代码
    private String stockID;

    // 仓位
    private Double value;

    public static Position deserialize(byte[] data) {
        JSONObject json = JSON.parseObject(new String(data, UTF_8));
        return deserialize(json);
    }

    public static Position deserialize(JSONObject json) {
        String strategyID = json.getString("strategyID");
        String id = json.getString("id");
        String stockID = json.getString("stockID");
        Double value = json.getDouble("value");
        return createInstance(strategyID, id, stockID, value);
    }

    public static Position createInstance(String StrategyID, String id, String stockID, Double value) {
        return new Position().setStrategyID(StrategyID).setId(id).setStockID(stockID).setValue(value);
    }
}
