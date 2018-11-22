package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.CaseDao;
import com.ngtesting.platform.dao.CaseStepDao;
import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.model.TstCaseStep;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.CaseStepService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseStepServiceImpl extends BaseServiceImpl implements CaseStepService {
    @Autowired
    CaseStepDao caseStepDao;
    @Autowired
    CaseDao caseDao;

    @Override
    @Transactional
    public TstCaseStep save(JSONObject json, TstUser user) {
        TstCaseStep step = JSON.parseObject(JSON.toJSONString(json), TstCaseStep.class);
        TstCase testCase = caseDao.get(step.getCaseId(), user.getDefaultPrjId());

        if (testCase == null) {
            return null;
        }

        if (step.getId() != null) {
            caseStepDao.update(step);
        } else {
            step.setId(null);
            caseStepDao.save(step);
        }

        return step;
    }

    @Override
    @Transactional
    public Boolean delete(Integer stepId, TstUser user) {
        TstCaseStep step = caseStepDao.get(stepId);

        TstCase testCase = caseDao.get(step.getCaseId(), user.getDefaultPrjId());
        if (testCase == null) {
            return false;
        }

        caseStepDao.delete(stepId);

        caseStepDao.moveOthersUp(step.getCaseId(), step.getOrdr());
        return true;
    }

    @Override
    @Transactional
    public Boolean changeOrder(JSONObject vo, String direction, TstUser user) {
        Integer id = vo.getInteger("id");
        TstCaseStep curr = caseStepDao.get(id);

        TstCase testCase = caseDao.get(curr.getCaseId(), user.getDefaultPrjId());

        if (testCase == null) {
            return false;
        }

        TstCaseStep neighbor = null;
        if ("up".equals(direction)) {
            neighbor = caseStepDao.getPrev(curr.getCaseId(), curr.getOrdr());
        } else if ("down".equals(direction)) {
            neighbor = caseStepDao.getNext(curr.getCaseId(), curr.getOrdr());
        }

        if (neighbor != null) {
            Integer stepOrder = curr.getOrdr();
            Integer neighborOrder = neighbor.getOrdr();

            caseStepDao.setOrder(curr.getId(), neighborOrder);
            caseStepDao.setOrder(neighbor.getId(), stepOrder);
        }

        return true;
    }
}
