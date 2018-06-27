package com.ngtesting.platform.config;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.ngtesting.platform.dao")
public class NgtestingWeb {

	public static void main(String[] args) {
		SpringApplication.run(NgtestingWeb.class, args);
	}

}
