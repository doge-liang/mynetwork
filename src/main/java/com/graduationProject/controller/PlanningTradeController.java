package com.graduationProject.controller;

import com.alibaba.fastjson.JSONObject;
import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.PlanningTradeOutput;
import com.graduationProject.dto.ResultDTO;
import com.graduationProject.entity.User;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;


import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * PlanningTradeController
 * <p>
 * 计划交易控制器
 * <p>
 * Created : 2021/5/17 16:00
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@RestController
@RequestMapping("/strategy/{id}/planningTrade")
public class PlanningTradeController {

    @GetMapping("/list")
    public ResultDTO<PlanningTradeOutput> getAllPlanningTradesByStrategyID(@PathVariable("id") String id) {
        try {
            // String userName = (String) map.get("userName");
            // String userSecret = (String) map.get("userSecret");
            // User user = new User("user1", "user1pw", "Subscriber");
            User user = new User("admin", "adminpw", "Provider");
            user.doEnroll();
            byte[] result = user.doQuery("GetPlanningTradesByStrategyID", id);
            if (result.length != 0) {
                // JSONObject json = new JSONObject();
                PlanningTradeOutput planningTrades = JSONObject.parseObject(new String(result, UTF_8), PlanningTradeOutput.class);
                // System.out.println(planningTrades);
                return new ResultDTO<>(planningTrades);
            }
            return new ResultDTO<>(StatusCode.NOT_FOUND);
        } catch (Exception e) {
            e.printStackTrace();
            return new ResultDTO<>(StatusCode.INTERNAL_ERROR);
        }
    }
}
