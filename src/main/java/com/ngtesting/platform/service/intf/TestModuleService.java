package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstModule;

import java.util.List;

public interface TestModuleService extends BaseService {
	List<TstModule> list(Integer projectId, String keywords, Boolean disabled);
	TstModule getById(Integer caseId, Integer projectId);
	TstModule save(JSONObject json, TstUser optUser);
	Boolean delete(Integer vo, Integer userId);

	Boolean changeOrder(Integer id, String act, Integer orgId);
}
