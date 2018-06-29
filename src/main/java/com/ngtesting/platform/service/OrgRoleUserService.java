package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface OrgRoleUserService extends BaseService {

	List<TstUser> listUserByOrgRole(Integer orgId, Integer orgRoleId);

	boolean saveOrgRoleUsers(Integer roleId, List<TstUser> orgUsers);

}
