package com.ngtesting.platform.config;

import com.ngtesting.platform.servlet.AuthInterceptor;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;

@Configuration
public class InterceptorConfig extends WebMvcConfigurerAdapter {

    @Bean
    public AuthInterceptor getAuthInterceptor() {
        return new AuthInterceptor();
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // 多个拦截器组成一个拦截器链
        // addPathPatterns 用于添加拦截规则
        // excludePathPatterns 用户排除拦截
        registry.addInterceptor(getAuthInterceptor()).addPathPatterns("/**");
        // registry.addInterceptor(new InterfaceAuthCheckInterceptor()).addPathPatterns("/api/**");
        // 如果interceptor中不注入redis或其他项目可以直接new，否则请使用上面这种方式
        super.addInterceptors(registry);
    }
}
