package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstEnv;

import java.util.List;

public interface TestEnvService extends BaseService {
	List<TstEnv> list(Integer projectId, String keywords, Boolean disabled);
	TstEnv getById(Integer caseId, Integer projectId);
	TstEnv save(JSONObject json, TstUser optUser);
	Boolean delete(Integer vo, Integer projectId);

	Boolean changeOrder(Integer id, String act, Integer orgId);
}
