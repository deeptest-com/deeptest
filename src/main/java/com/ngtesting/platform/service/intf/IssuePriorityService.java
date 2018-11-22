package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuPriority;

import java.util.List;

public interface IssuePriorityService extends BaseService {
	List<IsuPriority> list(Integer orgId);

    List<IsuPriority> list(Integer orgId, Integer prjId);
	List<IsuPriority> listNotInSolution(Integer solutionId, Integer orgId);

    IsuPriority get(Integer id, Integer orgId);

	IsuPriority save(IsuPriority vo, Integer orgId);

	Boolean delete(Integer id, Integer orgId);

	Boolean setDefault(Integer id, Integer orgId);

	Boolean changeOrder(Integer id, String act, Integer orgId);
}
