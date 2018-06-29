package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;
import java.util.Map;

public interface ProjectService extends BaseService {

	List<TstProject> listVos(Integer orgId, Integer userId, String keywords, String disabled);
	List<TstProject> listProjectGroups(Integer orgId);

	TstProject getDetail(Integer id);

	TstProject save(TstProject vo, Integer orgId, TstUser userVo);
	Boolean delete(Integer id);

	List<TstProject> listBrothers(Integer projectId);
	List<TstProjectAccessHistory> listRecentProject(Integer orgId, Integer userId);
	List<TstProjectAccessHistory> listRecentProjectVo(Integer orgId, Integer userId);

	TstProject viewPers(Integer projectId, TstUser userVo);
	List<TstProject> list(Integer orgId, String keywords, String disabled);

    void updateNameInHisotyPers(Integer projectId, Integer userId);

	void genHistoryPers(Integer orgId, Integer userId, Integer projectId, String projectName);

	boolean isLastestProjectGroup(Integer orgId, Integer projectGroupId);

	TstProject genVo(TstProject po, Map<String, Map<String, Boolean>> privs);
	List<TstProject> genVos(List<TstProject> pos, String keywords, String disabled, Map<String, Map<String, Boolean>> privs);
	List<TstProject> genVos(List<TstProject> pos, Map<String, Map<String, Boolean>> privs);
	List<TstProject> genGroupVos(List<TstProject> pos);
	TstProjectAccessHistory genHistoryVo(TstProjectAccessHistory po);

    List<Integer> listBrotherIds(Integer projectId);

    List<TstProjectAccessHistory> genHistoryVos(List<TstProjectAccessHistory> pos);

}
