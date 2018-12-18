package com.ngtesting.platform.config;

import org.mybatis.spring.annotation.MapperScan;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.transaction.annotation.EnableTransactionManagement;

import javax.annotation.PostConstruct;

@SpringBootApplication
@EnableTransactionManagement
@MapperScan("com.ngtesting.platform.dao")
@ComponentScan(basePackages={"com.ngtesting.platform"})
public class AppLaunch {
    Logger logger = LoggerFactory.getLogger(AppLaunch.class);

	public static void main(String[] args) {
		SpringApplication.run(AppLaunch.class, args);
	}

    @PostConstruct
    void setDefaultTimezone() {
//        TimeZone.setDefault(TimeZone.getTimeZone("UTC"));
//        System.out.println(new Date().toString());
    }

}
