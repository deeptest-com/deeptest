package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.*;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.service.intf.CaseInTaskHistoryService;
import com.ngtesting.platform.service.intf.CaseInTaskService;
import com.ngtesting.platform.utils.StringUtil;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class CaseInTaskServiceImpl extends BaseServiceImpl implements CaseInTaskService {
    @Autowired
    CaseHistoryService caseHistoryService;
    @Autowired
    CaseInTaskHistoryService caseInTaskHistoryService;

    @Autowired
    CaseDao caseDao;
    @Autowired
    CaseInTaskDao caseInTaskDao;
    @Autowired
    TestTaskDao taskDao;
    @Autowired
    TestPlanDao planDao;

    @Autowired
    CaseInTaskHistoryDao caseInTaskHistoryDao;

    @Override
    public List<TstCaseInTask> query(Integer taskId, Integer projectId) {
        List<TstCaseInTask> ls = caseInTaskDao.query(taskId, projectId);

        return ls;
    }

    @Override
    public TstCaseInTask getDetail(Integer id, Integer projectId) {
        TstCaseInTask po = caseInTaskDao.getDetail(id, projectId);

        return po;
    }

    @Override
    public TstCaseInTask rename(JSONObject json, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        Integer caseId = json.getInteger("id");
        Integer entityId = json.getInteger("entityId");
        String name = json.getString("name");

        TstCase testCase = caseDao.get(caseId, projectId);
        if (testCase == null) {
            return null;
        }

        testCase.setUpdateById(user.getId());

        testCase.setName(name);
        testCase.setReviewResult(null);

        caseDao.renameUpdate(testCase);

        caseHistoryService.saveHistory(user, Constant.EntityAct.rename, caseId,null);
        caseInTaskHistoryService.saveHistory(user, Constant.EntityAct.rename, entityId,null);

        return caseInTaskDao.getDetail(entityId, projectId);
    }

    @Override
    @Transactional
    public TstCaseInTask setResult(Integer caseInTaskId, Integer caseId, String result, String status, Integer nextId, TstUser user) {
        Integer projectId = user.getDefaultPrjId();

        TstCaseInTask testCase = caseInTaskDao.getDetail(caseInTaskId, projectId);
        if (testCase == null) {
            return null;
        }

        caseInTaskDao.setResult(caseInTaskId, result, status, user.getId());

        caseInTaskHistoryService.saveHistory(caseId, caseInTaskId,
                Constant.EntityAct.exe_result, user, status, result==null?"":result.trim());

        taskDao.start(testCase.getTaskId());
        planDao.start(testCase.getPlanId());

        if (nextId != null) {
            return caseInTaskDao.getDetail(nextId, projectId);
        } else {
            return caseInTaskDao.getDetail(caseInTaskId, projectId);
        }
    }

}
