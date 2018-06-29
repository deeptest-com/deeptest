package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProjectPrivilegeDefine;

import java.util.List;

public interface ProjectRolePriviledgeRelationService extends BaseService {

	String addPriviledgeForLeaderPers(List<TstProjectPrivilegeDefine> allProjectPrivileges, Integer projectRole);

	String addPriviledgeForDesignerPers(List<TstProjectPrivilegeDefine> allProjectPrivileges, Integer projectRole);

	String addPriviledgeForTesterPers(List<TstProjectPrivilegeDefine> allProjectPrivileges, Integer projectRole);

	String addPriviledgePers(Integer projectPrivilegeId, Integer projectRoleId);
}
