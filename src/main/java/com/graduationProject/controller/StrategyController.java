package com.graduationProject.controller;

import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.ResultDTO;
import com.graduationProject.entity.UserDO;
import org.hyperledger.fabric.gateway.ContractException;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.ta4j.core.Strategy;

import java.io.IOException;
import java.util.List;
import java.util.Map;

/**
 * @类名 : StrategyController
 * @说明 : 策略控制器
 * @创建日期 : 2021/3/24
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@RestController
@RequestMapping("/strategy")
public class StrategyController {

    @GetMapping("/all")
    public ResultDTO<List<Strategy>> getAllStrategies(@RequestBody Map map) throws IOException, ContractException {
        String userName = (String) map.get("userName");
        UserDO user = new UserDO(userName, "", "Subscriber");
        user.doQuery("GetAllStrategies", "");
        return new ResultDTO<List<Strategy>>(StatusCode.INTERNAL_ERROR);
    }
}
