package com.graduationProject.controller;

import com.alibaba.fastjson.JSON;
import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.PositionOutput;
import com.graduationProject.dto.ResultDTO;
import com.graduationProject.entity.User;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import javax.servlet.http.HttpSession;
import java.io.IOException;
import java.nio.charset.StandardCharsets;
import java.util.Map;

/**
 * PositionController
 * <p>
 * 持仓控制器
 * <p>
 * Created : 2021/5/18 2:03
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@RestController
@RequestMapping("/strategy/{id}/position")
public class PositionController {

    @GetMapping("/list")
    public ResultDTO<PositionOutput> getAllPositions(HttpSession session, @PathVariable String id) {
        Map map = (Map) session.getAttribute("loginUser");
        String userName = (String) map.get("userName");
        String userSecret = (String) map.get("userSecret");
        String orgName = (String) map.get("orgName");

        try {
            User user = new User(userName, userSecret, orgName);
            // User user = new User("admin", "adminpw", "Provider");
            if (user.doEnroll()) {
                byte[] result = user.doQuery("GetPositionsByStrategyID", id);
                PositionOutput po = JSON.parseObject(new String(result, StandardCharsets.UTF_8), PositionOutput.class);
                System.out.println(po);
                return new ResultDTO<>(po);
            }
            return new ResultDTO<>(StatusCode.NOT_FOUND);
        } catch (Exception e) {
            e.printStackTrace();
            return new ResultDTO<>(StatusCode.INTERNAL_ERROR);
        }
    }
}
