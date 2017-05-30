package com.ngtesting.platform.service;

import java.util.List;
import java.util.Map;

import com.ngtesting.platform.entity.TestOrgGroup;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.OrgPrivilegeVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectPrivilegeVo;

public interface ProjectPrivilegeService extends BaseService {

	Map<String, List<ProjectPrivilegeVo>> listPrivilegesByOrg(Long orgId, Long projectRoleId);

	boolean saveProjectPrivileges(Long roleId, Map<String, List<ProjectPrivilegeVo>> map);

}
