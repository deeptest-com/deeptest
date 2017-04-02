package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.vo.GroupVo;
import com.ngtesting.platform.vo.Page;

public interface GroupService extends BaseService {

	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysOrgGroup save(GroupVo vo, Long orgId);
	boolean delete(Long id);
	boolean disable(Long id);

	List<GroupVo> genVos(List<SysOrgGroup> pos);
	GroupVo genVo(SysOrgGroup user);

}
