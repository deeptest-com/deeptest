package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestProject;
import com.ngtesting.platform.entity.TestProjectAccessHistory;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
import com.ngtesting.platform.vo.TestProjectVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface ProjectService extends BaseService {

	List<TestProjectVo> listVos(Long orgId, String keywords, String disabled);
	List<TestProjectVo> listProjectGroups(Long orgId);
	
	TestProject getDetail(Long id);

	TestProject save(TestProjectVo vo, Long orgId, Long userId);
	Boolean delete(Long id);
	
//	List<TestProjectAccessHistoryVo> setDefaultPers(Long orgId, Long projectId, UserVo userVo);
	
	List<TestProjectAccessHistory> listRecentProject(Long orgId, Long userId);
	List<TestProjectAccessHistoryVo> listRecentProjectVo(Long orgId, Long userId);
//	Long initDefaultBasicDataPers(Long id, Long userId);

	TestProjectVo viewPers(Long projectId, UserVo userVo);
	List<TestProject> list(Long orgId, String keywords, String disabled);
	boolean isLastestProjectGroup(Long orgId, Long projectGroupId);

	TestProjectVo genVo(TestProject po);
	List<TestProjectVo> genVos(List<TestProject> pos, String keywords, String disabled);
	List<TestProjectVo> genGroupVos(List<TestProject> pos);
	List<TestProjectVo> genVos(List<TestProject> pos);
	TestProjectAccessHistoryVo genHistoryVo(TestProjectAccessHistory po);
	List<TestProjectAccessHistoryVo> genHistoryVos(List<TestProjectAccessHistory> pos);

//	void removeCache(Long orgId);
	
}
