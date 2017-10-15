package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestProjectRole;
import com.ngtesting.platform.vo.ProjectRoleVo;

import java.util.List;

public interface ProjectRoleService extends BaseService {

	List list(Long orgId, String keywords, String disabled);

	TestProjectRole save(ProjectRoleVo vo, Long orgId);
	boolean delete(Long id);

	List<ProjectRoleVo> genVos(List<TestProjectRole> pos);
	ProjectRoleVo genVo(TestProjectRole role);

}
