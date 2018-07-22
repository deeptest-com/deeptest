package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface OrgGroupUserRelationService extends BaseService {
	List<TstOrgGroupUserRelation> listRelationsByUser(Integer orgId, Integer userId);
	List<TstOrgGroupUserRelation> listRelationsByGroup(Integer orgId, Integer orgGroupId);

	List<TstOrgGroupUserRelation> listRelations(Integer orgId, Integer orgGroupId, Integer userId);

	List<TstOrgGroup> listAllOrgGroups(Integer orgId);

	List<TstUser> listAllOrgUsers(Integer orgId);

	boolean saveRelations(List<TstOrgGroupUserRelation> orgGroupUserVos);

    boolean saveRelations(Integer userId, List<TstOrgGroupUserRelation> orgGroupUserVos);

	TstOrgGroupUserRelation getRelationOrgGroupUser(Integer orgGroupId, Integer userId);

	TstOrgGroupUserRelation genVo(Integer orgId, TstOrgGroup group, Integer userId);

	TstOrgGroupUserRelation genVo(Integer orgId, TstUser user, Integer groupId);
}
