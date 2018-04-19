package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.AiProductBranch;
import com.ngtesting.platform.service.AiProductBranchService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.AiProductBranchVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AiProductBranchServiceImpl extends BaseServiceImpl implements AiProductBranchService {

    @Override
    public List<AiProductBranchVo> listForProductBranchVo(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(AiProductBranch.class);

//        dc.add(Restrictions.eq("projectId", projectId));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("displayOrder"));

        List<AiProductBranch> ls = findAllByCriteria(dc);
        List<AiProductBranchVo> vos = genVos(ls);

        return vos;
    }

    @Override
    public List<AiProductBranchVo> genVos(List<AiProductBranch> pos) {
        List<AiProductBranchVo> vos = new LinkedList<>();

        for (AiProductBranch langModel : pos) {
            AiProductBranchVo vo = genVo(langModel);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public AiProductBranchVo genVo(AiProductBranch po) {
        AiProductBranchVo vo = new AiProductBranchVo();
        BeanUtilEx.copyProperties(vo, po);
        return vo;
    }
}

