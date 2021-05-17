package com.graduationProject.dto;

import com.graduationProject.consts.StatusCode;
import lombok.Data;

/**
 * @ClassName : ResultDTO
 * @说明 : 返回结果类
 * @创建日期 : 2021/4/22
 * @author : Niaowuuu
 * @since : 1.0
 */
@Data
public class ResultDTO<T> {

    /* 状态码 */
    private Integer code;

    /* 消息 */
    private String msg;

    /* 数据对象 */
    private T data;

    public ResultDTO(Integer code, String msg){
        this.code = code;
        this.msg = msg;
    }

    public ResultDTO(StatusCode statusCode) {
        this.code = statusCode.getCode();
        this.msg = statusCode.getMsg();
    }

    public ResultDTO(T data) {
        this(StatusCode.SUCCESS);
        this.data = data;
    }
}
