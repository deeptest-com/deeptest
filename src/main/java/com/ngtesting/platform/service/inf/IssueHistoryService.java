package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;
import java.util.Map;

public interface IssueHistoryService extends BaseService {

	List<TstHistory> list(Integer projectId, String projectType);
	TstHistory getById(Integer id);

	TstHistory create(Integer projectI, TstUser optUser, String action,
					  TstHistory.TargetType entityType, Integer entityId, String name);

    Map<String, List<TstHistory>> genVosByDate(List<TstHistory> historyPos);

}
