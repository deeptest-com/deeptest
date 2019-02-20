package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.utils.MsgUtil;

public interface CaseHistoryService extends BaseService {
	void saveHistory(TstUser user, MsgUtil.MsgAction act, Integer caseId, String field);
}
