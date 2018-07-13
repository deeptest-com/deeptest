package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseInSuite;
import com.ngtesting.platform.entity.TestSuite;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestCaseInSuiteVo;
import com.ngtesting.platform.vo.TestSuiteVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface SuiteService extends BaseService {

    Page page(Long projectId, String keywords, Integer currentPage, Integer itemsPerPage);
    List<TestSuite> query(Long projectId, String keywords);

	TestSuiteVo getById(Long caseId);

	TestSuiteVo getById(Long caseId, Boolean withCases);

	TestSuite save(JSONObject json, UserVo optUser);
	TestSuite delete(Long vo, Long userId);

	List<TestSuite> list(Long projectId, String type);

	List<TestSuiteVo> genVos(List<TestSuite> pos);

    TestSuite saveCases(JSONObject json, UserVo optUser);

	TestSuite saveCases(Long projectId, Long caseProjectId, Long runId, Object[] ids, UserVo optUser);

	TestSuiteVo genVo(TestSuite po);

    TestSuiteVo genVo(TestSuite po, Boolean withCases);

    void addCasesPers(Long suiteId, List<Long> caseIds);

    Long countCase(Long suiteId);

	TestCaseInSuiteVo genCaseVo(TestCaseInSuite po);

    TestSuite updatePo(TestSuiteVo vo);
}
