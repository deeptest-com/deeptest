package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONArray;
import com.ngtesting.platform.entity.TestOrgGroup;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.Page;

import java.util.List;

public interface OrgGroupService extends BaseService {

	Page listByPage(Long orgId, String keywords, String disabled, Integer currentPage, Integer itemsPerPage);
	List search(Long orgId, String keywords, JSONArray exceptIds);
	
	TestOrgGroup save(OrgGroupVo vo, Long orgId);
	boolean delete(Long id);

//	void initDefaultBasicDataPers(TestOrg org);

	List<OrgGroupVo> genVos(List<TestOrgGroup> pos);
	OrgGroupVo genVo(TestOrgGroup user);

}
