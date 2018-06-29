package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.TstRelationOrgGroupUser;

import java.util.List;

public interface RelationOrgGroupUserService extends BaseService {
	List<TstRelationOrgGroupUser> listRelationsByUser(Integer orgId, Integer userId);
	List<TstRelationOrgGroupUser> listRelationsByGroup(Integer orgId, Integer orgGroupId);

	boolean saveRelations(List<TstRelationOrgGroupUser> orgGroupUserVos);

    boolean saveRelations(Integer userId, List<TstRelationOrgGroupUser> orgGroupUserVos);
}
