package cn.linkr.testspace.service;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.TestCase;
import cn.linkr.testspace.entity.TestProject;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;
import cn.linkr.testspace.vo.TestProjectVo;
import cn.linkr.testspace.vo.UserVo;

public interface TestProjectService extends BaseService {

	LinkedList<TestProjectVo> genVos(List<TestProject> pos, Map<String, Integer> ret);
	
	Boolean delete(Long id, Long userId);
	TestProject save(TestProjectVo vo, UserVo user);

	TestProject getDetail(Long id);

	void toOrderList(TestProjectVo root, LinkedList<TestProjectVo> resultList);

	Map<String, Object> listCache(Long companyId, String isActive);
	TestProjectVo genVo(TestProject po);

	void removeChildren(LinkedList<TestProjectVo> resultList);

	LinkedList<TestProjectVo> removeChildren(LinkedList<TestProjectVo> linkedList, TestProjectVo vo);

	void removeCache(Long companyId);
	
}
