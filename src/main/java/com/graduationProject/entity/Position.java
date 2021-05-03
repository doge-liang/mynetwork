package com.graduationProject.entity;

import lombok.Data;
import lombok.experimental.Accessors;
import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;
import org.json.JSONObject;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @类名 : Position
 * @说明 : 持仓记录
 * @创建日期 : 2021/4/16
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@DataType
@Data
@Accessors(chain = true)
public class Position {

    // 持仓股代码
    @Property
    private String stockID;
    // 仓位
    @Property
    private Double value;

    public static Position deserialize(byte[] data) {
        JSONObject json = new JSONObject(new String(data, UTF_8));
        return deserialize(json);
    }

    public static Position deserialize(JSONObject json) {
        String stockID = json.getString("stockID");
        Double value = json.getDouble("value");
        return createInstance(stockID, value);
    }

    public static Position createInstance(String stockID, Double value) {
        return new Position().setStockID(stockID).setValue(value);
    }
}
