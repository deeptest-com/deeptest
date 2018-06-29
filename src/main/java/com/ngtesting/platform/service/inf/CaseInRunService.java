package com.ngtesting.platform.service.inf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseInRun;
import com.ngtesting.platform.model.TstCaseInRunHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseInRunService extends BaseService {

	List<TstCaseInRun> query(Integer runId);
    TstCaseInRun getById(Integer id);
    TstCaseInRun setResultPers(Integer caseInRunId, String result, String status, Long next, TstUser userVo);

    TstCaseInRun renamePers(JSONObject json, TstUser userVo);

    List<TstCaseInRun> genVos(List<TstCaseInRun> pos);

    TstCaseInRun addCaseToRunPers(Integer runId, TstCase po, TstUser userVo);
//    TestCaseInRun removeCaseFromRunPers(Long entityId, TstUser userVo);
    TstCaseInRun movePers(JSONObject json, TstUser userVo);

    TstCaseInRun getByRunAndCaseId(Integer runId, Integer caseId);

    void updateLeafAccordingToCasePers(Integer pid);


    TstCaseInRun genVo(TstCaseInRun po, Boolean withSteps);

    List<TstCaseInRunHistory> findHistories(Integer id);

    void saveHistory(TstUser user, Constant.CaseAct act, TstCaseInRun testCaseInRun,
                     String status, String result);
}
