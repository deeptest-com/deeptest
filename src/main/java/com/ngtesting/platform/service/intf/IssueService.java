package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuType;

import java.util.Map;

public interface IssueService extends BaseService {
	IsuIssue get(Integer id, Integer orgId);
	IsuPage getPage(Integer orgId, Integer prjId, String opt);

    IsuType getProjectDefaultType(Integer orgId, Integer prjId);

    Map<String, Integer> getProjectDefaultPages(Integer orgId, Integer prjId, Integer typeId);
}
