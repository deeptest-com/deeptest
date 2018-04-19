package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.AiTestEnv;
import com.ngtesting.platform.service.AiTestEnvService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.AiTestEnvVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AiTestEnvServiceImpl extends BaseServiceImpl implements AiTestEnvService {
    @Override
    public List<AiTestEnvVo> listTestEnvVo(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiTestEnv.class);

//        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("displayOrder"));

        List<AiTestEnv> ls = findAllByCriteria(dc);
        List<AiTestEnvVo> vos = genVos(ls);

        return vos;
    }

    @Override
    public List<AiTestEnvVo> genVos(List<AiTestEnv> pos) {
        List<AiTestEnvVo> vos = new LinkedList<>();

        for (AiTestEnv langModel : pos) {
            AiTestEnvVo vo = genVo(langModel);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiTestEnvVo genVo(AiTestEnv po) {
        AiTestEnvVo vo = new AiTestEnvVo();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }
}

