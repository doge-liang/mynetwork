package com.graduationProject.consts;

/**
 * @类名 : StatusCode
 * @说明 : 状态码枚举类
 * @创建日期 : 2021/4/22
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
public enum StatusCode {

    LOGIN_FAIL(1000, "LOGIN FAILED"),
    NOT_FOUND(404, "404 NOT FOUND"),
    INTERNAL_ERROR(500, "INTERNAL_ERROR"),
    SUCCESS(200, "SUCCESS");

    private Integer code;
    private String msg;


    StatusCode(Integer code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    public Integer getCode() {
        return this.code;
    }

    public String getMsg() {
        return this.msg;
    }
}
