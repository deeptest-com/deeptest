package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.vo.TestCaseVo;

import java.util.List;

public interface CaseService extends BaseService {

	List<TestCase> query(Long projectId);
	List<TestCaseVo> queryForSelection(Long projectId, Long runId);

	TestCaseVo getById(Long caseId);

    TestCase renamePers(JSONObject json, Long userId);
	TestCase delete(Long vo, Long userId);

	TestCase renamePers(Long id, String name, Long pId, Long projectId, Long userId);

	TestCaseVo movePers(JSONObject json, Long userId);

	void createRoot(Long projectId, Long userId);

	void loadNodeTree(TestCaseVo vo, TestCase po);

	TestCase save(JSONObject json, Long userId);

	void updateParentIfNeededPers(Long pid);

	boolean cloneStepsAndChildrenPers(TestCase testcase, TestCase src);

    TestCase saveField(JSONObject json);
	TestCase saveCustomizedField(JSONObject json);

	List<TestCase> getChildren(Long caseId);

	List<TestCaseVo> genVos(List<TestCase> pos);
    List<TestCaseVo> genVos(List<TestCase> pos, boolean withSteps);

	List<TestCaseVo> genVos(List<TestCase> pos, List<Long> selectIds, boolean withSteps);

	TestCaseVo genVo(TestCase po);

	TestCaseVo genVo(TestCase po, List<Long> selectIds, boolean withSteps);

	TestCaseVo genVo(TestCase po, boolean withSteps);

	void copyProperties(TestCase testCasePo, TestCaseVo testCaseVo);

    TestCase changeContentTypePers(Long id, String contentType);

    TestCase reviewPassPers(Long id, Boolean pass);
}
