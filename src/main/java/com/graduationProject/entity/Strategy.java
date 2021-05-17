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
 * Strategy
 *
 * 投资策略实体类
 *
 * Updated : 2021/5/14 0:42
 * @author : Niaowuuu
 * @version : 1.0
 */

@Data
@Accessors(chain = true)
@EqualsAndHashCode(callSuper = true)
public class Strategy extends State {


    // 策略ID
    private String id;

    // 策略名
    private String name;

    // 发布者
    private String provider;

    // 最大回撤
    private Double maxDrawdown;

    // 年化收益率
    private Double annualReturn;

    // 夏普率
    private Double sharpeRatio;

    // 是否公开
    private Integer state;

    // 策略订阅者的用户凭证
    private List<String> subscribers;

    public static Strategy deserialize(byte[] data) {
        JSONObject json = JSON.parseObject(new String(data, UTF_8));
        return deserialize(json);
    }

    public static Strategy deserialize(JSONObject json) {

        String id = json.getString("id");
        String name = json.getString("name");
        String provider = json.getString("provider");
        Double maxDrawdown = json.getDouble("maxDrawdown");
        Double annualReturn = json.getDouble("annualReturn");
        Double sharpRatio = json.getDouble("sharpeRatio");
        List<String> subscribers = new ArrayList<>();
        JSONArray subscribersArray = json.getJSONArray("subscribers");
        for (int i = 0; i < subscribersArray.size(); i++) {
            subscribers.add(subscribersArray.getString(i));
        }

        Integer state = json.getInteger("state");
        // JSONArray TradesArray = json.getJSONArray("trades");
        // List<Trade> trades = new ArrayList<>();
        // for (int i = 0; i < TradesArray.length(); i++) {
        //     JSONObject tradeJSON = TradesArray.getJSONObject(i);
        //     trades.add(Trade.deserialize(tradeJSON));
        // }

        // JSONArray planningTradesArray = json.getJSONArray("planningTrades");
        // List<PlanningTrade> planningTrades = new ArrayList<>();
        // for (int i = 0; i < planningTradesArray.length(); i++) {
        //     JSONObject planningTradeJSON = planningTradesArray.getJSONObject(i);
        //     planningTrades.add(PlanningTrade.deserialize(planningTradeJSON));
        // }

        // JSONArray positionsArray = json.getJSONArray("positions");
        // List<Position> positions = new ArrayList<>();
        // for (int i = 0; i < positionsArray.length(); i++) {
        //     JSONObject positionJSON = positionsArray.getJSONObject(i);
        //     positions.add(Position.deserialize(positionJSON));
        // }

        // return createInstance(id, name, provider, maxDrawdown, annualReturn, sharpRatio, state, subscribers, trades,
        //         planningTrades, positions);
        return createInstance(id, name, provider, maxDrawdown, annualReturn, sharpRatio, state, subscribers);
    }

    public static List<Strategy> deserializeList(byte[] data) {
        JSONArray json = JSON.parseArray(new String(data, UTF_8));
        List<Strategy> strategies = new ArrayList<>();
        for (int i = 0; i < json.size(); i++) {
            strategies.add(deserialize(json.getJSONObject(i)));
        }

        return strategies;
    }


    public static Strategy createInstance(String id,
                                          String name,
                                          String provider,
                                          Double maxDrawdown,
                                          Double annualReturn,
                                          Double sharpeRatio,
                                          Integer state,
                                          List<String> subscribers
    ) {
        return new Strategy()
                .setId(id)
                .setName(name)
                .setProvider(provider)
                .setMaxDrawdown(maxDrawdown)
                .setAnnualReturn(annualReturn)
                .setSharpeRatio(sharpeRatio)
                .setState(state)
                .setSubscribers(subscribers);
    }
}
