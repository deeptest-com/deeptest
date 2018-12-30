package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;

public interface CaseHistoryService extends BaseService {
	void saveHistory(TstUser user, Constant.EntityAct act, Integer caseId, String field);
}
