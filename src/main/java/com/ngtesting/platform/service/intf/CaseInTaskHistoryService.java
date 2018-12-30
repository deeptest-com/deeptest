package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstUser;

public interface CaseInTaskHistoryService extends BaseService {
	void saveHistory(TstUser user, Constant.EntityAct act, TstCaseInTask testCase, String field);
}
