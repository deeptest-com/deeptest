package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.*;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseHistoryService;
import com.ngtesting.platform.service.CaseInTaskService;
import com.ngtesting.platform.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class CaseInTaskServiceImpl extends BaseServiceImpl implements CaseInTaskService {
    @Autowired
    CaseHistoryService caseHistoryService;
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

        caseHistoryService.saveHistory(user, Constant.CaseAct.rename, testCase,null);

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

        saveHistory(caseId, caseInTaskId,
                Constant.CaseAct.exe_result, user, status, result==null?"":result.trim());

        taskDao.start(testCase.getTaskId());
        planDao.start(testCase.getPlanId());

        if (nextId != null) {
            return caseInTaskDao.getDetail(nextId, projectId);
        } else {
            return caseInTaskDao.getDetail(caseInTaskId, projectId);
        }
    }

    @Override
    public void saveHistory(Integer caseId, Integer caseInTaskId, Constant.CaseAct act, TstUser user,
                            String status, String result) {
        String action = act.msg;

        String msg = "用户" + StringUtil.highlightDict(user.getNickname()) + action
                + "为\"" + Constant.ExeStatus.get(status) + "\"";
        if (!StringUtil.IsEmpty(result)) {
            msg += ", 结果内容：" + result;
        }

        TstCaseInTaskHistory his = new TstCaseInTaskHistory();
        his.setTitle(msg);
        his.setCaseId(caseId);
        his.setCaseInTaskId(caseInTaskId);

        caseInTaskHistoryDao.save(his);
    }

}
