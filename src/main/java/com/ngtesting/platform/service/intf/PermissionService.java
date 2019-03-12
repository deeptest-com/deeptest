package com.ngtesting.platform.service.intf;

import java.util.List;

public interface PermissionService extends BaseService {
    List<String> getShiroStylePermissions(Integer id);

//    Boolean hasPerm(String scope, String[] perms, String opt,
//                    Integer userId, Integer id,
//                    HttpServletRequest request);
//
//    Boolean viewPerm(String scope, String opt, Integer userId, Integer id, HttpServletRequest request);
}
