package com.graduationProject.service;

import com.graduationProject.dto.PlanningTradeOutput;
import org.springframework.stereotype.Service;


/**
 * @ClassName : PlanningTradeService
 * @说明 : 交易信号服务接口
 * @创建日期 : 2021/5/13
 * @author : Niaowuuu
 * @since : 1.0
 */

@Service
public interface PlanningTradeService {

    /**
     * 获取信号列表
     */
    PlanningTradeOutput getAllPlanningTradeByStrategyID(String strategyID);

}
