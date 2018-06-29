package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.inf.AiProductBranchService;
import com.ngtesting.platform.model.AiProductBranch;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AiProductBranchServiceImpl extends BaseServiceImpl implements AiProductBranchService {

    @Override
    public List<AiProductBranch> listForProductBranchVo(Long projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(AiProductBranch.class);
//
////        dc.add(Restrictions.eq("projectId", projectId));
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("displayOrder"));
//
//        List<AiProductBranch> ls = findAllByCriteria(dc);
//        List<AiProductBranch> vos = genVos(ls);

        return null;
    }

}

