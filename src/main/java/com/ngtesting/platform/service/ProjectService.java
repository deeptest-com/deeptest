package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstProjectAccessHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;
import java.util.Map;

public interface ProjectService extends BaseService {

	List<TstProject> list(Integer orgId, Integer userId, String keywords, Boolean disabled);
	List<TstProject> listProjectGroups(Integer orgId);

	List<TstProjectAccessHistory> listRecentProject(Integer orgId, Integer userId);

	TstProject get(Integer id);

	TstProject getWithPrivs(Integer id, Integer userId);

	TstProject save(TstProject vo, Integer orgId, TstUser userVo);
	Boolean delete(Integer id, Integer userId);

	TstProject view(Integer projectId, TstUser userVo);

    void updateNameInHisoty(Integer projectId, Integer userId);

	boolean isLastestProjectGroup(Integer orgId, Integer projectGroupId);

	TstProject genVo(TstProject po, Map<String, Map<String, Boolean>> privs);
	List<TstProject> genVos(List<TstProject> pos, Map<String, Map<String, Boolean>> privs);
	List<TstProject> genGroupVos(List<TstProject> pos);
}
