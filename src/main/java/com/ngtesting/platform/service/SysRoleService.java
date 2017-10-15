package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

import java.util.List;

public interface SysRoleService extends BaseService {

	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);

	SysRole save(RoleVo vo, Long orgId);
	boolean delete(Long id);
	boolean disable(Long id);

	List<RoleVo> genVos(List<SysRole> pos);
	RoleVo genVo(SysRole role);

}
