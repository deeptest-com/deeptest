package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.*;
import com.ngtesting.platform.service.CaseAttachmentService;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.service.CaseInRunService;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.StringUtil;
import com.ngtesting.platform.vo.*;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.Iterator;
import java.util.LinkedList;
import java.util.List;

@Service
public class CaseInRunServiceImpl extends BaseServiceImpl implements CaseInRunService {
    @Autowired
    CaseService caseService;
    @Autowired
    CaseCommentsService caseCommentsService;
    @Autowired
    CaseAttachmentService caseAttachmentService;

    @Override
    public List<TestCaseInRunVo> query(Long runId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRun.class);
        dc.createAlias("testCase", "cs");

        dc.add(Restrictions.eq("runId", runId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("pId"));
        dc.addOrder(Order.asc("cs.ordr"));

        List<TestCaseInRun> ls = findAllByCriteria(dc);

        List<TestCaseInRunVo> vos = genVos(ls);

        return vos;
    }

    @Override
    public TestCaseInRunVo getById(Long id) {
        TestCaseInRun po = (TestCaseInRun) get(TestCaseInRun.class, id);
        TestCaseInRunVo vo = genVo(po, true);

        return vo;
    }

    @Override
    public TestCaseInRunVo setResultPers(Long caseInRunId, String result, String status, Long nextId, UserVo userVo) {
        TestCaseInRun po = (TestCaseInRun) get(TestCaseInRun.class, caseInRunId);
        po.setResult(result);
        po.setStatus(status);
        po.setExeById(userVo.getId());
//        if (!"block".equals(status)) {
            po.setExeTime(new Date());
//        }
        saveOrUpdate(po);

        saveHistory(userVo, Constant.CaseAct.exe_result, po, status, result==null?"":result.trim());

        TestRun run = po.getRun();
        TestPlan plan = run.getPlan();
        if (run.getStatus().equals(TestRun.RunStatus.not_start)) {
            run.setStatus(TestRun.RunStatus.in_progress);
            saveOrUpdate(run);
        }
        if (plan.getStatus().equals(TestPlan.PlanStatus.not_start)) {
            plan.setStatus(TestPlan.PlanStatus.in_progress);
            saveOrUpdate(plan);
        }

        if (nextId != null) {
            return getById(nextId);
        } else {
            return genVo(po, true);
        }
    }

    @Override
    public TestCaseInRunVo renamePers(JSONObject json, UserVo userVo) {
        Long caseId = json.getLong("id");
        Long entityId = json.getLong("entityId");
        Long runId = json.getLong("runId");
        String name = json.getString("name");
        Long pId = json.getLong("pId");
        Long projectId = json.getLong("projectId");

        TestCaseInRunVo vo;
        TestCase casePo = caseService.renamePers(caseId, name, pId, projectId, userVo);

        if (caseId == null || caseId <= 0) {
            vo = addCaseToRunPers(runId, casePo, userVo);
        } else {
            vo = genVo((TestCaseInRun) get(TestCaseInRun.class, entityId), false);
        }

        getDao().flush();
        caseService.updateParentIfNeededPers(vo.getpId());
        updateLeafAccordingToCasePers(vo.getpId());

        return vo;
    }

    @Override
    public TestCaseInRunVo movePers(JSONObject json, UserVo userVo) {
        Long runId = json.getLong("runId");
        Long caseId = json.getLong("srcId");

        Long srcId = json.getLong("srcId");
        TestCase src = (TestCase) get(TestCase.class, srcId);;
        Long targetId = json.getLong("targetId");

        Long parentId = src.getpId();

        TestCaseVo vo = caseService.movePers(json, userVo);

        TestCaseInRun caseInRun = getByRunAndCaseId(runId, caseId);
        caseInRun.setpId(vo.getpId());
        caseInRun.setLeaf(vo.getLeaf());
        saveOrUpdate(caseInRun);

        getDao().flush();
        updateLeafAccordingToCasePers(targetId);
        updateLeafAccordingToCasePers(parentId);

        return genVo(caseInRun, false);
    }

//    @Override
//    public TestCaseInRun removeCaseFromRunPers(Long entityId, UserVo userVo) {
//        TestCaseInRun po = (TestCaseInRun) get(TestCaseInRun.class, entityId);
//
//        getDao().querySql("{call remove_case_in_run_and_its_children(?,?,?)}",
//                po.getRunId(), po.getCaseId(), po.getpId());
//
//        return po;
//    }

    @Override
    public TestCaseInRun getByRunAndCaseId(Long runId, Long caseId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRun.class);

