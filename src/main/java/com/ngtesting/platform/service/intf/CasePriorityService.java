package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstCasePriority;

import java.util.List;

public interface CasePriorityService extends BaseService {
	List<TstCasePriority> list(Integer orgId);

	TstCasePriority get(Integer id, Integer orgId);

    TstCasePriority getDefault(Integer orgId);

    TstCasePriority save(TstCasePriority vo, Integer orgId);
	Boolean delete(Integer id, Integer orgId);

	Boolean setDefault(Integer id, Integer orgId);

	Boolean changeOrder(Integer id, String act, Integer orgId);
}
