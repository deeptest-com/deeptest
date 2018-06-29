package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVer;

import java.util.List;

public interface VerService extends BaseService {
	List<TstVer> list(Long projectId, String keywords, String disabled);
	TstVer getById(Long caseId);
	TstVer save(JSONObject json, TstUser optUser);
	TstVer delete(Long vo, Long userId);

	boolean changeOrderPers(Long id, String act, Long orgId);
}
