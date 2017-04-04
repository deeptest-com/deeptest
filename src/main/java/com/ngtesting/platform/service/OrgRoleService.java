package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrgRole;
import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.vo.OrgRoleVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

public interface OrgRoleService extends BaseService {

	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysOrgRole save(OrgRoleVo vo, Long orgId);
	boolean delete(Long id);
	boolean disable(Long id);

	List<OrgRoleVo> genVos(List<SysOrgRole> pos);
	OrgRoleVo genVo(SysOrgRole role);

}
