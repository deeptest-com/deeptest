package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstProjectPrivilegeDefine;

import java.util.List;
import java.util.Map;

public interface ProjectPrivilegeService extends BaseService {

    Map<String, Map<String, TstProjectPrivilegeDefine>> listPrivilegesByOrgAndProjectRole(Integer orgId, Integer projectRoleId);

    boolean addUserAsProjectTestLeaderPers(Integer orgId, Integer projectId, String roleCode, Integer userId);

    boolean saveProjectPrivileges(Integer orgId, Integer roleId, Map<String, List<TstProjectPrivilegeDefine>> map);

	Map<String, Boolean> listByUser(Integer userId, Integer prjId, Integer orgId);
}
