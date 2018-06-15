package com.ngtesting.platform.servlet;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.SpringContextHolder;
import com.ngtesting.platform.util.WebUtils;
import com.ngtesting.platform.vo.UserVo;
import org.apache.commons.lang.StringUtils;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

public class SystemInterceptor implements HandlerInterceptor {

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {

        if (Constant.WEB_ROOT == null) {
            Constant.WEB_ROOT = request.getScheme() + "://" + request.getServerName() +
                    (request.getServerPort() != 80? ":" + request.getServerPort() : "")
                    + request.getContextPath() + "/";
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
            if (request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY) != null
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
					AccountService accountService = SpringContextHolder.getBean(AccountService.class);
					UserService userService = SpringContextHolder.getBean(UserService.class);

					TestUser user = accountService.getByToken(token.trim());
                    UserVo userVo = userService.genVo(user);
                    if (user != null) {
						request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userVo);
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
