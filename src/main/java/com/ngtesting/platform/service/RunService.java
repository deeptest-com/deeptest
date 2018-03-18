package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseInRun;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestCaseInRunVo;
import com.ngtesting.platform.vo.TestRunVo;
import com.ngtesting.platform.vo.TestSuiteVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface RunService extends BaseService {

	List<TestCaseInRun> lodaCase(Long runId);
	TestRunVo getById(Long caseId);
	TestRun save(JSONObject json, UserVo optUser);

	boolean importSuiteCasesPers(TestRun run, List<TestSuiteVo> suites);

	TestRun saveCases(Long planId, Long runId, Object[] ids, UserVo optUser);

    TestRun saveCases(JSONObject json, UserVo optUser);

	void addCasesBySuitesPers(Long suiteId, List<Long> suiteIds);
	void addCasesPers(Long suiteId, List<Long> caseIds);

	TestRun delete(Long id, Long userId);
	TestRun closePers(Long id, Long userId);

    void closePlanIfAllRunClosedPers(Long planId);

    List<TestRunVo> genVos(List<TestRun> pos);
	TestRunVo genVo(TestRun po);

	List<TestCaseInRunVo> genCaseVos(List<TestCaseInRun> ls);
	TestCaseInRunVo genCaseVo(TestCaseInRun po);
}
