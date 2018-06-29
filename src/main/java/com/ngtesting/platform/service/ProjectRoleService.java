package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProjectRole;

import java.util.List;

public interface ProjectRoleService extends BaseService {

	List list(Integer orgId, String keywords, String disabled);

	TstProjectRole save(TstProjectRole vo, Integer orgId);
	boolean delete(Integer id);

//	TestProjectRoleForOrg createDefaultBasicDataPers(Integer orgId);

    List<TstProjectRole> genVos(List<TstProjectRole> pos);
	TstProjectRole genVo(TstProjectRole role);

}
