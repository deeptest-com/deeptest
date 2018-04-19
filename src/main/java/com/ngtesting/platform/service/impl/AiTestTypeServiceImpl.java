package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.AiTestType;
import com.ngtesting.platform.service.AiTestTypeService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.AiTestTypeVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AiTestTypeServiceImpl extends BaseServiceImpl implements AiTestTypeService {

    @Override
    public List<AiTestTypeVo> listTestTypeVo(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiTestType.class);

//        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("displayOrder"));

        List<AiTestType> ls = findAllByCriteria(dc);
        List<AiTestTypeVo> vos = genVos(ls);

        return vos;
    }

    @Override
    public List<AiTestTypeVo> genVos(List<AiTestType> pos) {
        List<AiTestTypeVo> vos = new LinkedList<>();

        for (AiTestType langModel : pos) {
            AiTestTypeVo vo = genVo(langModel);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiTestTypeVo genVo(AiTestType po) {
        AiTestTypeVo vo = new AiTestTypeVo();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }

}

