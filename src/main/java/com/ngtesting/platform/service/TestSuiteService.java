package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseInSuite;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface TestSuiteService extends BaseService {

	List listByPage(Integer projectId, String keywords, Boolean disabled);

    List<TstSuite> query(Integer projectId, String keywords);
    TstSuite get(Integer id);

    TstSuite getWithCases(Integer id);

    TstSuite save(JSONObject json, TstUser optUser);
	TstSuite delete(Integer vo, Integer userId);

	List<TstSuite> list(Integer projectId, String type);

	List<TstSuite> genVos(List<TstSuite> pos);

    TstSuite saveCases(JSONObject json, TstUser optUser);

	TstSuite saveCases(Integer projectId, Integer caseProjectId, Integer runId, List<Integer> ids, TstUser optUser);

	TstSuite genVo(TstSuite po);

    TstSuite genVo(TstSuite po, Boolean withCases);

    Integer countCase(Integer suiteId);

	TstCaseInSuite genCaseVo(TstCaseInSuite po);

    TstSuite updatePo(TstSuite vo);

}
