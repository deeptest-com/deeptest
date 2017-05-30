package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestOrgRole;
import com.ngtesting.platform.entity.TestProjectRole;
import com.ngtesting.platform.entity.TestRole;
import com.ngtesting.platform.vo.OrgRoleVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectRoleVo;
import com.ngtesting.platform.vo.RoleVo;

public interface ProjectRoleService extends BaseService {

	List list(Long orgId, String keywords, String disabled);
	
	TestProjectRole save(ProjectRoleVo vo, Long orgId);
	boolean delete(Long id);

	List<ProjectRoleVo> genVos(List<TestProjectRole> pos);
	ProjectRoleVo genVo(TestProjectRole role);

}
