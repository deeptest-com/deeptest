package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseStep;
import com.ngtesting.platform.vo.TestCaseStepVo;

public interface CaseStepService extends BaseService {

	TestCaseStep save(JSONObject vo, Long userId);
	TestCaseStep changeOrder(JSONObject vo, String direction, Long userId);
	void createSampleStep(Long caseId);
	boolean delete(Long stepId, Long userId);

	TestCaseStepVo genVo(TestCaseStep step);


}
