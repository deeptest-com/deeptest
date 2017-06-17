package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.vo.TestCaseTreeVo;
import com.ngtesting.platform.vo.TestCaseVo;

import java.util.List;

public interface TestCaseService extends BaseService {

	List<TestCase> query(Long projectId, Long moduleId);
	TestCaseVo getById(Long caseId);

	TestCaseTreeVo buildTree(List<TestCase> ls);

	List<TestCaseVo> genVos(List<TestCase> pos);

	TestCaseVo genVo(TestCase po);
	
	TestCase delete(Long vo, Long userId);
	TestCase move(Long id, Long pid, Long prePid, Long userId);
	TestCase create(Long id, String value, String type, Long pid, Long userId);
	TestCase rename(Long id, String value, Long userId);


}
