package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.vo.TestCaseVo;

import java.util.List;

public interface CaseService extends BaseService {

	List<TestCase> query(Long projectId);
	TestCaseVo getById(Long caseId);

    TestCase rename(JSONObject json, Long userId);
	TestCase delete(Long vo, Long userId);
	TestCaseVo movePers(JSONObject json, Long userId);

	TestCase create(Long id, String value, String type, Long pid, Long userId);

	void loadNodeTree(TestCaseVo vo, TestCase po);

	TestCase save(JSONObject json, Long userId);
	boolean cloneStepsAndChildrenPers(TestCase testcase, TestCase src);

    TestCase saveField(JSONObject json);
	TestCase saveCustomizedField(JSONObject json);

	List<TestCase> getChildren(Long caseId);

	List<TestCaseVo> genVos(List<TestCase> pos);
    List<TestCaseVo> genVos(List<TestCase> pos, boolean withSteps);

    TestCaseVo genVo(TestCase po);
    TestCaseVo genVo(TestCase po, boolean withSteps);

	void copyProperties(TestCase testCasePo, TestCaseVo testCaseVo);
}
