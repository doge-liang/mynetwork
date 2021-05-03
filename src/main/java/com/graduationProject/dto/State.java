package com.graduationProject.dto;

import org.json.JSONObject;

import static java.nio.charset.StandardCharsets.UTF_8;

/**
 * @类名 : StateDTO
 * @说明 : 状态传输实体类
 * @创建日期 : 2021/5/3
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
public class State {

    private String key;

    public State() {

    }

    String getKey() {
        return this.key;
    }

    public String[] getSplitKey() {
        return State.splitKey(this.key);
    }

    public static byte[] serialize(Object object) {
        String jsonStr = new JSONObject(object).toString();
        return jsonStr.getBytes(UTF_8);
    }

    public static String makeKey(String[] keyParts) {
        return String.join(":", keyParts);
    }

    public static String[] splitKey(String key) {
        System.out.println("splitting key " + key + "   " + java.util.Arrays.asList(key.split(":")));
        return key.split(":");
    }
}
