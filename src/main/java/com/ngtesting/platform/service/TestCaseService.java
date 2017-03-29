package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.EvtGuest;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.vo.GuestVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestCaseTreeVo;
import com.ngtesting.platform.vo.TestCaseVo;

public interface TestCaseService extends BaseService {

	List<TestCase> query(Long projectId, Long moduleId);
	TestCaseTreeVo buildTree(List<TestCase> ls);

	List<TestCaseVo> genVos(List<TestCase> pos);

	TestCaseVo genVo(TestCase po);
	
	TestCase delete(Long vo, Long clientId);
	TestCase move(Long id, Long pid, Long prePid, Long id2);
	TestCase create(Long id, String value, Integer type, Long pid, Long id2);
	TestCase rename(Long id, String value, Long id2);
	
}
