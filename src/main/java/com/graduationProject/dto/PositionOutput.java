package com.graduationProject.dto;

import com.graduationProject.entity.Position;
import lombok.Data;

import java.util.List;

/**
 * @ClassName : PositionOutput
 * @说明 : 区块链网络传回持仓信息实体
 * @创建日期 : 2021/5/13
 * @author : Niaowuuu
 * @since : 1.0
 */
@Data
public class PositionOutput {

    List<Position> positions;

    List<PositionHash> positionsHash;

    @Data
    public static class PositionHash {
        String id;
        String hashcode;
    }
}
