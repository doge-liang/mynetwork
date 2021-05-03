package com.graduationProject.entity;

import com.graduationProject.dto.State;
import lombok.Data;
import lombok.experimental.Accessors;
import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.contract.annotation.Property;
import org.json.JSONArray;
import org.json.JSONObject;

import java.util.ArrayList;
import java.util.List;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @类名 : Stategy
 * @说明 : 投资策略实体类
 * @创建日期 : 2021/3/24
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@DataType
@Accessors(chain = true)
@Data
public class Strategy extends State {

    // type Strategy struct {
    //     ID             string          `json:"ID"`             // 策略 ID
    //     Name           string          `json:"name"`           // 策略名
    //     Provider       string          `json:"provider"`       // 发布者
    //     MaxDrawdown    float64         `json:"maxDrawdown"`    // 最大回撤
    //     AnnualReturn   float64         `json:"annualReturn"`   // 年化收益率
    //     SharpeRatio    float64         `json:"sharpeRatio"`    // 夏普率
    //     Subscribers    []string        `json:"subscribers"`    // 订阅者证书列表
    //     State          string          `json:"state"`          // 是否公开
    //     Trades         []Trade         `json:"trades"`         // 交易记录
    //     PlanningTrades []PlanningTrade `json:"planningTrades"` // 计划交易
    //     Positions      []Position      `json:"positions"`      // 持仓记录
    // }

    // 策略ID
    @Property
    private String ID;

    // 策略名
    @Property
    private String name;

    // 发布者
    @Property
    private String provider;

    // 最大回撤
    @Property
    private Double maxDrawdown;

    // 年化收益率
    @Property
    private Double annualReturn;

    // 夏普率
    @Property
    private Double sharpeRatio;

    // 是否公开
    @Property
    private String state;

    // 策略订阅者的用户凭证
    @Property
    private List<String> subscribers;

    // 策略交易记录
    @Property
    private List<Trade> trades;

    // 计划交易
    @Property
    private List<PlanningTrade> planningTrades;

    // 持仓记录
    @Property
    private List<Position> positions;

    public static Strategy deserialize(byte[] data) {
        JSONObject json = new JSONObject(new String(data, UTF_8));

        String issuer = json.getString("");
        String ID = json.getString("ID");
        String name = json.getString("name");
        String provider = json.getString("provider");
        Double maxDrawdown = json.getDouble("maxDrawdown");
        Double annualReturn = json.getDouble("annualReturn");
        Double sharpRatio = json.getDouble("sharpRatio");
        List<String> subscribers = new ArrayList<>();
        JSONArray subscribersArray = json.getJSONArray("subscribers");
        for (int i = 0; i < subscribersArray.length(); i++) {
            subscribers.add(subscribersArray.getString(i));
        }

        String state = json.getString("state");
        JSONArray TradesArray = json.getJSONArray("trades");
        List<Trade> trades = new ArrayList<>();
        for (int i = 0; i < TradesArray.length(); i++) {
            JSONObject tradeJSON = TradesArray.getJSONObject(i);
            trades.add(Trade.deserialize(tradeJSON));
        }

        JSONArray planningTradesArray = json.getJSONArray("planningTrades");
        List<PlanningTrade> planningTrades = new ArrayList<>();
        for (int i = 0; i < planningTradesArray.length(); i++) {
            JSONObject planningTradeJSON = planningTradesArray.getJSONObject(i);
            planningTrades.add(PlanningTrade.deserialize(planningTradeJSON));
        }

        JSONArray positionsArray = json.getJSONArray("positions");
        List<Position> positions = new ArrayList<>();
        for (int i = 0; i < positionsArray.length(); i++) {
            JSONObject positionJSON = positionsArray.getJSONObject(i);
            positions.add(Position.deserialize(positionJSON));
        }

        return createInstance(ID, name, provider, maxDrawdown, annualReturn, sharpRatio, state, subscribers, trades,
                planningTrades, positions);
    }


    public static Strategy createInstance(String ID,
                                          String name,
                                          String provider,
                                          Double maxDrawdown,
                                          Double annualReturn,
                                          Double sharpeRatio,
                                          String state,
                                          List<String> subscribers,
                                          List<Trade> trades,
                                          List<PlanningTrade> planningTrades,
                                          List<Position> positions) {
        return new Strategy()
                .setID(ID)
                .setName(name)
                .setProvider(provider)
                .setMaxDrawdown(maxDrawdown)
                .setAnnualReturn(annualReturn)
                .setSharpeRatio(sharpeRatio)
                .setState(state)
                .setSubscribers(subscribers)
                .setTrades(trades)
                .setPlanningTrades(planningTrades)
                .setPositions(positions);
    }
}
