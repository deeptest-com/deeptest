package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgGroup;

import java.util.List;

public interface OrgGroupService extends BaseService {

	List<TstOrgGroup> listByPage(Integer orgId, String keywords, Boolean disabled, Integer currentPage, Integer itemsPerPage);
	List search(Integer orgId, String keywords, String exceptIds);
	List<TstOrgGroup> list(Integer orgId);
	TstOrgGroup get(Integer id);
	TstOrgGroup save(TstOrgGroup vo, Integer orgId);
	boolean delete(Integer id);

//	void initDefaultBasicDataPers(TestOrg org);

}
