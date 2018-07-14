package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgPrivilegeDefine;

import java.util.List;
import java.util.Map;

public interface OrgRolePrivilegeService extends BaseService {

	List<TstOrgPrivilegeDefine> listPrivilegesByOrgRole(Integer orgId, Integer orgRoleId);

	boolean saveOrgRolePrivileges(Integer roleId, List<TstOrgPrivilegeDefine> orgPrivileges);

	Map<String, Boolean> listByUser(Integer userId, Integer orgId);

}
