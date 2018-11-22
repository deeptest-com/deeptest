package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface TestPlanService extends BaseService {

	List<TstPlan> listByPage(Integer projectId, String keywords, String status);

	TstPlan getById(Integer caseId, Integer projectId);
	TstPlan save(TstPlan vo, TstUser optUser, Integer projectId);
	Boolean delete(Integer vo, Integer projectId);

	List<TstPlan> listByOrg(Integer orgId);

	List<TstPlan> listByProject(Integer projectId, TstProject.ProjectType type);

	void genVos(List<TstPlan> pos);
	void genVo(TstPlan po);
}
