package com.ngtesting.platform.servlet;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.exception.AuthException;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.PermissionService;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.servlet.http.HttpServletRequest;
import java.lang.reflect.Method;
import java.lang.reflect.Parameter;
import java.util.HashMap;
import java.util.Map;

@Aspect
@Component
public class AuthAspect {
    @Autowired
    PermissionService permissionService;

    @Pointcut("@annotation(com.ngtesting.platform.servlet.PrivOrg)")
    public void orgPointCut(){}

    @Pointcut("@annotation(com.ngtesting.platform.servlet.PrivPrj)")
    public void prjPointCut(){}

    @Before("orgPointCut()")
    public void beforeOrg(JoinPoint joinPoint){
        beforeCheck(joinPoint, "org");
    }

    @Before("prjPointCut()")
    public void beforePrj(JoinPoint joinPoint){
        beforeCheck(joinPoint, "prj");
    }

    public void beforeCheck(JoinPoint joinPoint, String scope){
        MethodSignature signature = (MethodSignature) joinPoint.getSignature();
        PrivOrg authAnnotation = signature.getMethod().getAnnotation(PrivOrg.class);

        Map<String, Object> map = getParam(joinPoint);
        HttpServletRequest request = (HttpServletRequest)map.get("request");
        JSONObject json = (JSONObject)map.get("json");

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = authAnnotation.src().equals("session")? user.getDefaultOrgId(): json.getInteger(authAnnotation.key());

        String[] perms = authAnnotation.perms();
        String opt = authAnnotation.opt();

        Boolean pass = permissionService.hasOrgPerm(scope, perms, opt, user.getId(), orgId, request);

        if (!pass) {
            throw new AuthException();
        }
    }

//    @Before("execution(* com.ngtesting.platform.service.impl.UserServiceImpl.list(..))")
//    public void before(JoinPoint joinPoint){
//        MethodSignature signature = (MethodSignature) joinPoint.getSignature();
//        Method method = signature.getMethod();
//        System.out.println("方法规则式拦截,"+method.getName());
//    }

    private Map<String, Object> getParam(JoinPoint joinPoint){
        Map<String, Object> map = new HashMap<>();

        MethodSignature methodSignature = (MethodSignature) joinPoint.getSignature();
        Method method = methodSignature.getMethod();
        PrivOrg authAnnotation = method.getAnnotation(PrivOrg.class);

        Parameter[] params =  method.getParameters();

        for (int i = 0; i < params.length; i++) {
            Parameter p = params[i];

            if (p.getType().getSimpleName().equals("HttpServletRequest")) {
                map.put("request", joinPoint.getArgs()[i]);
            }
            if (p.getType().getSimpleName().equals("JSONObject")) {
                map.put("json", joinPoint.getArgs()[i]);
            }
        }

        return map;
    }

}