package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.AiTestSet;
import com.ngtesting.platform.service.AiTestSetService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AiTestSetServiceImpl extends BaseServiceImpl implements AiTestSetService {

    @Override
    public List<AiTestSet> listTestSetVo(Integer projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(AiTestSet.class);
//
////        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("name"));
//
//        List<AiTestSet> ls = findAllByCriteria(dc);
//        List<AiTestSet> vos = genVos(ls);
//
//        return vos;

        return null;
    }

}

