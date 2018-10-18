package com.ngtesting.platform.servlet;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.OrgPrivilegeService;
import com.ngtesting.platform.service.PropService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.utils.AuthPassport;
import com.ngtesting.platform.utils.WebUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.util.StringUtils;
import org.springframework.web.method.HandlerMethod;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.HashMap;
import java.util.Map;

@Component
public class AuthInterceptor implements HandlerInterceptor {
    private Logger logger = LoggerFactory.getLogger(getClass());

    @Autowired
    private PropService propService;
    @Autowired
    private UserService userService;
    @Autowired
    private OrgPrivilegeService orgPrivilegeService;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {

        WebUtils.InitWebContext(request, propService.getWorkDir());

        if (handler.getClass().isAssignableFrom(HandlerMethod.class)) {
            // 方法上是否有身份验证注解
            AuthPassport authPassport = ((HandlerMethod) handler).getMethodAnnotation(AuthPassport.class);
            // 声明不验证权限
            if (authPassport != null && authPassport.validate() == false) {
                return true;
            }

            // 已经登录
            if (request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE) != null) {
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
                if (!StringUtils.isEmpty(token)) {
                    // 登录验证
//                    UserService userService = SpringContextHolder.getBean(UserService.class);

                    TstUser user = userService.getByToken(token.trim());
                    if (user != null) {
                        request.getSession().setAttribute(Constant.HTTP_SESSION_USER_PROFILE, user);
                        return true;
                    }
                }

                Map<String, Object> result = new HashMap<String, Object>();
                result.put("code", Constant.RespCode.NOT_LOGIN.getCode());
                result.put("msg", "not login");
                WebUtils.renderJson(response, JSON.toJSONString(result));
                return false;
            }
        }
        return true;
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
