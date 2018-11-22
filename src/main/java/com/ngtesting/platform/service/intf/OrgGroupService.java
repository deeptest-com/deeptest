package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstOrgGroup;

import java.util.List;

public interface OrgGroupService extends BaseService {

	List<TstOrgGroup> listByPage(Integer orgId, String keywords, Boolean disabled, Integer currentPage, Integer itemsPerPage);
	List search(Integer orgId, String keywords,  List<Integer> exceptIds);
	List<TstOrgGroup> list(Integer orgId);
	TstOrgGroup get(Integer id, Integer orgId);
	TstOrgGroup save(TstOrgGroup vo, Integer orgId);
	boolean delete(Integer id, Integer orgId);

//	void initDefaultBasicDataPers(TestOrg org);

}
