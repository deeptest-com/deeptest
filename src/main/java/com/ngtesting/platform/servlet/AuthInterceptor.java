package com.ngtesting.platform.servlet;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.intf.PropService;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.utils.WebUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

@Component
public class AuthInterceptor implements HandlerInterceptor {
    private Logger logger = LoggerFactory.getLogger(getClass());

    @Autowired
    private PropService propService;
    @Autowired
    private UserService userService;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        if (Constant.WEB_ROOT == null) {
            WebUtils.InitWebContext(request, propService.getWorkDir());
        }

//        if (handler.getClass().isAssignableFrom(HandlerMethod.class)) {
//            AuthPassport authPassport = ((HandlerMethod) handler).getMethodAnnotation(AuthPassport.class);
//            // 声明无需权限，返回
//            if (authPassport != null && authPassport.validate() == false) {
//                return true;
//            }
//
//            // 已经登录，返回
//            if (request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE) != null) {
//                return true;
//            }
//
//            String token = request.getHeader("token");
//            if (token == null) {
//                token = request.getParameter("token");
//            }
//
//            String packageName = ((HandlerMethod) handler).getBeanType().getPackage().getName();
//            // client请求鉴权
//            if (packageName.startsWith(Constant.API_PACKAGE_FOR_CLIENT)) {
//                if (!StringUtils.isEmpty(token)) {
//                    TstUser user = userService.getByToken(token.trim());
//                    // token对应上用户，返回
//                    if (user != null) {
//                        request.getSession().setAttribute(Constant.HTTP_SESSION_USER_PROFILE, user);
//                        return true;
//                    }
//                }
//
//                Map<String, Object> result = new HashMap<String, Object>();
//                result.put("code", Constant.RespCode.NOT_LOGIN.getCode());
//                result.put("msg", "not login");
//                WebUtils.renderJson(response, JSON.toJSONString(result));
//                return false;
//            }
//        }
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
