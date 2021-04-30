package com.graduationProject.entity;

import lombok.Data;

import java.util.Date;

/**
 * @类名 : Message
 * @说明 : 消息实体类
 * @创建日期 : 2021/4/9
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@Data
public class Trade {

    //    股票代码
    private String StockID;
    //    交易份额
    private Double Amount;
    //    手续费
    private Double Commission;
    //    交易时间
    private Date DateTime;
    //    交易价格
    private Double Price;

}
