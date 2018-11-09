package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPrioritySolution;

import java.util.List;

public interface IssuePrioritySolutionService extends BaseService {
	List<IsuPrioritySolution> list(Integer orgId);

    List<IsuPrioritySolution> list(Integer orgId, Integer prjId);

    IsuPrioritySolution get(Integer id, Integer orgId);

	IsuPrioritySolution save(IsuPrioritySolution vo, Integer orgId);

	Boolean delete(Integer id, Integer orgId);
}
