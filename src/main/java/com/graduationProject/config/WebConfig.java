package com.graduationProject.config;

import com.graduationProject.component.LoginInterceptor;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

/**
 * WebConfig
 * <p>
 * 拦截器配置类
 * <p>
 * Created : 2021/5/20 0:16
 *
 * @author : Niaowuuu
 * @version : 1.0
 */
@Configuration
public class WebConfig implements WebMvcConfigurer {

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // 排除 swagger
        registry.addInterceptor(new LoginInterceptor())
                .excludePathPatterns("/swagger-ui.html", "/webjars/springfox-swagger-ui/**", "/swagger-resources/**",
                        "/csrf", "/error")
                .excludePathPatterns("/", "/user/login", "/user/login-admin", "/strategy/distribute/**", "/strategy" +
                        "/update/**")
                .addPathPatterns("/**");


    }
}
