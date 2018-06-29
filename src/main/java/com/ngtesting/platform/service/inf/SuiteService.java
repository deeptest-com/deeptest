package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseInSuite;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.vo.Page;

import java.util.List;

public interface SuiteService extends BaseService {

    Page page(Integer projectId, String keywords, Integer currentPage, Integer itemsPerPage);
    List<TstSuite> query(Integer projectId, String keywords);

	TstSuite getById(Integer caseId);

	TstSuite getById(Integer caseId, Boolean withCases);

	TstSuite save(JSONObject json, TstUser optUser);
	TstSuite delete(Integer vo, Integer userId);

	List<TstSuite> list(Integer projectId, String type);

	List<TstSuite> genVos(List<TstSuite> pos);

    TstSuite saveCases(JSONObject json, TstUser optUser);

	TstSuite saveCases(Integer projectId, Integer caseProjectId, Integer runId, Object[] ids, TstUser optUser);

	TstSuite genVo(TstSuite po);

    TstSuite genVo(TstSuite po, Boolean withCases);

    void addCasesPers(Integer suiteId, List<Integer> caseIds);

    Integer countCase(Integer suiteId);

	TstCaseInSuite genCaseVo(TstCaseInSuite po);

    TstSuite updatePo(TstSuite vo);
}
