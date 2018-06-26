package com.ngtesting.platform;

import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
@MapperScan("com.ngtesting.platform.dao")
public class NgtestingWebApplication {

	public static void main(String[] args) {
		SpringApplication.run(NgtestingWebApplication.class, args);
	}

}
