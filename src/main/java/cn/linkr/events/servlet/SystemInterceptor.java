package cn.linkr.events.servlet;

import java.util.Enumeration;
import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

import org.apache.commons.lang.StringUtils;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.SysUser;
import cn.linkr.events.service.ClientService;
import cn.linkr.events.service.UserService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.Constant;
import cn.linkr.events.util.SpringContextHolder;
import cn.linkr.events.util.WebUtils;
import cn.linkr.events.util.Constant.RespCode;

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
        
        System.out.println("===" + request.getSession().getId());
        
        if (handler.getClass().isAssignableFrom(HandlerMethod.class)) {
            // 方法上是否有身份验证注解
            AuthPassport authPassport = ((HandlerMethod) handler).getMethodAnnotation(AuthPassport.class);
            // 声明不验证权限
            if (authPassport != null && authPassport.validate() == false) { 
                return true;
            }
            
            // 已经登录
            if (request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY) != null
            		|| request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY) != null) {
            	return true;
            }

            // 根据不同package处理不同身份认证逻辑
            String packageName = ((HandlerMethod) handler).getBeanType().getPackage().getName();
            String token = request.getHeader("token");
            if (token == null) {
            	token = request.getParameter("token");
            }
            
			// client请求鉴权
			if (packageName.startsWith(Constant.API_PACKAGE_FOR_CLIENT)) {
				if (StringUtils.isNotBlank(token)) {
					// 登录验证
					ClientService clientService = SpringContextHolder.getBean(ClientService.class);
					EvtClient client = clientService.getByToken(token.trim());
					if (client != null) {
						request.getSession().setAttribute(Constant.HTTP_SESSION_CLIENT_KEY, client);
						return true;
					}
				}
				
				Map<String, Object> result = new HashMap<String, Object>();
				result.put("code", RespCode.NOT_LOGIN.getCode());
				result.put("msg", "not login");
				WebUtils.renderJson(response, JSON.toJSONString(result));
				return false;
			} else if (packageName.startsWith(Constant.API_PACKAGE_FOR_ADMIN)) { // admin请求鉴权
				if (StringUtils.isNotBlank(token)) {
					// 登录验证
					UserService userService = SpringContextHolder.getBean(UserService.class);
					SysUser user = userService.getByToken(token.trim());
					if (user != null) {
						request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userService.genVo(user));
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
