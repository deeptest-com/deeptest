package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProjectRole;

import java.util.List;

public interface ProjectRoleService extends BaseService {

	List list(Integer orgId, String keywords, Boolean disabled);
	TstProjectRole get(Integer roleId);
	TstProjectRole save(TstProjectRole vo, Integer orgId);
	boolean delete(Integer id);

}
