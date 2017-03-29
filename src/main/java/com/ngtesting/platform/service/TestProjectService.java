package com.ngtesting.platform.service;

import java.util.HashSet;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import com.ngtesting.platform.entity.EvtGuest;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.vo.GuestVo;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestCaseTreeVo;
import com.ngtesting.platform.vo.TestCaseVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;

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
