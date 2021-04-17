package com.graduationProject.entity;

import lombok.Data;
import org.hyperledger.fabric.contract.annotation.DataType;
import org.hyperledger.fabric.sdk.Enrollment;

import java.util.List;

/**
 * @类名 : Stategy
 * @说明 : 投资策略实体类
 * @创建日期 : 2021/3/24
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@Data
@DataType
public class Stategy {

//    策略ID
    private Integer id;
//    策略名
    private String name;
//    策略创建者的用户凭证
    private Enrollment creator;
//    策略订阅者的用户凭证
    private List<Enrollment> subscribers;
//    策略交易记录
    private List<Trade> trades;
//    持仓记录
    private List<Position> positiions;

}
