package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseStep;
import com.ngtesting.platform.model.TstUser;

public interface CaseStepService extends BaseService {

	TstCaseStep save(JSONObject vo, TstUser user);

	Boolean changeOrder(JSONObject vo, String direction, TstUser user);
	Boolean delete(Integer stepId, TstUser user);

}
