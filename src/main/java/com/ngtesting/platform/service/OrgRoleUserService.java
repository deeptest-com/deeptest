package com.ngtesting.platform.service;

import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface OrgRoleUserService extends BaseService {

	List<UserVo> listUserByOrgRole(Long orgRoleId);

	boolean saveOrgRoleUsers(Long roleId, List<UserVo> orgUsers);

}
