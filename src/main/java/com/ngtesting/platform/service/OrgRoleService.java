package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgPrivilege;
import com.ngtesting.platform.model.TstOrgRole;

import java.util.List;

public interface OrgRoleService extends BaseService {

	List list(Integer orgId, String keywords, String disabled);

	TstOrgRole save(TstOrgRole vo, Integer orgId);
	boolean delete(Integer id);

//	void initOrgRoleBasicDataPers(Integer orgId);

    List<TstOrgPrivilege> getDefaultPrivByRoleCode(TstOrgRole.OrgRoleCode e);

}
