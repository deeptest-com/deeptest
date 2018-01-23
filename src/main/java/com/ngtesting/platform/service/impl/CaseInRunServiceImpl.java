package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.*;
import com.ngtesting.platform.service.CaseCommentsService;
import com.ngtesting.platform.service.CaseInRunService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseCommentsVo;
import com.ngtesting.platform.vo.TestCaseInRunVo;
import com.ngtesting.platform.vo.TestCaseStepVo;
import com.ngtesting.platform.vo.UserVo;
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
    CaseCommentsService caseCommentsService;

    @Override
    public List<TestCaseInRunVo> query(Long runId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRun.class);

        dc.add(Restrictions.eq("runId", runId));

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("pId"));
        dc.addOrder(Order.asc("ordr"));

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
        if (!"block".equals(status)) {
            po.setExeTime(new Date());
        }
        saveOrUpdate(po);

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
    public TestCaseInRunVo addCaseToRunPers(Long runId, TestCase po, UserVo userVo) {
        TestRun run = (TestRun)get(TestRun.class, runId);

        TestCaseInRun caseInRun = new TestCaseInRun(runId, po.getId(), po.getOrdr(), po.getpId());
        run.getTestcases().add(caseInRun);

        saveOrUpdate(caseInRun);
        TestCaseInRunVo vo = genVo(caseInRun, false);

        return vo;
    }

    @Override
    public void removeCasePers(Long runId, Long entityId, UserVo userVo) {
        TestCaseInRun po = (TestCaseInRun) get(TestCaseInRun.class, entityId);
        po.setDeleted(true);
        saveOrUpdate(po);
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
        } else {
            vo.setSteps(null);
            vo.setComments(null);
        }

        return vo;
    }

}

