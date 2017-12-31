package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestProjectPrivilegeDefine;

import java.util.List;

public interface ProjectRolePriviledgeRelationService extends BaseService {

	String addPriviledgeForLeaderPers(List<TestProjectPrivilegeDefine> allProjectPrivileges, Long projectRole);

	String addPriviledgeForDesignerPers(List<TestProjectPrivilegeDefine> allProjectPrivileges, Long projectRole);

	String addPriviledgeForTesterPers(List<TestProjectPrivilegeDefine> allProjectPrivileges, Long projectRole);

	String addPriviledgePers(Long projectPrivilegeId, Long projectRoleId);
}
