package com.graduationProject.entity;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.graduationProject.dto.Page;
import com.graduationProject.dto.State;
import lombok.Data;
import lombok.experimental.Accessors;

import java.util.ArrayList;
import java.util.Date;
import java.util.List;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @author : Niaowuuu
 * @ClassName : Message
 * @说明 : 消息实体类
 * @创建日期 : 2021/4/9
 * @since : 1.0
 */
@Accessors(chain = true)
@Data
public class Trade extends State {

    // 关联策略 id
    private String StrategyID;

    // 交易 id
    private String id;

    // 股票代码
    private String stockID;

    // 交易份额
    private Double amount;

    // 手续费
    private Double commission;

    // 交易时间
    private String dateTime;

    // 交易价格
    private Double price;

    @Override
    public boolean equals(Object obj) {
        // 地址相等
        if (this == obj) {
            return true;
        }

        // 比较对象不能为空
        if (obj == null) {
            return false;
        }

        if (obj instanceof Trade) {
            Trade other = (Trade) obj;
            return this.StrategyID.equals(other.StrategyID) && this.stockID.equals(other.stockID) && this.amount.equals(other.amount) && this.dateTime.equals(other.dateTime);
        }

        if (obj instanceof PlanningTrade) {
            PlanningTrade other = (PlanningTrade) obj;
            return this.StrategyID.equals(other.getStrategyID()) && this.stockID.equals(other.getStockID()) && this.amount.equals(other.getAmount());
        }
        return false;
    }

    public static Trade deserialize(byte[] data) {
        JSONObject json = JSON.parseObject(new String(data, UTF_8));

        return deserialize(json);

    }

    public static Trade deserialize(JSONObject json) {

        String StrategyID = json.getString("strategyID");
        String id = json.getString("id");
        String stockID = json.getString("stockID");
        Double amount = json.getDouble("amount");
        Double commission = json.getDouble("commission");
        String dateTime = json.getString("dateTime");
        Double price = json.getDouble("price");

        return createInstance(StrategyID, id, stockID, amount, commission, dateTime, price);

    }

    public static List<Trade> deserializeList(byte[] data) {
        JSONArray json = JSON.parseArray(new String(data, UTF_8));
        List<Trade> trades = new ArrayList<>();
        for (int i = 0; i < json.size(); i++) {
            trades.add(deserialize(json.getJSONObject(i)));
        }

        return trades;
    }

    public static Page<List<Trade>> deserializePage(byte[] data) {
        JSONObject json = JSON.parseObject(new String(data, UTF_8));
        JSONArray tradesArray = json.getJSONArray("trades");
        String bookmark = json.getString("bookmark");
        List<Trade> trades = new ArrayList<>();
        for (int i = 0; i < tradesArray.size(); i++) {
            trades.add(deserialize(tradesArray.getJSONObject(i)));
        }

        return new Page<>(trades, bookmark);
    }

    public static Trade createInstance(
            String StrategyID,
            String id,
            String stockID,
            Double amount,
            Double commission,
            String dateTime,
            Double price) {

        return new Trade()
                .setStrategyID(StrategyID)
                .setId(id)
                .setStockID(stockID)
                .setAmount(amount)
                .setCommission(commission)
                .setDateTime(dateTime)
                .setPrice(price);
    }

}
