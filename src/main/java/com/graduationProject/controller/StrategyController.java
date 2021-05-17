package com.graduationProject.controller;

import com.alibaba.fastjson.JSON;
import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.AnalyseReturn;
import com.graduationProject.dto.PlanningTradeOutput;
import com.graduationProject.dto.PositionOutput;
import com.graduationProject.dto.ResultDTO;
import com.graduationProject.entity.*;
import com.graduationProject.utils.IDGenerator;
import com.graduationProject.utils.RestMock;
import org.hyperledger.fabric.gateway.ContractException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

/**
 * @author : Niaowuuu
 * @ClassName : StrategyController
 * @说明 : 策略控制器
 * @创建日期 : 2021/3/24
 * @since : 1.0
 */
@RestController
@RequestMapping("/strategy")
public class StrategyController {

    @Autowired
    RestMock restApi;

    @GetMapping("/list")
    public ResultDTO<List<Strategy>> getAllStrategies(@RequestBody Map map) throws IOException, ContractException {
        // public ResultDTO<List<Strategy>> getAllStrategies() throws IOException, ContractException {
        String userName = (String) map.get("userName");
        String userSecret = (String) map.get("userSecret");
        String orgName = (String) map.get("orgName");
        if (userName == null) {
            userName = "admin";
            userSecret = "adminpw";
            orgName = "Subscriber";
        }
        try {
            User user = new User(userName, userSecret, orgName);
            user.doEnroll();
            byte[] result = user.doQuery("GetAllStrategies", "");
            if (result.length != 0) {
                List<Strategy> strategies = Strategy.deserializeList(result);
                System.out.println(strategies);
                return new ResultDTO<>(strategies);
            }
            return new ResultDTO<>(new ArrayList<>());
        } catch (Exception e) {
            e.printStackTrace();
            return new ResultDTO<>(StatusCode.INTERNAL_ERROR);
        }
    }

    @GetMapping("/distribute/{name}")
    public ResultDTO<Strategy> distribute(@PathVariable("name") String name) throws Exception {
        User admin = new User("admin", "adminpw", "Provider");
        admin.doEnroll();
        int state;
        IDGenerator idWorker = new IDGenerator(0, 0);
        String id = String.valueOf(idWorker.nextId());
        if (name.equals("RSI")) {
            state = 0;
            id = "6799980250357825536";
        } else if (name.equals("MA")) {
            state = 1;
            id = "6799979985630134272";
        } else {
            return new ResultDTO<>(StatusCode.NOT_FOUND);
        }
        AnalyseReturn res = restApi.helloFlask(name);
        // System.out.println(res);
        Strategy strat = Strategy.createInstance(
                id,
                name,
                "",
                res.getMaxDrawdown(),
                res.getAnnualReturn(),
                res.getSharpeRatio(),
                state,
                new ArrayList<>()
        );
        admin.doInvoke("Distribute", Strategy.serialize(strat));
        return new ResultDTO<>(strat);

    }

    @GetMapping("/update/{name}/{id}")
    public ResultDTO<AnalyseReturn> update(@PathVariable("name") String name, @PathVariable("id") String id) throws Exception {
        User admin = new User("admin", "adminpw", "Provider");
        admin.doEnroll();
        IDGenerator idWorker = new IDGenerator(0, 0);
        int state;
        if (name.equals("RSI")) {
            state = 0;
            id = "6799980250357825536";
        } else if (name.equals("MA")) {
            id = "6799979985630134272";
            state = 1;
        } else {
            return new ResultDTO<>(StatusCode.NOT_FOUND);
        }
        AnalyseReturn res = restApi.helloFlask(name);
        Strategy strat = Strategy.createInstance(
                id,
                name,
                "",
                res.getMaxDrawdown(),
                res.getAnnualReturn(),
                res.getSharpeRatio(),
                state,
                new ArrayList<>()
        );
        List<Trade> tradelist = new ArrayList<>();
        List<PlanningTrade> planningTradelist = new ArrayList<>();
        List<Position> positionlist = new ArrayList<>();
        for (AnalyseReturn.Trade t : res.getTrades()
        ) {

            tradelist.add(
                    Trade.createInstance(
                            id,
                            String.valueOf(idWorker.nextId()),
                            t.getStockID(),
                            t.getAmount(),
                            t.getCommission(),
                            t.getDateTime(),
                            t.getPrice()
                    )
            );
        }
        for (AnalyseReturn.PlanningTrade pt : res.getPlanningTrades()) {
            planningTradelist.add(
                    PlanningTrade.createInstance(
                            id,
                            String.valueOf(idWorker.nextId()),
                            pt.getStockID(),
                            pt.getAmount()
                    )
            );
        }
        for (AnalyseReturn.Position p : res.getPositions()) {
            positionlist.add(
                    Position.createInstance(
                            id,
                            String.valueOf(idWorker.nextId()),
                            p.getStockID(),
                            p.getAmount()
                    )
            );
        }
        String tradelistJSON = JSON.toJSONString(tradelist);
        String planningTradelistJSON = JSON.toJSONString(planningTradelist);
        String positionlistJSON = JSON.toJSONString(positionlist);
        // System.out.println(tradelistJSON);
        byte[] result = admin.doQuery("GetPlanningTradesByStrategyID", id);
        PlanningTradeOutput pto = JSON.parseObject(new String(result, StandardCharsets.UTF_8),
                PlanningTradeOutput.class);
        System.out.println(pto);
        result = admin.doQuery("GetPositionsByStrategyID", id);
        PositionOutput po = JSON.parseObject(new String(result, StandardCharsets.UTF_8),
                PositionOutput.class);
        System.out.println(po);
        admin.doInvoke("DelPrivateData",
                JSON.toJSONString(pto.getPlanningTrades()),
                JSON.toJSONString(po.getPositions()));

        admin.doInvoke("Update",
                Strategy.serialize(strat),
                tradelistJSON,
                planningTradelistJSON,
                positionlistJSON);
        // admin.doInvoke("UpdatePublicCollection", Strategy.serialize(strat),
        //         planningTradelistJSON,
        //         positionlistJSON);
        return new ResultDTO<>(res);

    }

    @GetMapping("/{id}/subscribe")
    public ResultDTO<?> subscribe(@PathVariable("id") String id) throws Exception {
        User user = new User("user1", "user1pw", "Subscriber");
        if (user.doEnroll()) {
            user.doInvoke("Subscribe", id);
            return new ResultDTO<>(StatusCode.SUCCESS);
        }
        return new ResultDTO<>(StatusCode.NOT_FOUND);
    }

    @GetMapping("/{id}/unsubscribe")
    public ResultDTO<?> unsubscribe(@PathVariable("id") String id) throws Exception {
        User user = new User("user1", "user1pw", "Subscriber");
        if (user.doEnroll()) {
            user.doInvoke("UnSubscribe", id);
            return new ResultDTO<>(StatusCode.SUCCESS);
        }
        return new ResultDTO<>(StatusCode.NOT_FOUND);
    }

}
