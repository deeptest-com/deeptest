package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestProjectRoleForOrg;
import com.ngtesting.platform.vo.ProjectRoleVo;

import java.util.List;

public interface ProjectRoleService extends BaseService {

	List list(Long orgId, String keywords, String disabled);

	TestProjectRoleForOrg save(ProjectRoleVo vo, Long orgId);
	boolean delete(Long id);

//	TestProjectRoleForOrg createDefaultBasicDataPers(Long orgId);

    List<ProjectRoleVo> genVos(List<TestProjectRoleForOrg> pos);
	ProjectRoleVo genVo(TestProjectRoleForOrg role);

}
