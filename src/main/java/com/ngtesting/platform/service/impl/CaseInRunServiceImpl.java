package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestCaseInRun;
import com.ngtesting.platform.entity.TestCaseStep;
import com.ngtesting.platform.service.CaseInRunService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseInRunVo;
import com.ngtesting.platform.vo.TestCaseStepVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CaseInRunServiceImpl extends BaseServiceImpl implements CaseInRunService {

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
        if (withSteps) {
            List<TestCaseStep> steps = testcase.getSteps();
            for (TestCaseStep step : steps) {
                TestCaseStepVo stepVo = new TestCaseStepVo(
                        step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());

                vo.getSteps().add(stepVo);
            }
        } else {
            vo.setSteps(null);
        }

        return vo;
    }

}

