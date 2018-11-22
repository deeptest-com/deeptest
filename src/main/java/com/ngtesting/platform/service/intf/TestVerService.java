package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;

import java.util.List;

public interface TestVerService extends BaseService {
	List<TstVer> list(Integer projectId, String keywords, Boolean disabled);
	TstVer getById(Integer caseId, Integer projectId);
	TstVer save(JSONObject json, TstUser optUser);
	Boolean delete(Integer vo, Integer userId);

	Boolean changeOrder(Integer id, String act, Integer orgId);
}
