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

    private String ID;
    private String StockID;
    private Double Amount;
    private Double Commission;
    private Date DateTime;
    private Double Price;

}
