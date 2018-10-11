package com.ngtesting.platform.config;

import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.exception.AuthException;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AuthService;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.AfterReturning;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;

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
        String className = joinPoint.getSignature().getDeclaringType().getSimpleName();
        if (className.equals("OrgAdmin")) {
            return;
        }

        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        HttpServletRequest request = attributes.getRequest();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer orgId = getOrgId(joinPoint);
        if (orgId != null && !authService.hasOrgAdminPrivilege(user.getId(), orgId)) {
            throw new AuthException();
        }
    }

    @AfterReturning(returning = "ret", pointcut = "authAmin()")
    public void doAfterReturning(Object ret) throws Throwable {
        logger.info("Response: " + ret);
    }

    Integer getOrgId(JoinPoint joinPoint) throws Throwable {
        ServletRequestAttributes attributes = (ServletRequestAttributes) RequestContextHolder.getRequestAttributes();
        HttpServletRequest request = attributes.getRequest();

//        logger.info("URL : " + request.getRequestURL().toString());
//        logger.info("HTTP_METHOD : " + request.getMethod());
//        logger.info("IP : " + request.getRemoteAddr());
//        logger.info("CLASS_METHOD : " + joinPoint.getSignature().getDeclaringTypeName() + "." + joinPoint.getSignature().getName());
//        logger.info("ARGS : " + Arrays.toString(joinPoint.getArgs()));

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer orgId = user.getDefaultOrgId();
        return orgId;
    }

//        MethodSignature methodSignature = (MethodSignature) joinPoint.getSignature();
//        Annotation[][] annotationMatrix = methodSignature.getMethod().getParameterAnnotations();
//        int index = -1;
//        for (Annotation[] annotations : annotationMatrix) {
//            index++;
//            for (Annotation annotation : annotations) {
//                if (annotation instanceof RequestBody) {
//                    Object requestBody = joinPoint.getArgs()[index];
//                    if (requestBody instanceof JSONObject) {
//                        logger.info("RequestBody: " + requestBody);
//                    }
//                }
//            }
//        }

}
