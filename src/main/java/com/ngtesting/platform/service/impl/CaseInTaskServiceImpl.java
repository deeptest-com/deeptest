package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstCaseInTaskHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.CaseAttachmentService;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.service.CaseInTaskService;
import com.ngtesting.platform.service.CaseService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CaseInTaskServiceImpl extends BaseServiceImpl implements CaseInTaskService {
    @Autowired
    CaseService caseService;
    @Autowired
    CaseCommentsService caseCommentsService;
    @Autowired
    CaseAttachmentService caseAttachmentService;

    @Override
    public List<TstCaseInTask> query(Integer taskId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInTask.class);
//        dc.createAlias("testCase", "cs");
//
//        dc.add(Restrictions.eq("taskId", taskId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("pId"));
//        dc.addOrder(Order.asc("cs.ordr"));
//
//        List<TestCaseInTask> ls = findAllByCriteria(dc);
//
//        List<TstCaseInTask> vos = genVos(ls);
//
//        return vos;

        return null;
    }

    @Override
    public TstCaseInTask getById(Integer id) {
//        TestCaseInTask po = (TestCaseInTask) get(TestCaseInTask.class, id);
//        TstCaseInTask vo = genVo(po, true);
//
//        return vo;

        return null;
    }

    @Override
    public TstCaseInTask setResultPers(Integer caseInTaskId, String result, String status, Integer nextId, TstUser TstUser) {
//        TestCaseInTask po = (TestCaseInTask) get(TestCaseInTask.class, caseInTaskId);
//        po.setResult(result);
//        po.setStatus(status);
//        po.setExeById(TstUser.getId());
////        if (!"block".equals(status)) {
//            po.setExeTime(new Date());
////        }
//        saveOrUpdate(po);
//
//        saveHistory(TstUser, Constant.CaseAct.exe_result, po, status, result==null?"":result.trim());
//
//        TestTask task = po.getTask();
//        TestPlan plan = task.getPlan();
//        if (task.getStatus().equals(TestTask.TaskStatus.not_start)) {
//            task.setStatus(TestTask.TaskStatus.in_progress);
//            saveOrUpdate(task);
//        }
//        if (plan.getStatus().equals(TestPlan.PlanStatus.not_start)) {
//            plan.setStatus(TestPlan.PlanStatus.in_progress);
//            saveOrUpdate(plan);
//        }
//
//        if (nextId != null) {
//            return getWithCasesById(nextId);
//        } else {
//            return genVo(po, true);
//        }

        return null;
    }

    @Override
    public TstCaseInTask renamePers(JSONObject json, TstUser TstUser) {
//        Long caseId = json.getLong("id");
//        Long entityId = json.getLong("entityId");
//        Long taskId = json.getLong("taskId");
//        String name = json.getString("name");
//        Long pId = json.getLong("pId");
//        Long projectId = json.getLong("projectId");
//
//        TstCaseInTask vo;
//        TestCase casePo = caseService.renamePers(caseId, name, pId, projectId, TstUser);
//
//        if (caseId == null || caseId <= 0) {
//            vo = addCaseToTaskPers(taskId, casePo, TstUser);
//        } else {
//            vo = genVo((TestCaseInTask) get(TestCaseInTask.class, entityId), false);
//        }
//
//        getDao().flush();
//        caseService.updateParentIfNeededPers(vo.getpId());
//        updateLeafAccordingToCasePers(vo.getpId());
//
//        return vo;

        return null;
    }

    @Override
    public TstCaseInTask movePers(JSONObject json, TstUser TstUser) {
//        Long taskId = json.getLong("taskId");
//        Long caseId = json.getLong("srcId");
//
//        Long srcId = json.getLong("srcId");
//        TestCase src = (TestCase) get(TestCase.class, srcId);;
//        Long targetId = json.getLong("targetId");
//
//        Long parentId = src.getpId();
//
//        TstCase vo = caseService.movePers(json, TstUser);
//
//        TestCaseInTask caseInTask = getByTaskAndCaseId(taskId, caseId);
//        caseInTask.setpId(vo.getpId());
//        caseInTask.setLeaf(vo.getLeaf());
//        saveOrUpdate(caseInTask);
//
//        getDao().flush();
//        updateLeafAccordingToCasePers(targetId);
//        updateLeafAccordingToCasePers(parentId);
//
//        return genVo(caseInTask, false);

        return null;
    }

