package com.ngtesting.platform.service.intf;

import javax.servlet.http.HttpServletRequest;

public interface PermissionService extends BaseService {

    Boolean hasOrgPerm(String scope, String[] perms, String opt,
                       Integer userId, Integer orgId,
                       HttpServletRequest request);
}
