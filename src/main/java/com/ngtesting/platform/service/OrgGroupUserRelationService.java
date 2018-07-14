package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstOrgGroupUserRelation;

import java.util.List;

public interface OrgGroupUserRelationService extends BaseService {
	List<TstOrgGroupUserRelation> listRelationsByUser(Integer orgId, Integer userId);
	List<TstOrgGroupUserRelation> listRelationsByGroup(Integer orgId, Integer orgGroupId);

	boolean saveRelations(List<TstOrgGroupUserRelation> orgGroupUserVos);

    boolean saveRelations(Integer userId, List<TstOrgGroupUserRelation> orgGroupUserVos);
}
