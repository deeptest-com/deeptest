package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstOrgGroup;
import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface OrgGroupUserRelationService extends BaseService {
    List<TstOrgGroupUserRelation> listRelationsByGroup(Integer orgId, Integer orgGroupId);
	List<TstOrgGroupUserRelation> listRelationsByUser(Integer orgId, Integer userId);

	List<TstOrgGroup> listAllOrgGroups(Integer orgId);

	List<TstUser> listAllOrgUsers(Integer orgId);

	boolean saveRelationsForUser(Integer orgId, Integer userId, List<TstOrgGroupUserRelation> orgGroupUserVos);
    boolean saveRelationsForGroup(Integer orgId, Integer groupId, List<TstOrgGroupUserRelation> orgGroupUserVos);

	TstOrgGroupUserRelation genVo(Integer orgId, TstOrgGroup group, Integer userId);

	TstOrgGroupUserRelation genVo(Integer orgId, TstUser user, Integer groupId);
}
