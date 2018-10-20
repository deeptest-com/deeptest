package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.service.IssueService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueServiceImpl extends BaseServiceImpl implements IssueService {
    Log logger = LogFactory.getLog(IssueServiceImpl.class);

    @Autowired
    IssueDao issueDao;

    @Override
	public IsuIssue getById(Integer caseId) {
//		TstCase po = (TstCase) getDetail(TstCase.class, caseId);
//		TstCase vo = genVo(po, true);
//
//		return vo;

        return null;
	}

	@Override
	public void genVos(List<IsuIssue> pos) {
		for (IsuIssue po: pos) {
			genVo(po);
		}
	}

	@Override
	public void genVo(IsuIssue po) {

	}

}

