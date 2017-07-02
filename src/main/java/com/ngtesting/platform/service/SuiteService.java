package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestSuite;
import com.ngtesting.platform.vo.TestSuiteTreeVo;
import com.ngtesting.platform.vo.TestSuiteVo;

import java.util.List;

public interface SuiteService extends BaseService {

	List<TestSuite> query(Long projectId);
	TestSuiteVo getById(Long caseId);

	TestSuiteTreeVo buildTree(List<TestSuite> ls);

	List<TestSuiteVo> genVos(List<TestSuite> pos);

	TestSuiteVo genVo(TestSuite po);

	TestSuite delete(Long vo, Long userId);
	TestSuite move(Long id, Long pid, Long prePid, Long userId);
	TestSuite create(Long id, String value, String type, Long pid, Long userId);
	TestSuite rename(Long id, String value, Long userId);
	TestSuite save(JSONObject json);

}
