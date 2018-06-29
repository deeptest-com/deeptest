package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.vo.Page;

import java.util.List;

public interface PlanService extends BaseService {

	Page page(Integer projectId, String status, String keywords, Integer currentPage, Integer itemsPerPage);
	TstPlan getById(Integer caseId);
	TstPlan save(JSONObject json, TstUser optUser);
	TstPlan delete(Integer vo, Integer userId);

	List<TstPlan> listByOrg(Integer orgId);

	List<TstPlan> listByProject(Integer projectId, String type);

	List<TstPlan> genVos(List<TstPlan> pos);
	TstPlan genVo(TstPlan po);

    TstPlan updatePo(TstPlan vo);
}
