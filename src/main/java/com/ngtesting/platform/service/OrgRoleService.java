package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgPrivilegeDefine;
import com.ngtesting.platform.model.TstOrgRole;

import java.util.List;

public interface OrgRoleService extends BaseService {

	List list(Integer orgId, String keywords, Boolean disabled);
    List<TstOrgRole> listAllOrgRoles(Integer orgId);

	TstOrgRole get(Integer orgRoleId);

	TstOrgRole save(TstOrgRole vo, Integer orgId);
	boolean delete(Integer id);

    List<TstOrgPrivilegeDefine> getDefaultPrivByRoleCode(TstOrgRole.OrgRoleCode e);
}
