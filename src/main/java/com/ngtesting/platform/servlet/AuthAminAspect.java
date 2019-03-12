package com.ngtesting.platform.servlet;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.AuthService;
import org.apache.shiro.SecurityUtils;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.AfterReturning;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.Map;

@Aspect
@Component
public class AuthAminAspect extends AuthAspectBase {
    @Autowired
    AuthService authService;
    @Autowired
    AuthDao authDao;

    @Pointcut("execution(public * com.ngtesting.platform.action.admin..*(..))")
    public void authAmin(){}

    @Before("authAmin()")
    public void doBefore(JoinPoint joinPoint){
        MethodSignature signature = (MethodSignature) joinPoint.getSignature();
        PrivCommon authAnnotation = signature.getMethod().getAnnotation(PrivCommon.class);

        Map<String, Object> map = getParam(joinPoint);
        JSONObject json = (JSONObject) map.get("json");

        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

        if (authAnnotation != null && authAnnotation.check() != null
                && authAnnotation.check().equals("false")) { // 不需要权限
            log("NONE", signature, user, null, null);
            return;
        }

        Integer orgId = json.getInteger("orgId") != null? json.getInteger("orgId") : user.getDefaultOrgId();
        String perm = "org_org:*:" + orgId;
        String opt = "and";

        checkAndLog(perm, signature, user, orgId, opt);
    }

    @AfterReturning(returning = "ret", pointcut = "authAmin()")
    public void doAfterReturning(Object ret) throws Throwable {
        logger.info("Response: " + ret);
    }

}
