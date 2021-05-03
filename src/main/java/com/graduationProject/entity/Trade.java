package com.graduationProject.entity;

import com.graduationProject.dto.State;
import lombok.Data;
import lombok.experimental.Accessors;
import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;
import org.json.JSONObject;

import java.util.Date;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @类名 : Message
 * @说明 : 消息实体类
 * @创建日期 : 2021/4/9
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@DataType
@Accessors(chain = true)
@Data
public class Trade extends State {


    // 交易 ID
    @Property
    private String ID;

    // 股票代码
    @Property
    private String stockID;

    // 交易份额
    @Property
    private Double amount;

    // 手续费
    @Property
    private Double commission;

    // 交易时间
    @Property
    private String dateTime;

    // 交易价格
    @Property
    private Double price;

    public static Trade deserialize(byte[] data) {
        JSONObject json = new JSONObject(new String(data, UTF_8));

        return deserialize(json);

    }

    public static Trade deserialize(JSONObject json) {

        String ID = json.getString("ID");
        String stockID = json.getString("stockID");
        Double amount = json.getDouble("amount");
        Double commission = json.getDouble("commission");
        String dateTime = json.getString("dateTime");
        Double price = json.getDouble("price");

        return createInstance(ID, stockID, amount, commission, dateTime, price);

    }

    public static Trade createInstance(String ID,
                                       String stockID,
                                       Double amount,
                                       Double commission,
                                       String dateTime,
                                       Double price) {

        return new Trade().setID(ID)
                .setStockID(stockID)
                .setAmount(amount)
                .setCommission(commission)
                .setDateTime(dateTime)
                .setPrice(price);
    }
}
