package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestRelationProjectRoleUser;

public interface RelationProjectRoleUserService extends BaseService {

	List<TestRelationProjectRoleUser> listRelationProjectRoleUsers(
			Long projectRoleId);

	TestRelationProjectRoleUser getRelationProjectRoleUser(Long projectRoleId);

}
