package com.graduationProject.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;

/**
 * @类名 : CrossOriginConfig
 * @说明 : 跨域请求配置类
 * @创建日期 : 2021/4/23
 * @作者 : Niaowuuu
 * @版本 : 1.0
 */
@Configuration
public class CrossOriginConfig extends WebMvcConfigurerAdapter {
    static final String ORIGINS[] = new String[] { "GET", "POST", "PUT", "DELETE" };
    @Override
    public void addCorsMappings(CorsRegistry registry) {
        registry.addMapping("/**").allowedOriginPatterns("*").allowCredentials(true).allowedMethods(ORIGINS).maxAge(3600);
    }
}
