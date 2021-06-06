package com.graduationProject.controller;

import com.alibaba.fastjson.JSON;
import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.Page;
import com.graduationProject.dto.ResultDTO;
import com.graduationProject.dto.TradeOutput;
import com.graduationProject.entity.Trade;
import com.graduationProject.entity.User;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpSession;
import java.nio.charset.StandardCharsets;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

/**
 * TradeController
 * <p>
 * 交易记录控制类
 * <p>
 * Created : 2021/5/14 15:53
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@RestController
@RequestMapping("/strategy/{id}/trade")
@Slf4j
public class TradeController {

    @GetMapping("/list")
    // public ResultDTO<List<Strategy>> getAllStrategies(@RequestBody Map map) throws IOException, ContractException {
    public ResultDTO<Page<List<Trade>>> getTradesPageByStrategyID(
            HttpSession session,
            @PathVariable("id") String id,
            @RequestParam(required = false, defaultValue = "") String bookmark,
            @RequestParam(required = false, defaultValue = "40") Integer pageSize) {
        // System.out.println(bookmark);
        log.info(bookmark);
        try {
            Map map = (Map) session.getAttribute("loginUser");
            String userName = (String) map.get("userName");
            String userSecret = (String) map.get("userSecret");
            String orgName = (String) map.get("orgName");

            User user = new User(userName, userSecret, orgName);
            user.doEnroll();
            byte[] result = user.doQuery("GetTradesPageByStrategyID", id, bookmark.replaceAll(" ", "\u0000"),
                    pageSize.toString());
            // byte[] result = user.doQuery("GetTradesPageByStrategyID", id, bookmark, pageSize.toString());
            if (result.length != 0) {
                TradeOutput to = JSON.parseObject(new String(result, StandardCharsets.UTF_8), TradeOutput.class);
                Page<List<Trade>> trades = new Page<>(to.getTrades(), to.getBookmark(), pageSize);
                // System.out.println(trades);
                return new ResultDTO<>(trades);
            }
            return new ResultDTO<>(new Page<>(new ArrayList<>(), "", pageSize));
        } catch (Exception e) {
            e.printStackTrace();
            return new ResultDTO<>(StatusCode.INTERNAL_ERROR);
        }
    }

    @GetMapping("/deleteAll")
    public ResultDTO<?> delTrades(@PathVariable String id) {
        try {
            User user = new User("user1", "user1pw", "Subscriber");
            user.doEnroll();
            byte[] result = user.doQuery("GetTradesPageByStrategyID", id, "", "40");
            if (result.length != 0) {
                Page<List<Trade>> trades = Trade.deserializePage(result, 40);
                System.out.println(trades);
                while (trades.getTotalPage() != 0) {
                    user.doInvoke("DelTradesByStrategyID", JSON.toJSONString(trades.getData()));
                    result = user.doQuery("GetTradesPageByStrategyID", id, "", "40");
                    trades = Trade.deserializePage(result, 40);
                    System.out.println(trades);
                }
            }
            return new ResultDTO<>(StatusCode.SUCCESS);
        } catch (Exception e) {
            e.printStackTrace();
            return new ResultDTO<>(StatusCode.INTERNAL_ERROR);
        }
    }
}
