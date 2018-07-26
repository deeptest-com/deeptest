package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseInTaskService extends BaseService {

	List<TstCaseInTask> query(Integer runId);
    TstCaseInTask getById(Integer id);
    TstCaseInTask setResultPers(Integer caseInRunId, String result, String status, Integer next, TstUser userVo);

    TstCaseInTask renamePers(JSONObject json, TstUser userVo);

    List<TstCaseInTask> genVos(List<TstCaseInTask> pos);

    TstCaseInTask addCaseToRunPers(Integer runId, TstCase po, TstUser userVo);
//    TestCaseInRun removeCaseFromRunPers(Long entityId, TstUser userVo);
    TstCaseInTask movePers(JSONObject json, TstUser userVo);

    TstCaseInTask getByRunAndCaseId(Integer runId, Integer caseId);

    void updateLeafAccordingToCasePers(Integer pid);


    TstCaseInTask genVo(TstCaseInTask po, Boolean withSteps);

    List<TstCaseInTaskHistory> findHistories(Integer id);

    void saveHistory(TstUser user, Constant.CaseAct act, TstCaseInTask testCaseInRun,
                     String status, String result);
}
