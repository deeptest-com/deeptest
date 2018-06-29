package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.TstCasePriority;

import java.util.List;

public interface IssuePriorityService extends BaseService {
	List<TstCasePriority> list(Integer orgId);

	TstCasePriority save(TstCasePriority vo, Integer orgId);
	boolean delete(Integer id);

	boolean setDefaultPers(Integer id, Integer orgId);

	List<TstCasePriority> listVos(Integer orgId);

	boolean changeOrderPers(Integer id, String act, Integer orgId);

//    void createDefaultBasicDataPers(Integer id);
}
