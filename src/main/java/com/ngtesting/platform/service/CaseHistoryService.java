package com.ngtesting.platform.service;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstUser;

public interface CaseHistoryService extends BaseService {
	void saveHistory(TstUser user, Constant.CaseAct act, TstCase testCase, String field);
}
