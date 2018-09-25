package com.ngtesting.platform.config;

import com.ngtesting.platform.servlet.CustomMultipartResolver;
import org.mybatis.spring.annotation.MapperScan;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.web.servlet.MultipartAutoConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.transaction.annotation.EnableTransactionManagement;
import org.springframework.web.multipart.MultipartResolver;

import javax.annotation.PostConstruct;

@EnableAutoConfiguration(exclude = {MultipartAutoConfiguration.class})

@SpringBootApplication
@EnableTransactionManagement
@MapperScan("com.ngtesting.platform.dao")
@ComponentScan(basePackages={"com.ngtesting.platform"})
public class AppLaunch {
    Logger logger = LoggerFactory.getLogger(AppLaunch.class);

	public static void main(String[] args) {
		SpringApplication.run(AppLaunch.class, args);
	}

    @Bean(name = "multipartResolver")
    public MultipartResolver multipartResolver() {
        CustomMultipartResolver customMultipartResolver = new CustomMultipartResolver();
        return customMultipartResolver;
    }

    @PostConstruct
    void setDefaultTimezone() {
//        TimeZone.setDefault(TimeZone.getTimeZone("UTC"));
//        System.out.println(new Date().toString());
    }

}
