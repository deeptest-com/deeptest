package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.vo.TestProjectVo;

public interface TestProjectService extends BaseService {

	List<TestProjectVo> list(Long orgId, String keywords, String disabled);
	List<TestProjectVo> listGroups(Long orgId);
	
	TestProject getDetail(Long id);

	TestProject save(TestProjectVo vo, Long orgId);
	Boolean delete(Long id, Long userId);
	
	TestProjectVo genVo(TestProject po);
	List<TestProjectVo> genVos(List<TestProject> pos, String keywords, String disabled);
	List<TestProjectVo> genGroupVos(List<TestProject> pos);
	
//	void removeCache(Long orgId);
	
	
	
}
