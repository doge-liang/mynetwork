package com.graduationProject.dto;

import com.alibaba.fastjson.JSON;
import com.graduationProject.consts.StatusCode;
import com.graduationProject.entity.Strategy;
import lombok.Data;
import lombok.experimental.Accessors;

import java.nio.charset.StandardCharsets;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @author : Niaowuuu
 * @ClassName : StateDTO
 * @说明 : 状态传输实体类
 * @创建日期 : 2021/5/3
 * @since : 1.0
 */
// @Data
// @Accessors(chain = true)
public class State {

    // private String key;
    //
    // public State() {
    //
    // }
    //
    // String getKey() {
    //     return this.key;
    // }
    //
    // public String[] getSplitKey() {
    //     return State.splitKey(this.key);
    // }

    public static String serialize(Object object) {
        // String jsonStr = new JSONObject(object).toString();
        // return jsonStr.getBytes(UTF_8);
        return JSON.toJSONString(object);
    }

    // public static String makeKey(String[] keyParts) {
    //     return String.join(":", keyParts);
    // }
    //
    // public static String[] splitKey(String key) {
    //     System.out.println("splitting key " + key + "   " + java.util.Arrays.asList(key.split(":")));
    //     return key.split(":");
    // }
}
