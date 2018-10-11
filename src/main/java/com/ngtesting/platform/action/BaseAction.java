package com.ngtesting.platform.action;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.AuthDao;
import com.ngtesting.platform.service.AuthService;
import org.springframework.beans.factory.annotation.Autowired;

import java.util.HashMap;
import java.util.Map;

public class BaseAction {
    @Autowired
    AuthService authService;
    @Autowired
    AuthDao authDao;

	public Boolean hasNoOrgAdminPriviledge(Integer userId, Integer orgId) {
		return !authService.hasOrgAdminPrivilege(userId, orgId);
	}

	public Boolean userNotInOrg(Integer userId, Integer orgId) {
		return authDao.userNotInOrg(userId, orgId);
	}

	public Boolean userNotInProject(Integer userId, Integer projectId) {
		return authDao.userNotInProject(userId, projectId);
	}

	public Map<String, Object> authFail() {
		Map<String, Object> ret = new HashMap<String, Object>();
		ret.put("code", Constant.RespCode.AUTH_FAIL.getCode());
		ret.put("msg", "权限不足");
		return ret;
	}

    public Map<String, Object> bizFail() {
        Map<String, Object> ret = new HashMap<String, Object>();
        ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
        ret.put("msg", "业务处理错误");
        return ret;
    }

}
