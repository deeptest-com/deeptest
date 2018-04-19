package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.AiTestSet;
import com.ngtesting.platform.service.AiTestSetService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.AiTestSetVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AiTestSetServiceImpl extends BaseServiceImpl implements AiTestSetService {

    @Override
    public List<AiTestSetVo> listTestSetVo(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiTestSet.class);

//        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("name"));

        List<AiTestSet> ls = findAllByCriteria(dc);
        List<AiTestSetVo> vos = genVos(ls);

        return vos;
    }

    @Override
    public List<AiTestSetVo> genVos(List<AiTestSet> pos) {
        List<AiTestSetVo> vos = new LinkedList<>();

        for (AiTestSet po : pos) {
            AiTestSetVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiTestSetVo genVo(AiTestSet po) {
        AiTestSetVo vo = new AiTestSetVo();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }

}

