package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseStep;

public interface CaseStepService extends BaseService {

	TstCaseStep save(JSONObject vo, Integer userId);

    TstCaseStep changeOrderPers(JSONObject vo, String direction, Integer userId);
	boolean delete(Integer stepId, Integer userId);

}
