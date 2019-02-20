package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.utils.MsgUtil;

public interface CaseInTaskHistoryService extends BaseService {
	void saveHistory(TstUser user, MsgUtil.MsgAction act, Integer caseInTaskIn, String field);

	void saveHistory(Integer caseId, Integer caseInTaskId, MsgUtil.MsgAction act, TstUser user,
					 String status, String result);
}
