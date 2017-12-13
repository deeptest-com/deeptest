package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseInRun;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestCaseInRunVo;
import com.ngtesting.platform.vo.TestRunVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface RunService extends BaseService {

	List<TestCaseInRun> lodaCase(Long runId);
	TestRunVo getById(Long caseId);
	TestRun save(JSONObject json, UserVo optUser);
	TestRun saveCases(JSONObject json);
	TestRun delete(Long vo, Long userId);

	List<TestRunVo> genVos(List<TestRun> pos);
	TestRunVo genVo(TestRun po);


	List<TestCaseInRunVo> genCaseVos(List<TestCaseInRun> ls);
	TestCaseInRunVo genCaseVo(TestCaseInRun po);
}
