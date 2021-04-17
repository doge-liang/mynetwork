package com.graduationProject.entity;

import lombok.Data;

/**
 * @类名 : Position
 * @说明 : 持仓记录
 * @创建日期 : 2021/4/16
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@Data
public class Position {

//    持仓股代码
    private String stockcode;
//    仓位
    private Double Size;
}
