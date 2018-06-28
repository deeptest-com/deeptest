package com.ngtesting.platform.servlet;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.config.Constants;
import com.ngtesting.platform.utils.WebUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

@Component
public class AuthInterceptor implements HandlerInterceptor {
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
