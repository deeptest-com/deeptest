package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestProjectPrivilegeDefine;
import com.ngtesting.platform.entity.TestProjectRolePriviledgeRelation;
import com.ngtesting.platform.vo.ProjectPrivilegeDefineVo;

import java.util.List;
import java.util.Map;

public interface ProjectPrivilegeService extends BaseService {

	Map<String, List<ProjectPrivilegeDefineVo>> listPrivilegesByOrgAndProjectRole(Long orgId, Long projectRoleId);
    List<TestProjectRolePriviledgeRelation> listProjectRolePrivileges(Long projectRoleId);

    List<TestProjectPrivilegeDefine> listAllProjectPrivileges();

    boolean saveProjectPrivileges(Long roleId, Map<String, List<ProjectPrivilegeDefineVo>> map);

	Map<String, Boolean> listByUserPers(Long userId, Long prjId, Long orgId);
}
