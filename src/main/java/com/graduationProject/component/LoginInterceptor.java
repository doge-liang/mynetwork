package com.graduationProject.component;

import lombok.extern.slf4j.Slf4j;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

/**
 * @author : Niaowuuu
 * @ClassName : LoginInterceptor
 * @说明 : 登录拦截器
 * @创建日期 : 2021/4/23
 * @since : 1.0
 */
@Slf4j
public class LoginInterceptor implements HandlerInterceptor {


    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        String RequestURI = request.getRequestURI();
        log.info("preHandler 拦截的链接为：{}", RequestURI);

        // 登录检查
        HttpSession session = request.getSession();

        Object loginUserMap = session.getAttribute("loginUser");

        if (loginUserMap != null) {
            // request.getRequestDispatcher("/");
            // 放行
            log.info("loginUser: {}", loginUserMap);
            return true;
        } else {
            // 拦截
            request.setAttribute("msg", "未登录！");
            request.getRequestDispatcher("/").forward(request, response);
            return false;
        }
    }

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler,
                           ModelAndView modelAndView) throws Exception {

    }

    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception ex) throws Exception {

    }

    // public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
    //     log.info("==========登录状态拦截");
    //
    //     HttpSession session = request.getSession();
    //     log.info("sessionId为：" + session.getId());
    //
    //     // 获取用户信息，如果没有用户信息直接返回提示信息
    //     Object userInfo = session.getAttribute("userInfo");
    //     if (userInfo == null) {
    //         log.info("没有登录");
    //         response.getWriter().write("Please Login In");
    //         return false;
    //     } else {
    //         log.info("已经登录过啦，用户信息为：" + session.getAttribute("userInfo"));
    //     }
    //
    //     return true;
    // }

}
