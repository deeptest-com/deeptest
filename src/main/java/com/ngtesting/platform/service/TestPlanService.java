package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface TestPlanService extends BaseService {

	List<TstPlan> listByPage(Integer projectId, String keywords, String status);

	TstPlan getById(Integer caseId);
	TstPlan save(JSONObject json, TstUser optUser);
	void delete(Integer vo, Integer userId);

	List<TstPlan> listByOrg(Integer orgId);

	List<TstPlan> listByProject(Integer projectId, TstProject.ProjectType type);

	void genVos(List<TstPlan> pos);
	void genVo(TstPlan po);
}
