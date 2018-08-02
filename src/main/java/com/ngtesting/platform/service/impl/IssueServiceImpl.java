package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstCase;
import com.ngtesting.platform.service.IssueService;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueServiceImpl extends BaseServiceImpl implements IssueService {
//    @Autowired
//    CaseCommentsService caseCommentsService;

	@Override
	public List<TstCase> query(Integer projectId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstCase.class);
//
//        if (projectId != null) {
//        	dc.add(Restrictions.eq("projectId", projectId));
//        }
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("pId"));
//        dc.addOrder(Order.asc("ordr"));
//
//        List<TstCase> ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
	}

    @Override
	public TstCase getById(Integer caseId) {
//		TstCase po = (TstCase) getDetail(TstCase.class, caseId);
//		TstCase vo = genVo(po, true);
//
//		return vo;

        return null;
	}

}

