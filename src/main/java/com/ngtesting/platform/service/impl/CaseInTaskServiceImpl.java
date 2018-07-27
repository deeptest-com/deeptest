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
    public List<TstCaseInTask> query(Integer runId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRun.class);
//        dc.createAlias("testCase", "cs");
//
//        dc.add(Restrictions.eq("runId", runId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("pId"));
//        dc.addOrder(Order.asc("cs.ordr"));
//
//        List<TestCaseInRun> ls = findAllByCriteria(dc);
//
//        List<TstCaseInTask> vos = genVos(ls);
//
//        return vos;

        return null;
    }

    @Override
    public TstCaseInTask getById(Integer id) {
//        TestCaseInRun po = (TestCaseInRun) get(TestCaseInRun.class, id);
//        TstCaseInTask vo = genVo(po, true);
//
//        return vo;

        return null;
    }

    @Override
    public TstCaseInTask setResultPers(Integer caseInRunId, String result, String status, Integer nextId, TstUser TstUser) {
//        TestCaseInRun po = (TestCaseInRun) get(TestCaseInRun.class, caseInRunId);
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
//        TestRun run = po.getRun();
//        TestPlan plan = run.getPlan();
//        if (run.getStatus().equals(TestRun.RunStatus.not_start)) {
//            run.setStatus(TestRun.RunStatus.in_progress);
//            saveOrUpdate(run);
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
//        Long runId = json.getLong("runId");
//        String name = json.getString("name");
//        Long pId = json.getLong("pId");
//        Long projectId = json.getLong("projectId");
//
//        TstCaseInTask vo;
//        TestCase casePo = caseService.renamePers(caseId, name, pId, projectId, TstUser);
//
//        if (caseId == null || caseId <= 0) {
//            vo = addCaseToRunPers(runId, casePo, TstUser);
//        } else {
//            vo = genVo((TestCaseInRun) get(TestCaseInRun.class, entityId), false);
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
//        Long runId = json.getLong("runId");
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
//        TestCaseInRun caseInRun = getByRunAndCaseId(runId, caseId);
//        caseInRun.setpId(vo.getpId());
//        caseInRun.setLeaf(vo.getLeaf());
//        saveOrUpdate(caseInRun);
//
//        getDao().flush();
//        updateLeafAccordingToCasePers(targetId);
//        updateLeafAccordingToCasePers(parentId);
//
//        return genVo(caseInRun, false);

        return null;
    }

//    @Override
//    public TestCaseInRun removeCaseFromRunPers(Long entityId, TstUser TstUser) {
//        TestCaseInRun po = (TestCaseInRun) get(TestCaseInRun.class, entityId);
//
//        getDao().querySql("{call remove_case_in_run_and_its_children(?,?,?)}",
//                po.getRunId(), po.getCaseId(), po.getpId());
//
//        return po;
//    }

    @Override
    public TstCaseInTask getByRunAndCaseId(Integer runId, Integer caseId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRun.class);
//
//        dc.add(Restrictions.eq("runId", runId));
//        dc.add(Restrictions.eq("caseId", caseId));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//
//        List<TestCaseInRun> ls = findAllByCriteria(dc);
//        if (ls.size() == 1) {
//            return ls.get(0);
//        } else {
//            return null;
//        }

        return null;
    }

    @Override
    public void updateLeafAccordingToCasePers(Integer pid) {

//        getDao().querySql("{call update_case_in_run_leaf(?)}", pid);
    }

    // 执行时新增的用例
    @Override
    public TstCaseInTask addCaseToRunPers(Integer runId, TstCase po, TstUser TstUser) {
//        TestRun run = (TestRun)get(TestRun.class, runId);
//
//        TestCaseInRun caseInRun = new TestCaseInRun(run.getProjectId(), run.getPlanId(),
//                run.getId(), po.getId(), po.getpId(), true);
//        run.getTestCases().add(caseInRun);
//
//        saveOrUpdate(caseInRun);
//        TstCaseInTask vo = genVo(caseInRun, false);
//
//        return vo;

        return null;
    }

    @Override
    public List<TstCaseInTask> genVos(List<TstCaseInTask> pos) {
        List<TstCaseInTask> vos = new LinkedList<>();

//        for (TestCaseInRun po: pos) {
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
//            List<TestCaseInRunHistory> histories = findHistories(po.getId());
//            for (TestCaseInRunHistory his : histories) {
//                TstCaseInTaskHistory historyVo = new TstCaseInTaskHistory(
//                        his.getId(), his.getTitle(), his.getDescr(), his.getTestCaseInRunId(), his.getCreateTime());
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
//        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRunHistory.class);
//        dc.add(Restrictions.eq("testCaseInRunId", id));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.desc("createTime"));
//
//        List<TestCaseInRunHistory> ls = findAllByCriteria(dc);
//        return ls;

        return null;
    }

    @Override
    public void saveHistory(TstUser user, Constant.CaseAct act, TstCaseInTask testCaseInRun,
                            String status, String result) {
//        String action = act.msg;
//
//        String msg = "用户" + StringUtil.highlightDict(user.getName()) + action
//                + "为\"" + Constant.ExeStatus.get(status) + "\"";
//        if (!StringUtil.IsEmpty(result)) {
//            msg += ", 内容：" + result;
//        }
//
//        TestCaseInRunHistory his = new TestCaseInRunHistory();
//        his.setTitle(msg);
//        his.setTestCaseInRunId(testCaseInRun.getId());
//        saveOrUpdate(his);
    }

}
