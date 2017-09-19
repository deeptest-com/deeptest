package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseStep;
import com.ngtesting.platform.service.CaseStepService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseStepVo;
import org.springframework.stereotype.Service;

@Service
public class CaseStepServiceImpl extends BaseServiceImpl implements CaseStepService {

    @Override
    public TestCaseStep save(JSONObject vo, Long userId) {
//		TestCaseStep step = (TestCaseStep) get(TestCaseStep.class, vo.getId());

        return null;
    }

    @Override
    public TestCaseStep changeOrder(JSONObject vo, String direction, Long userId) {

        return null;
    }

    @Override
    public void createSampleStep(Long caseId) {
        TestCaseStep step = new TestCaseStep(caseId, "步骤", "期待结果", 1);
        step.setTestCaseId(caseId);
        saveOrUpdate(step);
    }

    @Override
    public boolean delete(Long stepId, Long userId) {
        TestCaseStep step = (TestCaseStep) get(TestCaseStep.class, stepId);
        step.setDeleted(true);
        saveOrUpdate(step);
        return true;
    }

    @Override
    public TestCaseStepVo genVo(TestCaseStep po) {
        TestCaseStepVo vo = new TestCaseStepVo();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }
}
