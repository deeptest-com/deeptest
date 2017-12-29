package com.ngtesting.platform.service;

public interface ProjectRolePriviledgeRelationService extends BaseService {

	void addPriviledgeForLeaderPers(Long projectRole);

	void addPriviledgeForDesignerPers(Long projectRole);

	void addPriviledgeForTesterPers(Long projectRole);

	void addPriviledgePers(Long projectPrivilegeId, Long projectRoleId);
}
