package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysOrgGroup;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.Page;

public interface OrgGroupService extends BaseService {

	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	
	SysOrgGroup save(OrgGroupVo vo, Long orgId);
	boolean delete(Long id);

	List<OrgGroupVo> genVos(List<SysOrgGroup> pos);
	OrgGroupVo genVo(SysOrgGroup user);

}
