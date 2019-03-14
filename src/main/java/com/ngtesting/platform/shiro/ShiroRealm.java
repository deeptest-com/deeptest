package com.ngtesting.platform.shiro;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.impl.IssueOptServiceImpl;
import com.ngtesting.platform.service.intf.PermissionService;
import com.ngtesting.platform.service.intf.UserService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.apache.shiro.SecurityUtils;
import org.apache.shiro.authc.*;
import org.apache.shiro.authz.AuthorizationInfo;
import org.apache.shiro.authz.SimpleAuthorizationInfo;
import org.apache.shiro.realm.AuthorizingRealm;
import org.apache.shiro.subject.PrincipalCollection;
import org.apache.shiro.subject.Subject;
import org.apache.shiro.util.ByteSource;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.List;

public class ShiroRealm extends AuthorizingRealm {
    Log logger = LogFactory.getLog(IssueOptServiceImpl.class);

    @Autowired
    private UserService userService;
    @Autowired
    private PermissionService permissionService;

    /**
     * 认证信息.(身份验证) : Authentication 是用来验证用户身份
     *
     */
    @Override
    protected AuthenticationInfo doGetAuthenticationInfo(AuthenticationToken authcToken) throws AuthenticationException {
        UsernamePasswordToken token = (UsernamePasswordToken) authcToken;
        String email = token.getUsername();

        // 从数据库获取对应用户名密码的用户
        TstUser user = userService.getByEmail(email);
        if (user != null) {
            // 用户为禁用状态
            if (user.getLocked()) {
                throw new DisabledAccountException();
            }

            Subject subject = SecurityUtils.getSubject();
            String sessionId = subject.getSession().getId().toString();
            user.setToken(sessionId);

            SimpleAuthenticationInfo authenticationInfo = new SimpleAuthenticationInfo(
                    user,
                    user.getPassword(),
                    ByteSource.Util.bytes(user.getSalt()),
                    getName()
            );
            return authenticationInfo;
        }
        throw new UnknownAccountException();
    }

    /**
     * 授权
     */
    @Override
    protected AuthorizationInfo doGetAuthorizationInfo(PrincipalCollection principals) {
        Object principal = principals.getPrimaryPrincipal();
        SimpleAuthorizationInfo info = new SimpleAuthorizationInfo();

        if (principal instanceof TstUser) {
            TstUser user = (TstUser) principal;
            if(user != null){
                // info.addRole(role.getEnname());

//                Map<String, String> map = new HashMap<>();
                List<String> list = permissionService.getShiroStylePermissions(user.getId());
                for (String perm : list){
                    System.out.println(perm);
                    info.addStringPermission(perm);
//                    if (!perm.startsWith("org_")) {
//                        String id = perm.split(":")[2];
//                        if (!map.containsKey(id)) {
//                            info.addStringPermission("project:view:" + id); // add project view perm
//                            map.put(id, "");
//                        }
//                    }
                }
            }
        }

        return info;
    }
}
