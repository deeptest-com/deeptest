package cn.linkr.testspace.service;

import java.util.List;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.TestCase;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;

public interface TestCaseService extends BaseService {

	List<TestCase> query(Long projectId, Long moduleId, String keywords);
	TestCaseTreeVo buildTree(List<TestCase> ls);

	List<TestCaseVo> genVos(List<TestCase> pos);

	TestCaseVo genVo(TestCase po);
	
	TestCase delete(Long vo, Long clientId);
	TestCase move(Long id, Long pid, Long prePid, Long id2);
	TestCase create(Long id, String value, Integer type, Long pid, Long id2);
	TestCase rename(Long id, String value, Long id2);
	
}
