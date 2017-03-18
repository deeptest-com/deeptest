package cn.linkr.testspace.service;

import java.util.HashSet;
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

	List list(String status, String keywords, Long companyId);

	TestProjectVo genVo(TestProject po);
	HashSet<TestProjectVo> genVos(List<TestProject> pos, Map<String, Integer> ret);
	
	TestProject delete(Long vo, Long clientId);
	TestProject save(Long id, String value, Integer type, Long pid, Long id2);

	void toList(TestProjectVo root, HashSet<TestProjectVo> vos);

	int countChildrenNumb(TestProjectVo vo, int count);
	
}
