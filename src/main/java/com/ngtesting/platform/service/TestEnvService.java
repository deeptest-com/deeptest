package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstEnv;

import java.util.List;

public interface TestEnvService extends BaseService {
	List<TstEnv> list(Integer projectId, String keywords, Boolean disabled);
	TstEnv getById(Integer caseId);
	TstEnv save(JSONObject json, TstUser optUser);
	void delete(Integer vo, Integer userId);

	boolean changeOrder(Integer id, String act, Integer orgId);
}
