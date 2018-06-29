package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;
import java.util.Map;

public interface HistoryService extends BaseService {

	List<TstHistory> listByOrg(Integer orgId);

	List<TstHistory> listByProject(Integer projectId, String projectType);
	TstHistory getById(Integer id);

	TstHistory create(Integer projectI, TstUser optUser, String action,
					  String entityType, Integer entityId, String name);

    Map<String, List<TstHistory>> genVosByDate(List<TstHistory> historyPos);
	List<TstHistory> genVos(List<TstHistory> pos);
	TstHistory genVo(TstHistory po);

}
