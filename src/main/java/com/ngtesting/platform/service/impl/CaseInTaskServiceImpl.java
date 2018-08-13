package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.CaseInTaskDao;
import com.ngtesting.platform.dao.TestPlanDao;
import com.ngtesting.platform.dao.TestTaskDao;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseAttachmentService;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.service.CaseInTaskService;
import com.ngtesting.platform.service.CaseService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.Date;
import java.util.List;

@Service
public class CaseInTaskServiceImpl extends BaseServiceImpl implements CaseInTaskService {
    @Autowired
    CaseService caseService;
    @Autowired
    CaseInTaskDao caseInTaskDao;
    @Autowired
    TestTaskDao taskDao;
    @Autowired
    TestPlanDao planDao;

    @Autowired
    CaseCommentsService caseCommentsService;
    @Autowired
    CaseAttachmentService caseAttachmentService;

    @Override
    public List<TstCaseInTask> query(Integer taskId) {
        List<TstCaseInTask> ls = caseInTaskDao.query(taskId);

        return ls;
    }

    @Override
    public TstCaseInTask getDetail(Integer id) {
        TstCaseInTask po = caseInTaskDao.getDetail(id);

        return po;
    }

    @Override
    @Transactional
    public TstCaseInTask setResult(Integer caseInTaskId, String result, String status, Integer nextId, TstUser TstUser) {
        TstCaseInTask po = new TstCaseInTask();

        po.setId(caseInTaskId);
        po.setResult(result);
        po.setStatus(status);
        po.setExeBy(TstUser.getId());
        po.setExeTime(new Date());
        caseInTaskDao.setResult(po);

        saveHistory(TstUser, Constant.CaseAct.exe_result, po, status, result==null?"":result.trim());

        taskDao.start(po.getTaskId());
        planDao.start(po.getPlanId());

        if (nextId != null) {
            return caseInTaskDao.getDetail(nextId);
        } else {
            return caseInTaskDao.getDetail(caseInTaskId);
        }
    }

    @Override
    public List<TstCaseInTaskHistory> findHistories(Integer id) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInTaskHistory.class);
//        dc.add(Restrictions.eq("testCaseInTaskId", id));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.desc("createTime"));
//
//        List<TestCaseInTaskHistory> ls = findAllByCriteria(dc);
//        return ls;

        return null;
    }

    @Override
    public void saveHistory(TstUser user, Constant.CaseAct act, TstCaseInTask testCaseInTask,
                            String status, String result) {
//        String action = act.msg;
//
//        String msg = "用户" + StringUtil.highlightDict(user.getName()) + action
//                + "为\"" + Constant.ExeStatus.getDetail(status) + "\"";
//        if (!StringUtil.IsEmpty(result)) {
//            msg += ", 内容：" + result;
//        }
//
//        TestCaseInTaskHistory his = new TestCaseInTaskHistory();
//        his.setTitle(msg);
//        his.setTestCaseInTaskId(testCaseInTask.getId());
//        saveOrUpdate(his);
    }

}
