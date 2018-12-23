package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.PermissionDao;
import com.ngtesting.platform.service.intf.PermissionService;
import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class PermissionServiceImpl extends BaseServiceImpl implements PermissionService {
    private Logger logger = LoggerFactory.getLogger(getClass());

    @Autowired
    PermissionDao permissionDao;

    @Override
    public Boolean hasPerm(String scope, String[] perms, String opt, Integer userId, Integer orgId, HttpServletRequest request) {
        Map<String, Map<String, Boolean>> permsMap = getPermsMap(userId, request);

        Boolean pass = permsMap.get(scope) == null? false: checkPerm(perms, opt, permsMap.get(scope));

        // 基本都不用访问数据库，二次查询只会发生在：
        // 1. 权限有所更新
        // 2. 用户第一次鉴权
        // 3. 非法攻击（模拟了非法的orgId、prjId等请求参数）
        if (!pass) {
            permsMap = genPermsMap(userId, request);
            pass = checkPerm(perms, opt, permsMap.get(scope));
        }
        return pass;
    }

    private Boolean checkPerm(String[] perms, String opt, Map<String, Boolean> permsMap) {
        logger.info("AuthAspect Required = " + StringUtils.join(permsMap.keySet(), ","));

        if ("or".equals(opt)) {
            for (String p : perms) {
                if (permsMap.get(p) != null && permsMap.get(p)) {
                    return true;
                }
            }

            return false;
        } else {
            for (String p : perms) {
                if (permsMap.get(p) == null || permsMap.get(p) == false) {
                    return false;
                }
            }

            return true;
        }
    }

    private Map<String, Map<String, Boolean>> getPermsMap(Integer userId, HttpServletRequest request) {
        Map<String, Map<String, Boolean>> permsMap = new HashMap<>();

        Object obj = request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PERMISSION);
        if (obj == null) {
            permsMap = genPermsMap(userId, request);
        }

        return permsMap;
    }

    private Map<String, Map<String, Boolean>> genPermsMap(Integer userId, HttpServletRequest request) {
        Map<String, Map<String, Boolean>> permsMap = new HashMap<>();

        permsMap.put("org", genOrgPermsMap(userId));
        permsMap.put("prj", genPrjPermsMap(userId));

        request.getSession().setAttribute(Constant.HTTP_SESSION_USER_PERMISSION, permsMap);
        return permsMap;
    }

    private Map<String, Boolean> genOrgPermsMap(Integer userId) {
        List<Map> ls = permissionDao.listOrgPermission(userId);

        Map<String, Boolean> ret = new HashMap<>();
        for (Map map : ls) {
            ret.put(map.get("code").toString(), true);
        }

        return ret;
    }

    private Map<String, Boolean> genPrjPermsMap(Integer userId) {
        List<Map> ls = permissionDao.listPrjPermission(userId);

        Map<String, Boolean> ret = new HashMap<>();
        for (Map map : ls) {
            ret.put(map.get("code").toString(), true);
        }

        return ret;
    }
}
