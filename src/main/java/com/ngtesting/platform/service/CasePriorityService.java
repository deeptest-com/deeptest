package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstCasePriority;

import java.util.List;

public interface CasePriorityService extends BaseService {
	List<TstCasePriority> list(Integer orgId);

	TstCasePriority save(TstCasePriority vo, Integer orgId);
	boolean delete(Integer id);

	List<TstCasePriority> genVos(List<TstCasePriority> pos);
	TstCasePriority genVo(TstCasePriority user);

	boolean setDefaultPers(Integer id, Integer orgId);

	List<TstCasePriority> listVos(Integer orgId);

	boolean changeOrderPers(Integer id, String act, Integer orgId);

//    void createDefaultBasicDataPers(Long id);
}
