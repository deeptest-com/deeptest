package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.AiAsrLangModel;
import com.ngtesting.platform.service.AiAsrLangModelService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.AiAsrLangModelVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AiAsrLangModelServiceImpl extends BaseServiceImpl implements AiAsrLangModelService {
    @Override
    public List<AiAsrLangModelVo> listAsrLangModelVo(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiAsrLangModel.class);

//        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("displayOrder"));

        List<AiAsrLangModel> ls = findAllByCriteria(dc);
        List<AiAsrLangModelVo> vos = genVos(ls);

        return vos;
    }

    @Override
    public List<AiAsrLangModelVo> genVos(List<AiAsrLangModel> pos) {
        List<AiAsrLangModelVo> vos = new LinkedList<>();

        for (AiAsrLangModel langModel : pos) {
            AiAsrLangModelVo vo = genVo(langModel);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiAsrLangModelVo genVo(AiAsrLangModel po) {
        AiAsrLangModelVo vo = new AiAsrLangModelVo();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }
}

