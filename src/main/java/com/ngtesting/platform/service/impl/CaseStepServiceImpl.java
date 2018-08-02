package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.CaseStepDao;
import com.ngtesting.platform.model.TstCaseStep;
import com.ngtesting.platform.service.CaseStepService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseStepServiceImpl extends BaseServiceImpl implements CaseStepService {
    @Autowired
    CaseStepDao caseStepDao;

    @Override
    @Transactional
    public TstCaseStep save(JSONObject json, Integer userId) {
        TstCaseStep vo = JSON.parseObject(JSON.toJSONString(json), TstCaseStep.class);

        if (vo.getId() != null) {
            caseStepDao.update(vo);
        } else {
            vo.setId(null);
            caseStepDao.save(vo);
            caseStepDao.moveOthersDown(vo.getCaseId(), vo.getId(), vo.getOrdr());
        }

        return vo;
    }

    @Override
    public boolean delete(Integer stepId, Integer userId) {
        TstCaseStep step = caseStepDao.get(stepId);
        caseStepDao.delete(stepId);

        caseStepDao.moveOthersUp(step.getCaseId(), step.getOrdr());
        return true;
    }

    @Override
    @Transactional
    public TstCaseStep changeOrderPers(JSONObject vo, String direction, Integer userId) {
            TstCaseStep step = caseStepDao.get(vo.getInteger("id"));
            TstCaseStep neighbor = null;
            if ("up".equals(direction)) {
                neighbor = caseStepDao.getPrev(step.getCaseId(), step.getOrdr());
            } else if ("down".equals(direction)) {
                neighbor = caseStepDao.getNext(step.getCaseId(), step.getOrdr());
            }

            Integer stepOrder = step.getOrdr();
            Integer neighborOrder = neighbor.getOrdr();

            caseStepDao.setOrder(step.getId(), neighborOrder);
            caseStepDao.setOrder(neighbor.getId(), stepOrder);

//        TstCaseStep po = (TstCaseStep)getDetail(TstCaseStep.class, vo.getLong("id"));
//        String hql = "from TstCaseStep st where st.deleted = false and st.disabled = false "
//                + " and testCaseId = ?";
//        if ("up".equals(direction)) {
//            hql += " and st.ordr < ? order by ordr desc";
//        } else if ("down".equals(direction)) {
//            hql += " and st.ordr > ? order by ordr asc";
//        }
//        TstCaseStep neighbor = (TstCaseStep) getDao().findFirstByHQL(hql, vo.getLong("caseId"), vo.getInteger("ordr"));
//        TstCaseStep step = (TstCaseStep) getDetail(TstCaseStep.class, vo.getLong("id"));
//
//        Integer order = step.getOrdr();
//        step.setOrdr(neighbor.getOrdr());
//        neighbor.setOrdr(order);
//
//        saveOrUpdate(step);
//        saveOrUpdate(neighbor);
//
//        return po;

        return null;
    }
}
