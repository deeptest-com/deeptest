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

public interface TestProjectService extends BaseService {

	LinkedList<TestProjectVo> genVos(List<TestProject> pos, Map<String, Integer> ret);
	
	TestProject delete(Long vo, Long userId);
	TestProject save(TestProjectVo vo, Long companyId);

	TestProject getDetail(Long id);

	int countDescendantsNumb(Long id, String childrenPath);

	void toOrderList(TestProjectVo root, String childrenPath, LinkedList<TestProjectVo> resultList);

	Map<String, Object> listCache(Long companyId, String isActive);
	TestProjectVo genVo(TestProject po);

	void removeChildren(LinkedList<TestProjectVo> resultList);

	LinkedList<TestProjectVo> removeMe(LinkedList<TestProjectVo> linkedList, TestProjectVo vo);
	
}
