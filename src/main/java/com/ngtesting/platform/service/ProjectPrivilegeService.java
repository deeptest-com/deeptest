package com.ngtesting.platform.service;

import com.ngtesting.platform.vo.ProjectPrivilegeVo;

import java.util.List;
import java.util.Map;

public interface ProjectPrivilegeService extends BaseService {

	Map<String, List<ProjectPrivilegeVo>> listPrivilegesByOrg(Long orgId, Long projectRoleId);

	boolean saveProjectPrivileges(Long roleId, Map<String, List<ProjectPrivilegeVo>> map);

	Map<String, Boolean> listByUser(Long userId, Long prjId);
}
