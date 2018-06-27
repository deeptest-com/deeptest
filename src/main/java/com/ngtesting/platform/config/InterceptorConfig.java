package com.ngtesting.platform.config;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.utils.Constants;
import com.ngtesting.platform.utils.WebUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.util.StringUtils;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurerAdapter;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

@Configuration
public class InterceptorConfig extends WebMvcConfigurerAdapter {

    @Bean
    public InterfaceAuthCheckInterceptor getInterfaceAuthCheckInterceptor() {
        return new InterfaceAuthCheckInterceptor();
    }

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        // 多个拦截器组成一个拦截器链
        // addPathPatterns 用于添加拦截规则
        // excludePathPatterns 用户排除拦截
        registry.addInterceptor(getInterfaceAuthCheckInterceptor()).addPathPatterns("/**");
        // registry.addInterceptor(new InterfaceAuthCheckInterceptor()).addPathPatterns("/api/**");
        // 如果interceptor中不注入redis或其他项目可以直接new，否则请使用上面这种方式
        super.addInterceptors(registry);
    }

    class InterfaceAuthCheckInterceptor implements HandlerInterceptor {

        private Logger logger = LoggerFactory.getLogger(getClass());

        @Override
        public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object obj)
                throws Exception {
            String key = request.getParameter("key");
            if (StringUtils.isEmpty(key)) {
                response.setContentType("application/json;charset=utf-8");

                Map<String, Object> result = new HashMap();
                result.put("code", Constants.RespCode.NOT_LOGIN.getCode());
                result.put("msg", "not login");
                WebUtils.renderJson(response, JSON.toJSONString(result));
                return false;
            } else {
                return true;
            }
        }

        @Override
        public void afterCompletion(HttpServletRequest arg0, HttpServletResponse arg1, Object arg2, Exception arg3)
                throws Exception {

        }

        @Override
        public void postHandle(HttpServletRequest arg0, HttpServletResponse arg1, Object arg2, ModelAndView arg3)
                throws Exception {

        }

    }
}
