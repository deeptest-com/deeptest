package com.ngtesting.platform.service;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseInTaskService extends BaseService {

	List<TstCaseInTask> query(Integer taskId);
    TstCaseInTask getDetail(Integer id);
    TstCaseInTask setResult(Integer caseInTaskId, String result, String status, Integer next, TstUser userVo);

    List<TstCaseInTaskHistory> findHistories(Integer id);

    void saveHistory(TstUser user, Constant.CaseAct act, TstCaseInTask testCaseInTask,
                     String status, String result);
}
