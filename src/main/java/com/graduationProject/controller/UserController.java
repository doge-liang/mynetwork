package com.graduationProject.controller;

import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.ResultDTO;
import com.graduationProject.entity.AdminDO;
import com.graduationProject.entity.UserDO;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

/**
 * @类名 : UserController
 * @说明 : 用户接口控制器
 * @创建日期 : 2021/3/24
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@RestController
@RequestMapping("/user")
public class UserController {

    @PostMapping("/login")
    public ResultDTO<Object> login(@RequestBody Map map) throws Exception {
//        return new ResultDTO<Object>(StatusCode.SUCCESS);
        String userName = (String) map.get("userName");
        String userSecret = (String) map.get("userSecret");
        System.out.printf("userName:%s, password:%s \n", userName, userSecret);
        UserDO user = new UserDO(userName, userSecret, "Subscriber");
        if (user.doEnroll()) {
            return new ResultDTO<Object>(StatusCode.SUCCESS);
        }
        return new ResultDTO<Object>(StatusCode.LOGIN_FAIL);
    }

    @PostMapping("/login-admin")
    public ResultDTO<Object> loginAdmin(String adminName, String adminSecret) throws Exception {
        AdminDO adminDO = new AdminDO(adminName, adminSecret, "Subscriber");
        if (adminDO.doEnroll()) {
            return new ResultDTO<Object>(StatusCode.SUCCESS);
        }
        return new ResultDTO<Object>(StatusCode.LOGIN_FAIL);
    }

    @PostMapping("/register")
    public ResultDTO<Object> register(@RequestBody Map map) throws Exception {
        String userName = (String) map.get("userName");
        String userSecret = (String) map.get("userSecret");
        System.out.printf("userName:%s, password:%s \n", userName, userSecret);
        UserDO user = new UserDO(userName, userSecret, "Subscriber");
        if (user.doRegister("admin", "adminpw")) {
            return new ResultDTO<Object>(StatusCode.SUCCESS);
        }
        return new ResultDTO<Object>(StatusCode.LOGIN_FAIL);
    }

}
