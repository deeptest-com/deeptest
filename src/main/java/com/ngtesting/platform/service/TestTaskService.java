package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface TestTaskService extends BaseService {

	TstTask getById(Integer caseId, Integer projectId);
	TstTask save(JSONObject json, TstUser optUser);

	void importSuiteCasesPers(TstTask task, List<TstSuite> suites, TstUser optUser);

	TstTask saveCases(JSONObject json, TstUser optUser);

	void delete(Integer id, Integer userId);
	void closePers(Integer id, Integer userId);
    void closePlanIfAllTaskClosedPers(Integer planId);

	List<TstTask> listByPlan(Integer id);

    List<TstTask> genVos(List<TstTask> pos);
	TstTask genVo(TstTask po);

}
