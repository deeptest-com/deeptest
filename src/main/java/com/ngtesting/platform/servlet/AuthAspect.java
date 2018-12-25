package com.ngtesting.platform.servlet;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.exception.AuthException;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.PermissionService;
import org.apache.commons.lang3.StringUtils;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
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
    private Logger logger = LoggerFactory.getLogger(getClass());

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

        String key;
        String src;
        String[] perms;
        String opt;
        String classAndMethod = signature.getMethod().toString();

        if (scope.equals("org")) {
            PrivOrg authAnnotation = signature.getMethod().getAnnotation(PrivOrg.class);
            key = authAnnotation.key();
            src = authAnnotation.src();
            perms = authAnnotation.perms();
            opt = authAnnotation.opt();
        } else {
            PrivPrj authAnnotation = signature.getMethod().getAnnotation(PrivPrj.class);
            key = authAnnotation.key();
            src = authAnnotation.src();
            perms = authAnnotation.perms();
            opt = authAnnotation.opt();
        }

        Map<String, Object> map = getParam(joinPoint);
        HttpServletRequest request = (HttpServletRequest)map.get("request");
        JSONObject json = (JSONObject)map.get("json");

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer id;

        if (src.equals("request")) {
            id = json.getInteger(key);
        } else if (src.equals("session")) {
            id = scope.equals("org")? user.getDefaultOrgId(): user.getDefaultPrjId();
        } else { // 未指定
            id = json.getInteger(key) != null?
                    json.getInteger(key) :
                        scope.equals("org")? user.getDefaultOrgId(): user.getDefaultPrjId();
        }

        Boolean pass;
        if (perms.length == 0) { // 查看权限
            pass = permissionService.viewPerm(scope, opt, user.getId(), id, request);
        } else {
            pass = permissionService.hasPerm(scope, perms, opt, user.getId(), id, request);
        }

        log(perms, pass, classAndMethod, user, id, opt);

        if (!pass) {
            throw new AuthException();
        }
    }

    @Before("execution(* com.ngtesting.platform.action.admin.*.*(..))")
    public void before(JoinPoint joinPoint){
        MethodSignature signature = (MethodSignature) joinPoint.getSignature();
        PrivCommon authAnnotation = signature.getMethod().getAnnotation(PrivCommon.class);

        Map<String, Object> map = getParam(joinPoint);
        HttpServletRequest request = (HttpServletRequest)map.get("request");
        JSONObject json = (JSONObject)map.get("json");

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        String classAndMethod = signature.getMethod().toString();
        if (authAnnotation != null && authAnnotation.check() != null && authAnnotation.check().equals("false")) {
            String perms[] = { "@PrivCommon(check=\"false\")" };
            log(perms, true, classAndMethod, user, null, null);

            return;
        }

        String perms[] = { "org-admin" };
        String opt = "or";

        Integer orgId = json.getInteger("orgId") != null? json.getInteger("orgId") : user.getDefaultOrgId();

        Boolean pass = permissionService.hasPerm("org", perms, opt, user.getId(), orgId, request);

        log(perms, pass, classAndMethod, user, orgId, opt);

        if (!pass) {
            throw new AuthException();
        }
    }

    private Map<String, Object> getParam(JoinPoint joinPoint){
        Map<String, Object> map = new HashMap<>();

        MethodSignature methodSignature = (MethodSignature) joinPoint.getSignature();
        Method method = methodSignature.getMethod();

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

    private void log(String[] perms, Boolean pass, String classAndMethod, TstUser user, Integer orgId, String opt){

        logger.info("AuthAspect Require  = " + (perms.length == 0?"view":StringUtils.join(perms, ",")));
        logger.info("AuthAspect Result   = " + pass);

        logger.info("AuthAspect Detail   - " + classAndMethod);

        String info = "                      " + "user: " + user.getId() + ", " + user.getEmail();

        if (orgId != null) {
            info += ", orgId = " + orgId;
        }
        if (opt != null) {
            info += ", opt: " + opt;
        }

        logger.info(info);
    }

}