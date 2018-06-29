package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.AiTestEnv;
import com.ngtesting.platform.service.inf.AiTestEnvService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AiTestEnvServiceImpl extends BaseServiceImpl implements AiTestEnvService {
    @Override
    public List<AiTestEnv> listTestEnvVo(Long projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(AiTestEnv.class);
//
////        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("displayOrder"));
//
//        List<AiTestEnv> ls = findAllByCriteria(dc);
//        List<AiTestEnv> vos = genVos(ls);
//
//        return vos;

        return null;
    }

}

