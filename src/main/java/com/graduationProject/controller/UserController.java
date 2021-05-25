package com.graduationProject.controller;

import com.graduationProject.consts.StatusCode;
import com.graduationProject.dto.ResultDTO;
import com.graduationProject.entity.User;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpSession;
import java.util.HashMap;
import java.util.Map;

/**
 * @author : Niaowuuu
 * @ClassName : UserController
 * @说明 : 用户接口控制器
 * @创建日期 : 2021/3/24
 * @since : 1.0
 */
@RestController
@RequestMapping("/user")
public class UserController {

    @PostMapping("/login")
    // public ResultDTO<Object> login(HttpSession session,
    //                                @RequestParam("userName") String userName,
    //                                @RequestParam("userSecret") String userSecret,
    //                                @RequestParam("orgName") String orgName) throws Exception {
    public ResultDTO<Object> login(HttpSession session, @RequestBody Map map) throws Exception {

        String userName = (String) map.get("userName");
        String userSecret = (String) map.get("userSecret");
        String orgName = (String) map.get("orgName");

        System.out.printf("userName:%s, password:%s, orgName: %s\n", userName, userSecret, orgName);
        User user = new User(userName, userSecret, orgName);
        if (user.login()) {
            System.out.println(user.getEnrollment().getKey().toString());
            // HashMap<String, String> map = new HashMap<>();
            // map.put("userName", userName);
            // map.put("userSecret", userSecret);
            // map.put("orgName", orgName);
            session.setAttribute("loginUser", map);
            return new ResultDTO<>(StatusCode.SUCCESS);
        }
        return new ResultDTO<>(StatusCode.LOGIN_FAIL);
    }

    @PostMapping("/login-admin")
    public ResultDTO<Object> loginAdmin(String adminName, String adminSecret) throws Exception {
        User admin = new User(adminName, adminSecret, "Provider");
        // admin.doEnroll();
        // System.out.println(admin.getEnrollment().getCert());
        // System.out.println(admin.getEnrollment().getKey());
        if (admin.doEnroll()) {
            return new ResultDTO<>(StatusCode.SUCCESS);
        }
        return new ResultDTO<>(StatusCode.LOGIN_FAIL);
    }

    @PostMapping("/register")
    public ResultDTO<Object> register(@RequestBody Map map) throws Exception {
        String userName = (String) map.get("userName");
        String userSecret = (String) map.get("userSecret");
        System.out.printf("userName:%s, password:%s \n", userName, userSecret);
        User user = new User(userName, userSecret, "Subscriber");
        if (user.doRegister("admin", "adminpw")) {
            return new ResultDTO<Object>(StatusCode.SUCCESS);
        }
        return new ResultDTO<>(StatusCode.LOGIN_FAIL);
    }

}
