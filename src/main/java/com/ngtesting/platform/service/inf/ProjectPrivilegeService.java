package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import com.ngtesting.platform.model.TstProjectRolePriviledgeRelation;

import java.util.List;
import java.util.Map;

public interface ProjectPrivilegeService extends BaseService {

    Map<String, Map<String, TstProjectPrivilegeDefine>> listPrivilegesByOrgAndProjectRole(Integer orgId, Integer projectRoleId);
    List<TstProjectRolePriviledgeRelation> listProjectRolePrivileges(Integer projectRoleId);

    List<TstProjectPrivilegeDefine> listAllProjectPrivileges();

    boolean addUserAsProjectTestLeaderPers(Integer orgId, Integer projectId, String roleCode, Integer userId);

    boolean saveProjectPrivileges(Integer roleId, Map<String, List<TstProjectPrivilegeDefine>> map);

	Map<String, Boolean> listByUserPers(Integer userId, Integer prjId, Integer orgId);
}
