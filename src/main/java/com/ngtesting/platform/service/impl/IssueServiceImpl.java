package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.service.intf.IssueService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssueServiceImpl extends BaseServiceImpl implements IssueService {
    Log logger = LogFactory.getLog(IssueServiceImpl.class);

    @Autowired
    IssueDao issueDao;

	@Autowired
	IssuePageDao pageDao;

    @Override
	public IsuIssue get(Integer caseId, Integer orgId) {
		IsuIssue po = issueDao.get(caseId, orgId);

		return po;
	}

	@Override
	public IsuPage getPage(Integer orgId, Integer prjId, String opt) {
        IsuType type = getProjectDefaultType(orgId, prjId);
        Map<String, Integer> pageMap = getProjectDefaultPages(orgId, prjId, type.getId());

		IsuPage page = pageDao.get(pageMap.get(opt), orgId);
		return page;
	}

    @Override
    public IsuType getProjectDefaultType(Integer orgId, Integer prjId) {
        IsuType type = issueDao.getProjectDefaultType(orgId, prjId);
        return type;
    }

    @Override
    public Map<String, Integer>  getProjectDefaultPages(Integer orgId, Integer prjId, Integer typeId) {
        List<Map<String, Object>> ls = issueDao.getProjectDefaultPages(orgId, prjId, typeId);

        Map<String, Integer> ret = new HashMap();
        for (Map<String, Object> map : ls) {
            ret.put(map.get("opt").toString(), Integer.valueOf(map.get("pageId").toString()));
        }
        return ret;
    }

}

