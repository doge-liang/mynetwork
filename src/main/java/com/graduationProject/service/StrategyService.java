package com.graduationProject.service;

import com.graduationProject.entity.Strategy;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @ClassName : StrategyService
 * @说明 : 策略资源服务接口
 * @创建日期 : 2021/5/13
 * @author : Niaowuuu
 * @since : 1.0
 */
@Service
public interface StrategyService {

    /**
     * 读取策略列表
     */
    List<Strategy> getStrategyList();

    /**
     * 发布策略
     */
    void distributeStrategy(Strategy strategy);

    /**
     * 删除策略
     */
    void deleteStrategy(Strategy strategy);

    /**
     * 订阅策略
     */
    void subscribe(Strategy strategy);

    /**
     * 取消订阅
     */
    void unSubscribe(Strategy strategy);
}
