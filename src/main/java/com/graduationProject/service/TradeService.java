package com.graduationProject.service;

import com.graduationProject.entity.Trade;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @ClassName : TradeService
 * @说明 : 交易记录服务接口
 * @创建日期 : 2021/5/13
 * @author : Niaowuuu
 * @since : 1.0
 */
@Service
public interface TradeService {

    /**
     * 获取所有交易记录
     */
    List<Trade> getTradeListByStrategyID(String strategyID);

}
