package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface TestTaskService extends BaseService {

	TstTask getById(Integer caseId);
	TstTask save(JSONObject json, TstUser optUser);

	boolean importSuiteCasesPers(TstTask task, List<TstSuite> suites);

	TstTask saveCases(JSONObject json, TstUser optUser);

	void delete(Integer id, Integer userId);
	void closePers(Integer id, Integer userId);
    void closePlanIfAllTaskClosedPers(Integer planId);

	List<TstTask> listByPlan(Integer id);

    List<TstTask> genVos(List<TstTask> pos);
	TstTask genVo(TstTask po);

	List<TstCaseInTask> genCaseVos(List<TstCaseInTask> ls);
	TstCaseInTask genCaseVo(TstCaseInTask po);


}
