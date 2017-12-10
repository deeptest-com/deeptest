package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestCaseInRun;
import com.ngtesting.platform.vo.TestCaseInRunVo;

import java.util.List;

public interface CaseInRunService extends BaseService {

	List<TestCaseInRunVo> query(Long runId);
    TestCaseInRunVo getById(Long id);
    TestCaseInRunVo setResultPers(Long caseInRunId, String result, String status, Long next);

    List<TestCaseInRunVo> genVos(List<TestCaseInRun> pos);
	TestCaseInRunVo genVo(TestCaseInRun po, Boolean withSteps);

}
