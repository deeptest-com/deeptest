package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseStep;
import com.ngtesting.platform.service.CaseStepService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseStepVo;
import org.springframework.stereotype.Service;

@Service
public class CaseStepServiceImpl extends BaseServiceImpl implements CaseStepService {

    @Override
    public TestCaseStep save(JSONObject json, Long userId) {
        TestCaseStepVo vo = JSON.parseObject(JSON.toJSONString(json), TestCaseStepVo.class);

		TestCaseStep po = new TestCaseStep();

        if (vo.getId() != null) {
            po = (TestCaseStep)get(TestCaseStep.class, vo.getId());
            BeanUtilEx.copyProperties(po, vo);
        } else {
            BeanUtilEx.copyProperties(po, vo);
            po.setId(null);
        }
        saveOrUpdate(po);

        return po;
    }

    @Override
    public TestCaseStep changeOrderPers(JSONObject vo, String direction, Long userId) {
        TestCaseStep po = (TestCaseStep)get(TestCaseStep.class, vo.getLong("id"));
        String hql = "from TestCaseStep st where st.deleted = false and st.disabled = false "
                + " and testCaseId = ?";
        if ("up".equals(direction)) {
            hql += " and st.ordr < ? order by ordr desc";
        } else if ("down".equals(direction)) {
            hql += " and st.ordr > ? order by ordr asc";
        }
        TestCaseStep neighbor = (TestCaseStep) getDao().findFirstByHQL(hql, vo.getLong("caseId"), vo.getInteger("ordr"));
        TestCaseStep step = (TestCaseStep) get(TestCaseStep.class, vo.getLong("id"));

        Integer order = step.getOrdr();
        step.setOrdr(neighbor.getOrdr());
        neighbor.setOrdr(order);

        saveOrUpdate(step);
        saveOrUpdate(neighbor);

        return po;
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
