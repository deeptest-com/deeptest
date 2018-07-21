package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.vo.Page;

import java.util.List;

public interface OrgGroupService extends BaseService {

	Page listByPage(Integer orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	List search(Integer orgId, String keywords, String exceptIds);

	TstOrgGroup save(TstOrgGroup vo, Integer orgId);
	boolean delete(Integer id);

//	void initDefaultBasicDataPers(TestOrg org);

	List<TstOrgGroup> genVos(List<TstOrgGroup> pos);
	TstOrgGroup genVo(TstOrgGroup user);

}
