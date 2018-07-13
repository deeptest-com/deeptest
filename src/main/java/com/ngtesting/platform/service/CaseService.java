package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestCaseHistory;
import com.ngtesting.platform.vo.TestCaseVo;
import com.ngtesting.platform.vo.UserVo;
import org.apache.poi.ss.usermodel.Sheet;
import org.apache.poi.xssf.usermodel.XSSFCellStyle;

import java.util.List;

public interface CaseService extends BaseService {

	List<TestCase> query(Long projectId);

	List<TestCaseVo> queryForSuiteSelection(Long projectId, Long caseProjectId, Long suiteId);

	List<TestCaseVo> queryForRunSelection(Long projectId, Long caseProjectId, Long runId);

	TestCaseVo getById(Long caseId);

    TestCase renamePers(JSONObject json, UserVo user);
	TestCase delete(Long vo, UserVo user);

	TestCase renamePers(Long id, String name, Long pId, Long projectId, UserVo user);

	TestCaseVo movePers(JSONObject json, UserVo user);

	void createRoot(Long projectId, UserVo user);

	void loadNodeTree(TestCaseVo vo, TestCase po);

	TestCase save(JSONObject json, UserVo user);

	String export(Long projectId);

	void writeHeader(Sheet sheet, Integer rowCount, XSSFCellStyle cellStyle);

	void writeTestCase(TestCase testCase, Sheet sheet, Integer rowCount, XSSFCellStyle cellStyle);

    void updateParentIfNeededPers(Long pid);

	boolean cloneStepsAndChildrenPers(TestCase testcase, TestCase src);

	void saveHistory(UserVo user, Constant.CaseAct act, TestCase testCase, String field);

	TestCase saveField(JSONObject json, UserVo user);

	List<TestCase> getChildren(Long caseId);

	List<TestCaseVo> genVos(List<TestCase> pos);
    List<TestCaseVo> genVos(List<TestCase> pos, boolean withSteps);

	List<TestCaseVo> genVos(List<TestCase> pos, List<Long> selectIds, boolean withSteps);

	TestCaseVo genVo(TestCase po);

	TestCaseVo genVo(TestCase po, List<Long> selectIds, boolean withSteps);

	TestCaseVo genVo(TestCase po, boolean withSteps);

    List<TestCaseHistory> findHistories(Long testCaseId);

    void copyProperties(TestCase testCasePo, TestCaseVo testCaseVo);

    TestCase changeContentTypePers(Long id, String contentType);

    TestCase reviewPassPers(Long id, Boolean pass);
}
