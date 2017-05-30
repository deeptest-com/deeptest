package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestOrgRole;
import com.ngtesting.platform.entity.TestRole;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

public interface RoleService extends BaseService {

	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	TestRole save(RoleVo vo, Long orgId);
	boolean delete(Long id);
	boolean disable(Long id);

	List<RoleVo> genVos(List<TestRole> pos);
	RoleVo genVo(TestRole role);

}
