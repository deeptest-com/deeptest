package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestHistory;
import com.ngtesting.platform.vo.TestHistoryVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;
import java.util.Map;

public interface HistoryService extends BaseService {

	List<TestHistory> list(Long projectId, String projectType);
	TestHistoryVo getById(Long id);

	TestHistory create(Long projectI, UserVo optUser, String action,
					   TestHistory.TargetType entityType, Long entityId, String name);

    Map<String, List<TestHistoryVo>> genVosByDate(List<TestHistory> historyPos);
	List<TestHistoryVo> genVos(List<TestHistory> pos);
	TestHistoryVo genVo(TestHistory po);

}
