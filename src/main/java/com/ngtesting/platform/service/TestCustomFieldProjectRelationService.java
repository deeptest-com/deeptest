package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCustomFieldProjectRelation;
import com.ngtesting.platform.model.TstProject;

import java.util.List;

public interface TestCustomFieldProjectRelationService extends BaseService {
	List<TstCustomFieldProjectRelation> listRelationsByField(Integer orgId, Integer fieldId);

	boolean saveRelationsByField(Integer orgId, Integer id, List<TstCustomFieldProjectRelation> relations);

	TstCustomFieldProjectRelation genVo(Integer orgId, TstProject project, Integer fieldId);
}
