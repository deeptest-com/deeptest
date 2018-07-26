package com.ngtesting.platform.service;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseHistoryService extends BaseService {

	List<TstCaseHistory> findHistories(Integer testCaseId);

	void saveHistory(TstUser user, Constant.CaseAct act, TstCase testCase, String field);
}
