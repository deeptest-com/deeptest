package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.model.TstUser;

import java.util.Map;

public interface IssueService extends BaseService {
    IsuIssue get(Integer id);

    IsuIssue get(Integer id, Integer userId, Integer orgId);
    IsuIssue getDetail(Integer id, Integer userId, Integer prjId);

    IsuIssue getData(Integer id, Integer userId, Integer prjId);

	IsuPage getPage(Integer orgId, Integer prjId, String opt);

    Boolean update(JSONObject issue, Integer pageId, TstUser user);

    IsuType getProjectDefaultType(Integer orgId, Integer prjId);

    Map<String, Integer> getProjectDefaultPages(Integer orgId, Integer prjId, Integer typeId);

    IsuIssue save(JSONObject issue, Integer pageId, TstUser user);

    IsuIssue updateField(JSONObject json, TstUser user);

    void delete(Integer id, TstUser user);
}
