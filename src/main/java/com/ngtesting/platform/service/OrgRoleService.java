package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestOrgRole;
import com.ngtesting.platform.entity.TestRole;
import com.ngtesting.platform.vo.OrgRoleVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

public interface OrgRoleService extends BaseService {

	List list(Long orgId, String keywords, String disabled);
	
	TestOrgRole save(OrgRoleVo vo, Long orgId);
	boolean delete(Long id);

	List<OrgRoleVo> genVos(List<TestOrgRole> pos);
	OrgRoleVo genVo(TestOrgRole role);

}
