package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.AiAsrLangModel;
import com.ngtesting.platform.service.AiAsrLangModelService;
import com.ngtesting.platform.utils.BeanUtilEx;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AiAsrLangModelServiceImpl extends BaseServiceImpl implements AiAsrLangModelService {
    @Override
    public List<AiAsrLangModel> listAsrLangModelVo(Long projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(AiAsrLangModel.class);
//
////        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("displayOrder"));
//
//        List<AiAsrLangModel> ls = findAllByCriteria(dc);
//        List<AiAsrLangModel> vos = genVos(ls);

        return null;
    }

    @Override
    public List<AiAsrLangModel> genVos(List<AiAsrLangModel> pos) {
        List<AiAsrLangModel> vos = new LinkedList<>();

        for (AiAsrLangModel langModel : pos) {
            AiAsrLangModel vo = genVo(langModel);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiAsrLangModel genVo(AiAsrLangModel po) {
        AiAsrLangModel vo = new AiAsrLangModel();
        BeanUtilEx.copyProperties(po, vo);
        return vo;
    }
}

