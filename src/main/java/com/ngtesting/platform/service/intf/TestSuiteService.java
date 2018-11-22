package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface TestSuiteService extends BaseService {

	List listByPage(Integer projectId, String keywords, Boolean disabled);

    List<TstSuite> listForImport(Integer projectId);

    TstSuite get(Integer id, Integer projectId);

    TstSuite getWithCases(Integer id);

    TstSuite save(JSONObject json, TstUser optUser);
    Boolean delete(Integer vo, Integer projectId);

	TstSuite saveCases(Integer projectId, Integer caseProjectId, Integer runId, List<Integer> ids, TstUser optUser);

}