//    @Override
//    public TestCaseInTask removeCaseFromTaskPers(Long entityId, TstUser TstUser) {
//        TestCaseInTask po = (TestCaseInTask) get(TestCaseInTask.class, entityId);
//
//        getDao().querySql("{call remove_case_in_task_and_its_children(?,?,?)}",
//                po.getTaskId(), po.getCaseId(), po.getpId());
//
//        return po;
//    }

    @Override
    public TstCaseInTask getByTaskAndCaseId(Integer taskId, Integer caseId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInTask.class);
//
//        dc.add(Restrictions.eq("taskId", taskId));
//        dc.add(Restrictions.eq("caseId", caseId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//
//        List<TestCaseInTask> ls = findAllByCriteria(dc);
//        if (ls.size() == 1) {
//            return ls.get(0);
//        } else {
//            return null;
//        }

        return null;
    }

    @Override
    public void updateLeafAccordingToCasePers(Integer pid) {

//        getDao().querySql("{call update_case_in_task_leaf(?)}", pid);
    }

    // 执行时新增的用例
    @Override
    public TstCaseInTask addCaseToTaskPers(Integer taskId, TstCase po, TstUser TstUser) {
//        TestTask task = (TestTask)get(TestTask.class, taskId);
//
//        TestCaseInTask caseInTask = new TestCaseInTask(task.getProjectId(), task.getPlanId(),
//                task.getId(), po.getId(), po.getpId(), true);
//        task.getTestCases().add(caseInTask);
//
//        saveOrUpdate(caseInTask);
//        TstCaseInTask vo = genVo(caseInTask, false);
//
//        return vo;

        return null;
    }

    @Override
    public List<TstCaseInTask> genVos(List<TstCaseInTask> pos) {
        List<TstCaseInTask> vos = new LinkedList<>();

//        for (TestCaseInTask po: pos) {
//            TstCaseInTask vo = genVo(po, false);
//            vos.add(vo);
//        }
        return vos;
    }

    @Override
    public TstCaseInTask genVo(TstCaseInTask po, Boolean withSteps) {
        TstCaseInTask vo = new TstCaseInTask();

//        TestCase testcase = (TestCase)get(TestCase.class, po.getCaseId());
//        BeanUtilEx.copyProperties(vo, testcase);
//        BeanUtilEx.copyProperties(vo, po);
//
//        vo.setEntityId(po.getId());
//        vo.setId(testcase.getId());
//
//        vo.setSteps(new LinkedList<TstCaseStep>());
//        vo.setComments(new LinkedList<TstCaseComments>());
//        vo.setAttachments(new LinkedList<TstCaseAttachment>());
//        vo.setHistories(new LinkedList<TstCaseInTaskHistory>());
//
//        if (withSteps) {
//            List<TestCaseStep> steps = testcase.getSteps();
//            for (TestCaseStep step : steps) {
//                TstCaseStep stepVo = new TstCaseStep(
//                        step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());
//
//                vo.getSteps().add(stepVo);
//            }
//
//            List<TestCaseComments> comments = testcase.getComments();
//            Iterator<TestCaseComments> iterator  = comments.iterator();
//            while (iterator.hasNext()) {
//                TestCaseComments comment = iterator.next();
//                TstCaseComments commentVo = caseCommentsService.genVo(comment);
//                vo.getComments().add(commentVo);
//            }
//
//            List<TestCaseInTaskHistory> histories = findHistories(po.getId());
//            for (TestCaseInTaskHistory his : histories) {
//                TstCaseInTaskHistory historyVo = new TstCaseInTaskHistory(
//                        his.getId(), his.getTitle(), his.getDescr(), his.getTestCaseInTaskId(), his.getCreateTime());
//
//                vo.getHistories().add(historyVo);
//            }
//
//            List<TestCaseAttachment> attachments = testcase.getAttachments();
//            Iterator<TestCaseAttachment> iteratorAttach  = attachments.iterator();
//            while (iteratorAttach.hasNext()) {
//                TestCaseAttachment attachment = iteratorAttach.next();
//                TstCaseAttachment attachVo = caseAttachmentService.genVo(attachment);
//                vo.getAttachments().add(attachVo);
//            }
//        } else {
//            vo.setSteps(null);
//            vo.setComments(null);
//        }

        return vo;
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
//                + "为\"" + Constant.ExeStatus.get(status) + "\"";
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
