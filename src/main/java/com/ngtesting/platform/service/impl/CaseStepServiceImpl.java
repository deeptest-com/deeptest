package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseStep;
import com.ngtesting.platform.service.inf.CaseStepService;
import com.ngtesting.platform.utils.BeanUtilEx;
import org.springframework.stereotype.Service;

@Service
public class CaseStepServiceImpl extends BaseServiceImpl implements CaseStepService {

    @Override
    public TstCaseStep save(JSONObject json, Integer userId) {
//        TstCaseStep vo = JSON.parseObject(JSON.toJSONString(json), TstCaseStep.class);
//
//		TstCaseStep po = new TstCaseStep();
//
//        if (vo.getId() != null) {
//            po = (TstCaseStep)get(TstCaseStep.class, vo.getId());
//            BeanUtilEx.copyProperties(po, vo);
//        } else {
//            BeanUtilEx.copyProperties(po, vo);
//            po.setId(null);
//            moveOthersPers(vo.getTestCaseId(), vo.getOrdr(), "down");
//        }
//        saveOrUpdate(po);
//
//        return po;

        return null;
    }

    @Override
    public void moveOthersPers(Integer testCaseId, Integer ordr, String direction) {
//        String sql = "update tst_case_step set ";
//        if ("up".equals(direction)) {
//            sql += " ordr=ordr-1";
//        } else if ("down".equals(direction)) {
//            sql += " ordr=ordr+1";
//        }
//
//        sql += " where deleted = false and disabled = false and test_case_id = ? and ";
//        if ("up".equals(direction)) {
//            sql += " ordr > ? ";
//        } else if ("down".equals(direction)) {
//            sql += " ordr >= ? ";
//        }
//        sql += " order by ordr asc";
//        getDao().querySql(sql, testCaseId, ordr);
    }

    @Override
    public TstCaseStep changeOrderPers(JSONObject vo, String direction, Integer userId) {
//        TstCaseStep po = (TstCaseStep)get(TstCaseStep.class, vo.getLong("id"));
//        String hql = "from TstCaseStep st where st.deleted = false and st.disabled = false "
//                + " and testCaseId = ?";
//        if ("up".equals(direction)) {
//            hql += " and st.ordr < ? order by ordr desc";
//        } else if ("down".equals(direction)) {
//            hql += " and st.ordr > ? order by ordr asc";
//        }
//        TstCaseStep neighbor = (TstCaseStep) getDao().findFirstByHQL(hql, vo.getLong("caseId"), vo.getInteger("ordr"));
//        TstCaseStep step = (TstCaseStep) get(TstCaseStep.class, vo.getLong("id"));
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

    @Override
    public void createSampleStep(Integer caseId) {
//        TstCaseStep step = new TstCaseStep(caseId, "步骤", "期待结果", 1);
//        step.setTestCaseId(caseId);
//        saveOrUpdate(step);
    }

    @Override
    public boolean delete(Integer stepId, Integer userId) {
//        TstCaseStep step = (TstCaseStep) get(TstCaseStep.class, stepId);
//        step.setDeleted(true);
//        saveOrUpdate(step);
//
//        moveOthersPers(step.getTestCaseId(), step.getOrdr(), "up");

        return true;
    }

    @Override
    public TstCaseStep genVo(TstCaseStep po) {
        TstCaseStep vo = new TstCaseStep();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }
}
