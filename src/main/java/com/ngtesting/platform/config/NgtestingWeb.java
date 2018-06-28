package com.ngtesting.platform.config;

import com.ngtesting.platform.servlet.CustomMultipartResolver;
import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.web.servlet.MultipartAutoConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.web.multipart.MultipartResolver;

@EnableAutoConfiguration(exclude = {MultipartAutoConfiguration.class})

@SpringBootApplication
@MapperScan("com.ngtesting.platform.dao")
public class NgtestingWeb {

	public static void main(String[] args) {
		SpringApplication.run(NgtestingWeb.class, args);
	}

    @Bean(name = "multipartResolver")
    public MultipartResolver multipartResolver()
    {
        CustomMultipartResolver customMultipartResolver = new CustomMultipartResolver();
        return customMultipartResolver;
    }

}
