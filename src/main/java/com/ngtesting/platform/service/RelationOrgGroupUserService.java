package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.vo.RelationOrgGroupUserVo;

public interface RelationOrgGroupUserService extends BaseService {
	List<RelationOrgGroupUserVo> listRelationsByUser(Long orgId, Long userId);
	List<RelationOrgGroupUserVo> listRelationsByGroup(Long orgId, Long orgGroupId);

	boolean saveRelations(List<RelationOrgGroupUserVo> orgGroupUserVos);

}
