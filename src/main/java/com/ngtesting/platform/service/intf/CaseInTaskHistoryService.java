package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;

public interface CaseInTaskHistoryService extends BaseService {
	void saveHistory(TstUser user, Constant.EntityAct act, Integer caseInTaskIn, String field);

	void saveHistory(Integer caseId, Integer caseInTaskId, Constant.EntityAct act, TstUser user,
					 String status, String result);
}
