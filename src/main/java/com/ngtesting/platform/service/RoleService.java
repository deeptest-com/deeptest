package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysRole;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RoleVo;

public interface RoleService extends BaseService {

	Page listByPage(Long companyId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysRole save(RoleVo vo, Long companyId);
	boolean delete(Long id);
	boolean disable(Long id);

	List<RoleVo> genVos(List<SysRole> pos);
	RoleVo genVo(SysRole role);

}
