package com.ngtesting.platform.config;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.exception.AuthException;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AuthService;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.AfterReturning;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;
import java.lang.annotation.Annotation;

@Aspect
@Component
public class AuthAminAspect {
    final static Logger logger  =  LoggerFactory.getLogger(AuthAminAspect.class );

    @Autowired
    AuthService authService;
    @Autowired
    AuthDao authDao;

    @Pointcut("execution(public * com.ngtesting.platform.action.admin..*(..))")
    public void authAmin(){}

    @Before("authAmin()")
    public void doBefore(JoinPoint joinPoint) throws Throwable {
        Boolean hasOrgAdminPriviledge = false;

        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        HttpServletRequest request = attributes.getRequest();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer orgId = getOrgId(joinPoint);
        if (orgId != null) {
            hasOrgAdminPriviledge = authService.hasOrgAdminPrivilege(user.getId(), orgId);
        } else {
            hasOrgAdminPriviledge = true;
        }

        if (!hasOrgAdminPriviledge) {
            throw new AuthException();
        }

//        logger.info("URL : " + request.getRequestURL().toString());
//        logger.info("HTTP_METHOD : " + request.getMethod());
//        logger.info("IP : " + request.getRemoteAddr());
//        logger.info("CLASS_METHOD : " + joinPoint.getSignature().getDeclaringTypeName() + "." + joinPoint.getSignature().getName());
//        logger.info("ARGS : " + Arrays.toString(joinPoint.getArgs()));
    }

    @AfterReturning(returning = "ret", pointcut = "authAmin()")
    public void doAfterReturning(Object ret) throws Throwable {
        logger.info("Response: " + ret);
    }

    Integer getOrgId(JoinPoint joinPoint) throws Throwable {
        Integer orgId = null;

        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        HttpServletRequest request = attributes.getRequest();

        MethodSignature methodSignature = (MethodSignature) joinPoint.getSignature();
        Annotation[][] annotationMatrix = methodSignature.getMethod().getParameterAnnotations();
        int index = -1;
        for (Annotation[] annotations : annotationMatrix) {
            index++;
            for (Annotation annotation : annotations) {
                if (annotation instanceof RequestBody) {
                    Object requestBody = joinPoint.getArgs()[index];
                    if (requestBody instanceof JSONObject) {
                        logger.info("RequestBody: " + requestBody);

                        JSONObject json = (JSONObject) requestBody;
                        if (json.getInteger("orgId") != null) {
                            orgId = json.getInteger("orgId");
                        }
                    } else if (requestBody instanceof TstOrg) {
                        TstOrg json = (TstOrg) requestBody;
                        orgId = json.getId();
                    }
                }
            }
        }

        if (orgId == null) {
            TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
            logger.info("TstUser: " + user);

            orgId = user.getDefaultOrgId();
        }

        logger.info("***orgId: " + orgId);
        return orgId;
    }

}
