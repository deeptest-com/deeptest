package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface CaseInTaskService extends BaseService {

	List<TstCaseInTask> query(Integer taskId, Integer projectId);
    TstCaseInTask getDetail(Integer id, Integer projectId);
    TstCaseInTask setResult(Integer caseInTaskId, Integer caseId, String result, String status, Integer next, TstUser userVo);

    TstCaseInTask rename(JSONObject json, TstUser userVo);
}
