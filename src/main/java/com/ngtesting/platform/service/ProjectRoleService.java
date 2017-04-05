package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrgRole;
import com.ngtesting.platform.entity.SysProjectRole;
import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.vo.OrgRoleVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.ProjectRoleVo;
import com.ngtesting.platform.vo.RoleVo;

public interface ProjectRoleService extends BaseService {

	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysProjectRole save(ProjectRoleVo vo, Long orgId);
	boolean delete(Long id);

	List<ProjectRoleVo> genVos(List<SysProjectRole> pos);
	ProjectRoleVo genVo(SysProjectRole role);

}
