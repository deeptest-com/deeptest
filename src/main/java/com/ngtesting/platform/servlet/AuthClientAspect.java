package com.ngtesting.platform.servlet;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.PermissionService;
import org.apache.commons.lang3.StringUtils;
import org.apache.shiro.SecurityUtils;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.aspectj.lang.annotation.Pointcut;
import org.aspectj.lang.reflect.MethodSignature;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import javax.servlet.http.HttpServletRequest;
import java.util.ArrayList;
import java.util.List;
import java.util.Map;

@Aspect
@Component
public class AuthClientAspect extends AuthAspectBase {
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
        beforeCheck(joinPoint, "project");
    }

    public void beforeCheck(JoinPoint joinPoint, String scope){
        TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();

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

        Integer id;

        if (src.equals("request")) {
            id = json.getInteger(key);
        } else if (src.equals("session")) {
            id = scope.equals("org")? user.getDefaultOrgId(): user.getDefaultPrjId();
        } else { // 目前都未指定
            id = json.getInteger(key) != null?
                    json.getInteger(key) :
                        scope.equals("org")? user.getDefaultOrgId(): user.getDefaultPrjId();
        }

        String perm = "";
        List<String> ls = new ArrayList<>();
        if (perms.length > 0) {
            for (String p : perms) {
                ls.add(p + ":" + id);
            }
            perm = StringUtils.join(ls, ",");
        }

        checkAndLog(perm, signature, user, id, opt);
    }
}