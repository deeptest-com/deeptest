package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseInRun;
import com.ngtesting.platform.model.TstRun;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface RunService extends BaseService {

	List<TstCaseInRun> lodaCase(Integer runId);
	TstRun getById(Integer caseId);
	TstRun save(JSONObject json, TstUser optUser);

	boolean importSuiteCasesPers(TstRun run, List<TstSuite> suites);

	TstRun saveCases(Integer projectId, Integer caseProjectId, Integer planId, Integer runId, Object[] ids, TstUser optUser);

	TstRun saveCases(JSONObject json, TstUser optUser);

	void addCasesBySuitesPers(Integer suiteId, List<Integer> suiteIds);
	void addCasesPers(Integer suiteId, List<Integer> caseIds);

	TstRun delete(Integer id, Integer userId);
	TstRun closePers(Integer id, Integer userId);

    void closePlanIfAllRunClosedPers(Integer planId);

    List<TstRun> genVos(List<TstRun> pos);
	TstRun genVo(TstRun po);

	List<TstCaseInRun> genCaseVos(List<TstCaseInRun> ls);
	TstCaseInRun genCaseVo(TstCaseInRun po);
}
