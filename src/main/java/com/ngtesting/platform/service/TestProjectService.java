package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.entity.TestProjectAccessHistory;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;

public interface TestProjectService extends BaseService {

	List<TestProjectVo> list(Long orgId, String keywords, String disabled);
	List<TestProjectVo> listProjectGroups(Long orgId);
	
	TestProject getDetail(Long id);

	TestProject save(TestProjectVo vo, Long orgId);
	Boolean delete(Long id, Long userId);
	
	List<TestProjectAccessHistoryVo> setDefaultPers(Long orgId, Long projectId, UserVo userVo);
	
	List<TestProjectAccessHistory> listRecentProject(Long orgId, Long userId);
	List<TestProjectAccessHistoryVo> listRecentProjectVo(Long orgId, Long userId);
	
	TestProjectVo genVo(TestProject po);
	List<TestProjectVo> genVos(List<TestProject> pos, String keywords, String disabled);
	List<TestProjectVo> genGroupVos(List<TestProject> pos);
	List<TestProjectVo> genVos(List<TestProject> pos);
	TestProjectAccessHistoryVo genHistoryVo(TestProjectAccessHistory po);
	List<TestProjectAccessHistoryVo> genHistoryVos(
			List<TestProjectAccessHistory> pos);
	TestProjectVo viewPers(Long orgId, UserVo userVo, Long projectId);
	
//	void removeCache(Long orgId);
	
}
