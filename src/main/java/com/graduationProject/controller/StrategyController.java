package com.graduationProject.controller;

import com.alibaba.fastjson.JSON;
import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.*;
import com.graduationProject.entity.*;
import com.graduationProject.utils.IDGenerator;
import com.graduationProject.utils.RestMock;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;
import java.nio.charset.StandardCharsets;
import java.util.*;

/**
 * StrategyController
 * <p>
 * 策略控制器
 *
 * @author : Niaowuuu
 * @since : 1.0
 * <p>
 * 2021/3/24
 */
@RestController
@RequestMapping("/strategy")
public class StrategyController {

    @Autowired
    RestMock restApi;

    @GetMapping("/list")
    public ResultDTO<List<StrategyDTO>> getAllStrategies(HttpServletRequest request) {
        // public ResultDTO<List<Strategy>> getAllStrategies() throws IOException, ContractException {
        // String userName = (String) map.get("userName");
        // String userSecret = (String) map.get("userSecret");
        // String orgName = (String) map.get("orgName");
        HttpSession session = request.getSession();
        Map loginUserMap = (Map) session.getAttribute("loginUser");
        String userName = (String) loginUserMap.get("userName");
        String userSecret = (String) loginUserMap.get("userSecret");
        String orgName = (String) loginUserMap.get("orgName");

        if (userName == null) {
            userName = "admin";
            userSecret = "adminpw";
            orgName = "Subscriber";
        }
        System.out.printf("userName:%s, password:%s, orgName: %s\n", userName, userSecret, orgName);
        try {
            User user = new User(userName, userSecret, orgName);
            user.doEnroll();
            byte[] result = user.doQuery("GetAllStrategies", "");
            if (result.length != 0) {
                List<StrategyDTO> strategies = StrategyDTO.deserializeList(result);
                System.out.println(strategies);
                return new ResultDTO<>(strategies);
            }
            return new ResultDTO<>(StatusCode.NOT_FOUND);
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
        byte[] result;
        Strategy strat = Strategy.createInstance(
                id,
                name,
                "",
                res.getMaxDrawdown(),
                res.getAnnualReturn(),
                res.getSharpeRatio(),
                state,
                null
        );
        List<Trade> newTrades = new ArrayList<>();
        List<PlanningTrade> newPlanningTrades = new ArrayList<>();
        List<Position> newPositions = new ArrayList<>();
        for (AnalyseReturn.Trade t : res.getTrades()
        ) {

            newTrades.add(
                    Trade.createInstance(
                            id,
                            // "",
                            String.valueOf(idWorker.nextId()),
                            t.getStockID(),
                            t.getAmount(),
                            t.getCommission(),
                            t.getDateTime(),
                            t.getPrice()
                    )
            );
        }

        // 获取所有交易
        System.out.println("获取所有交易");
        result = admin.doQuery("GetTradesPageByStrategyID", id, "");
        TradeOutput to = JSON.parseObject(new String(result, StandardCharsets.UTF_8), TradeOutput.class);
        List<Trade> oldTrades = new ArrayList<>();
        if (Objects.nonNull(to) && to.getTrades().size() == 40) {
            oldTrades = new ArrayList<>(to.getTrades());

            while (to.getTrades().size() == 40) {
                System.out.println("有交易");
                result = admin.doQuery("GetTradesPageByStrategyID", id, to.getBookmark());
                to = JSON.parseObject(new String(result, StandardCharsets.UTF_8), TradeOutput.class);
                System.out.println(to);
                System.out.println(to.getBookmark());
                oldTrades.addAll(to.getTrades());
            }
        }
        System.out.println(oldTrades);
        System.out.println(oldTrades.size());

        // 出现新的交易记录
        if (oldTrades.size() < newTrades.size()) {
            System.out.println("出现新的交易记录");
            // 原来的记录非空，需要在生成的交易中排除
            if (oldTrades.size() != 0) {
                for (Iterator<Trade> newTradesIter = newTrades.iterator(); newTradesIter.hasNext(); ) {
                    Trade newTrade = newTradesIter.next();
                    System.out.println(newTrade);
                    for (Trade oldTrade : oldTrades) {
                        if (newTrade.equals(oldTrade)) {
                            newTradesIter.remove();
                            System.out.println("REMOVED");
                        }
                    }
                }
            }
            System.out.println("获取交易信号");
            result = admin.doQuery("GetPlanningTradesByStrategyID", id);
            PlanningTradeOutput pto = JSON.parseObject(new String(result, StandardCharsets.UTF_8), PlanningTradeOutput.class);
            System.out.println(pto);
            if (Objects.nonNull(pto) && pto.getPlanningTrades().size() != 0) {
                List<PlanningTrade> oldPlanningTrades = pto.getPlanningTrades();
                // 如果有计划交易
                if (oldPlanningTrades.size() != 0) {
                    System.out.println("有交易信号，尝试 ID 替换");
                    // 将计划交易变成已完成的交易
                    for (PlanningTrade planningTrade : oldPlanningTrades) {
                        for (Trade trade : newTrades) {
                            if (trade.equals(planningTrade)) {
                                trade.setId(planningTrade.getId());
                                System.out.println(trade);
                            }
                        }
                    }
                }
            }
            // 交易上链
            System.out.println("要添加的交易记录");
            System.out.println(newTrades);
            System.out.println(newTrades.size());
            if (newTrades.size() > 0) {
                admin.doInvoke("AddTrades", JSON.toJSONString(newTrades));
            }
        }

        for (AnalyseReturn.PlanningTrade pt : res.getPlanningTrades()) {
            newPlanningTrades.add(
                    PlanningTrade.createInstance(
                            id,
                            String.valueOf(idWorker.nextId()),
                            pt.getStockID(),
                            pt.getAmount()
                    )
            );
        }
        for (AnalyseReturn.Position p : res.getPositions()) {
            newPositions.add(
                    Position.createInstance(
                            id,
                            String.valueOf(idWorker.nextId()),
                            p.getStockID(),
                            p.getAmount()
                    )
            );
        }

        // 私有策略要清空原有的私有数据
        if (strat.getState() == 1) {
            // if (state == 1) {
            result = admin.doQuery("GetPlanningTradesByStrategyID", id);
            PlanningTradeOutput pto = JSON.parseObject(new String(result, StandardCharsets.UTF_8), PlanningTradeOutput.class);
            System.out.println(pto);
            result = admin.doQuery("GetPositionsByStrategyID", id);
            PositionOutput po = JSON.parseObject(new String(result, StandardCharsets.UTF_8), PositionOutput.class);
            System.out.println(po);

            if (Objects.nonNull(pto) || Objects.nonNull(po)) {
                admin.doInvoke("DelPrivateData",
                        JSON.toJSONString(pto.getPlanningTrades()),
                        JSON.toJSONString(po.getPositions()));
            }
        }
        System.out.println(newPlanningTrades);
        System.out.println(newPositions);
        // 私有数据部分上链
        admin.doInvoke("AddPlanningTrades",
                Strategy.serialize(strat),
                JSON.toJSONString(newPlanningTrades));
        admin.doInvoke("AddPositions",
                Strategy.serialize(strat),
                JSON.toJSONString(newPositions));

        // result = admin.doQuery("GetPlanningTradesByStrategyID", id);
        // PlanningTradeOutput pto = JSON.parseObject(new String(result, StandardCharsets.UTF_8), PlanningTradeOutput.class);
        // System.out.println(pto);
        // result = admin.doQuery("GetPositionsByStrategyID", id);
        // PositionOutput po = JSON.parseObject(new String(result, StandardCharsets.UTF_8), PositionOutput.class);
        // System.out.println(po);
        return new ResultDTO<>(res);
        // return new ResultDTO<>(StatusCode.SUCCESS);

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
