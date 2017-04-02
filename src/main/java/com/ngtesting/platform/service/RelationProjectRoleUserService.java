package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.SysRelationProjectRoleUser;

public interface RelationProjectRoleUserService extends BaseService {

	List<SysRelationProjectRoleUser> listRelationProjectRoleUsers(
			Long projectRoleId);

	SysRelationProjectRoleUser getRelationProjectRoleUser(Long projectRoleId);

}
