package com.graduationProject.dto;

import com.graduationProject.entity.Trade;
import lombok.Data;

import java.util.List;

/**
 * TradeOutput
 * <p>
 * 交易记录传输类
 * <p>
 * Created : 2021/5/16 4:47
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@Data
public class TradeOutput {

    List<Trade> trades;

    String bookmark;
}
