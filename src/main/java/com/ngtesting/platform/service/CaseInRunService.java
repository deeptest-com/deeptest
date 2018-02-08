package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestCaseInRun;
import com.ngtesting.platform.vo.TestCaseInRunVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface CaseInRunService extends BaseService {

	List<TestCaseInRunVo> query(Long runId);
    TestCaseInRunVo getById(Long id);
    TestCaseInRunVo setResultPers(Long caseInRunId, String result, String status, Long next, UserVo userVo);

    TestCaseInRunVo renamePers(JSONObject json, UserVo userVo);

    TestCaseInRunVo addCaseToRunPers(Long runId, TestCase po, UserVo userVo);
    TestCaseInRun deleteCaseFromRunPers(Long entityId, UserVo userVo);
    TestCaseInRunVo movePers(JSONObject json, UserVo userVo);

    TestCaseInRun getByRunAndCaseId(Long runId, Long caseId);

    void updateParentIfNeededPers(Long pid);

    List<TestCaseInRunVo> genVos(List<TestCaseInRun> pos);
	TestCaseInRunVo genVo(TestCaseInRun po, Boolean withSteps);

}
