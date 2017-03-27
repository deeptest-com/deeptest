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

	List<TestProjectVo> list(Long companyId, String keywords, String disabled);
	List<TestProjectVo> listGroups(Long companyId);
	
	TestProject getDetail(Long id);

	TestProject save(TestProjectVo vo, UserVo user);
	Boolean delete(Long id, Long userId);
	
	TestProjectVo genVo(TestProject po);
	List<TestProjectVo> genVos(List<TestProject> pos, String keywords, String disabled);
	List<TestProjectVo> genGroupVos(List<TestProject> pos);
	
//	void removeCache(Long companyId);
	
	
	
}
