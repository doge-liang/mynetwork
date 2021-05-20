package com.graduationProject;

import org.hyperledger.fabric.gateway.*;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import springfox.documentation.swagger2.annotations.EnableSwagger2;

import java.io.IOException;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.Map;
import java.util.concurrent.TimeoutException;

/**
 * @ClassName : com.graduationProject.MainApplication
 * @说明 : 主程序
 * @创建日期 : 2021/3/23
 * @author : Niaowuuu
 * @since : 1.0
 */
// @EnableSwagger2
@SpringBootApplication
public class MainApplication {

    public static void main(String[] args) {
        SpringApplication.run(MainApplication.class, args);
    }

}