        dc.add(Restrictions.eq("runId", runId));
        dc.add(Restrictions.eq("caseId", caseId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("id"));

        List<TestCaseInRun> ls = findAllByCriteria(dc);
        if (ls.size() == 1) {
            return ls.get(0);
        } else {
            return null;
        }
    }

    @Override
    public void updateLeafAccordingToCasePers(Long pid) {
        getDao().querySql("{call update_case_in_run_leaf(?)}", pid);
    }

    @Override
    public List<TestCaseInRunVo> genVos(List<TestCaseInRun> pos) {
        List<TestCaseInRunVo> vos = new LinkedList<>();

        for (TestCaseInRun po: pos) {
            TestCaseInRunVo vo = genVo(po, false);
            vos.add(vo);
        }
        return vos;
    }

    // 执行时新增的用例
    @Override
    public TestCaseInRunVo addCaseToRunPers(Long runId, TestCase po, UserVo userVo) {
        TestRun run = (TestRun)get(TestRun.class, runId);

        TestCaseInRun caseInRun = new TestCaseInRun(run.getProjectId(), run.getPlanId(),
                run.getId(), po.getId(), po.getpId(), true);
        run.getTestcases().add(caseInRun);

        saveOrUpdate(caseInRun);
        TestCaseInRunVo vo = genVo(caseInRun, false);

        return vo;
    }

    @Override
    public TestCaseInRunVo genVo(TestCaseInRun po, Boolean withSteps) {
        TestCaseInRunVo vo = new TestCaseInRunVo();

        TestCase testcase = (TestCase)get(TestCase.class, po.getCaseId());
        BeanUtilEx.copyProperties(vo, testcase);
        BeanUtilEx.copyProperties(vo, po);

        vo.setEntityId(po.getId());
        vo.setId(testcase.getId());

        vo.setSteps(new LinkedList<TestCaseStepVo>());
        vo.setComments(new LinkedList<TestCaseCommentsVo>());
        vo.setAttachments(new LinkedList<TestCaseAttachmentVo>());
        vo.setHistories(new LinkedList<TestCaseInRunHistoryVo>());

        if (withSteps) {
            List<TestCaseStep> steps = testcase.getSteps();
            for (TestCaseStep step : steps) {
                TestCaseStepVo stepVo = new TestCaseStepVo(
                        step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());

                vo.getSteps().add(stepVo);
            }

            List<TestCaseComments> comments = testcase.getComments();
            Iterator<TestCaseComments> iterator  = comments.iterator();
            while (iterator.hasNext()) {
                TestCaseComments comment = iterator.next();
                TestCaseCommentsVo commentVo = caseCommentsService.genVo(comment);
                vo.getComments().add(commentVo);
            }

            List<TestCaseInRunHistory> histories = findHistories(po.getId());
            for (TestCaseInRunHistory his : histories) {
                TestCaseInRunHistoryVo historyVo = new TestCaseInRunHistoryVo(
                        his.getId(), his.getTitle(), his.getDescr(), his.getTestCaseInRunId(), his.getCreateTime());

                vo.getHistories().add(historyVo);
            }

            List<TestCaseAttachment> attachments = testcase.getAttachments();
            Iterator<TestCaseAttachment> iteratorAttach  = attachments.iterator();
            while (iteratorAttach.hasNext()) {
                TestCaseAttachment attachment = iteratorAttach.next();
                TestCaseAttachmentVo attachVo = caseAttachmentService.genVo(attachment);
                vo.getAttachments().add(attachVo);
            }
        } else {
            vo.setSteps(null);
            vo.setComments(null);
        }

        return vo;
    }

    @Override
    public List<TestCaseInRunHistory> findHistories(Long id) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRunHistory.class);
        dc.add(Restrictions.eq("testCaseInRunId", id));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.desc("createTime"));

        List<TestCaseInRunHistory> ls = findAllByCriteria(dc);
        return ls;
    }

    @Override
    public void saveHistory(UserVo user, Constant.CaseAct act, TestCaseInRun testCaseInRun,
                            String status, String result) {
        String action = act.msg;

        String msg = "用户" + StringUtil.highlightDict(user.getName()) + action
                + "为\"" + Constant.ExeStatus.get(status) + "\"";
        if (!StringUtil.IsEmpty(result)) {
            msg += ", 内容：" + result;
        }

        TestCaseInRunHistory his = new TestCaseInRunHistory();
        his.setTitle(msg);
        his.setTestCaseInRunId(testCaseInRun.getId());
        saveOrUpdate(his);
    }

}
