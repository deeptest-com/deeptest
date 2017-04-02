package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.vo.RelationOrgGroupUserVo;

public interface RelationOrgGroupUserService extends BaseService {
	List<RelationOrgGroupUserVo> listRelationsOrgGroupUsers(Long orgId, Long orgGroupId, Long userId);

	boolean saveRelations(List<RelationOrgGroupUserVo> orgGroupUserVos);
}
