package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuPrioritySolution;

import java.util.List;

public interface IssuePrioritySolutionService extends BaseService {
	List<IsuPrioritySolution> list(Integer orgId);

    List<IsuPrioritySolution> list(Integer orgId, Integer prjId);

//    IsuPrioritySolution get(Integer id, Integer orgId);
    IsuPrioritySolution getDetail(Integer solutionId, Integer orgId);

	IsuPrioritySolution save(IsuPrioritySolution vo, Integer orgId);

	Boolean delete(Integer id, Integer orgId);

	Boolean setDefault(Integer id, Integer orgId);

	Boolean addPriority(Integer priorityId, Integer solutionId, Integer orgId);
	Boolean removePriority(Integer priorityId, Integer solutionId, Integer orgId);

	Boolean addAll(Integer solutionId, Integer orgId);
	Boolean removeAll(Integer solutionId, Integer orgId);

	// For Project
	IsuPrioritySolution getByProject(Integer projectId, Integer orgId);

	void setByProject(Integer solutionId, Integer projectId, Integer orgId);
}
