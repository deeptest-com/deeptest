package cn.mobiu.events.interceptor;

import java.util.Enumeration;
import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.lang.StringUtils;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.constants.Constant.RespCode;
import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.service.ClientService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.SpringContextHolder;
import cn.mobiu.events.util.WebUtils;

import com.alibaba.fastjson.JSON;

public class SystemInterceptor implements HandlerInterceptor {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        if (Constant.WEB_ROOT == null) {
            Constant.WEB_ROOT = request.getScheme() + "://" + request.getServerName() + ":" + request.getServerPort() + request.getContextPath() + "/";
        }
        if (Constant.WORK_DIR == null) {
            Constant.WORK_DIR = request.getSession().getServletContext().getRealPath("/");;
        }
        
        if (handler.getClass().isAssignableFrom(HandlerMethod.class)) {
            // 方法上是否有身份验证注解
            AuthPassport authPassport = ((HandlerMethod) handler).getMethodAnnotation(AuthPassport.class);
            // 声明不验证权限
            if (authPassport != null && authPassport.validate() == false) { 
                return true;
            }
            
            // 已经登录
            if (request.getSession(true).getAttribute(Constant.HTTP_SESSION_CLIENT_KEY) != null) {
            	return true;
            }

            // 根据不同package处理不同身份认证逻辑
            String packageName = ((HandlerMethod) handler).getBeanType().getPackage().getName();
            String token = request.getHeader("token");
            if (token == null) {
            	token = request.getParameter("token");
            }
            
			// app鉴权管理
			if (packageName.startsWith(Constant.API_PACKAGE_FOR_CLIENT)) {
				
				if (StringUtils.isNotBlank(token)) {
					// 登录验证
					ClientService clientService = SpringContextHolder.getBean(ClientService.class);
					EvtClient client = clientService.getByToken(token.trim());
					if (client != null) {
						request.getSession(true).setAttribute(Constant.HTTP_SESSION_CLIENT_KEY, client);
						return true;
					}
				}
				
				Map<String, Object> result = new HashMap<String, Object>();
				result.put("code", RespCode.NOT_LOGIN.getCode());
				result.put("msg", "not login");
				WebUtils.renderJson(response, JSON.toJSONString(result));
				return false;
			}
        }
        return true;
    }
    
    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, ModelAndView mv) throws Exception {
    }
    
    @Override
    public void afterCompletion(HttpServletRequest request, HttpServletResponse response, Object handler, Exception exception) throws Exception {
    }
}
