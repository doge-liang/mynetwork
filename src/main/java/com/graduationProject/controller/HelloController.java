package com.graduationProject.controller;

import com.graduationProject.utils.RestMock;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @类名 : HelloController
 * @说明 : 测试Controller
 * @创建日期 : 2021/3/23
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@RestController
public class HelloController {

    @Autowired
    RestMock restApi;

    @RequestMapping("/hello")
    public String handle01() {
        return "Hello, Springboot 2!";
    }

    @RequestMapping("/hello/flask")
    public Object handle02() {
        return restApi.helloFlask();
    }


}