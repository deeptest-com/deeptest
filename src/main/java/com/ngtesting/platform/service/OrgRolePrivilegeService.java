package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgPrivilege;

import java.util.List;
import java.util.Map;

public interface OrgRolePrivilegeService extends BaseService {

	List<TstOrgPrivilege> listPrivilegesByOrgRole(Integer orgId, Integer orgRoleId);

	boolean saveOrgRolePrivileges(Integer roleId, List<TstOrgPrivilege> orgPrivileges);

	Map<String, Boolean> listByUser(Integer userId, Integer orgId);

}
