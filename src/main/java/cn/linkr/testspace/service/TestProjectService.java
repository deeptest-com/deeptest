package cn.linkr.testspace.service;

import java.util.List;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.TestCase;
import cn.linkr.testspace.entity.TestProject;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.TestCaseTreeVo;
import cn.linkr.testspace.vo.TestCaseVo;
import cn.linkr.testspace.vo.TestProjectVo;

public interface TestProjectService extends BaseService {

	List<TestProject> list(String status, String keywords);

	TestProjectVo genVo(TestProject po);
	List<TestProjectVo> genVos(List<TestProject> pos);
	
	TestProject delete(Long vo, Long clientId);
	TestProject save(Long id, String value, Integer type, Long pid, Long id2);
	
}
