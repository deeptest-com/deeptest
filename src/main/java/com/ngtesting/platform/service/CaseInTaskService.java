package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseInTaskService extends BaseService {

	List<TstCaseInTask> query(Integer taskId);
    TstCaseInTask getDetail(Integer id);
    TstCaseInTask setResultPers(Integer caseInTaskId, String result, String status, Integer next, TstUser userVo);

    TstCaseInTask renamePers(JSONObject json, TstUser userVo);

    List<TstCaseInTask> genVos(List<TstCaseInTask> pos);

    TstCaseInTask addCaseToTaskPers(Integer taskId, TstCase po, TstUser userVo);
//    TestCaseInTask removeCaseFromTaskPers(Long entityId, TstUser userVo);
    TstCaseInTask movePers(JSONObject json, TstUser userVo);

    TstCaseInTask getByTaskAndCaseId(Integer taskId, Integer caseId);

    void updateLeafAccordingToCasePers(Integer pid);


    TstCaseInTask genVo(TstCaseInTask po, Boolean withSteps);

    List<TstCaseInTaskHistory> findHistories(Integer id);

    void saveHistory(TstUser user, Constant.CaseAct act, TstCaseInTask testCaseInTask,
                     String status, String result);
}
