package com.graduationProject.entity;

import com.graduationProject.dto.State;
import lombok.Data;
import lombok.experimental.Accessors;
import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;
import org.json.JSONObject;

import javax.json.JsonObject;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @类名 : PlanningTrade
 * @说明 : 计划交易实体类
 * @创建日期 : 2021/5/3
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@DataType
@Data
@Accessors(chain = true)
public class PlanningTrade extends State {

    // 交易股票
    @Property
    private String stockID;

    // 交易份额（买卖用正负来表示）
    @Property
    private Double amount;

    public static PlanningTrade deserialize(byte[] data) {
        JSONObject json = new JSONObject(new String(data, UTF_8));
        return deserialize(json);
    }

    public static PlanningTrade deserialize(JSONObject json) {
        String stockID = json.getString("stockID");
        Double amount = json.getDouble("amount");

        return createInstance(stockID, amount);
    }

    public static PlanningTrade createInstance(String stockID, Double amount) {
        return new PlanningTrade().setStockID(stockID).setAmount(amount);
    }
}
