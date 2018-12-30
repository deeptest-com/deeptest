package com.ngtesting.platform.service.intf;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.*;

import java.util.List;
import java.util.Map;

public interface IssueService extends BaseService {
	IsuIssue get(Integer id, Integer userId, Integer orgId);
    IsuIssue getDetail(Integer id, Integer userId, Integer prjId);

    IsuIssue getData(Integer id, Integer userId, Integer prjId);

	IsuPage getPage(Integer orgId, Integer prjId, String opt);

    Boolean update(JSONObject issue, Integer pageId, TstUser user);

    IsuType getProjectDefaultType(Integer orgId, Integer prjId);

    Map<String, Integer> getProjectDefaultPages(Integer orgId, Integer prjId, Integer typeId);

    IsuIssue save(JSONObject issue, Integer pageId, TstUser user);

    IsuIssue updateField(JSONObject json, TstUser user);

    void  genDataForExtTable(List<IsuPageElement> elems,
                             List<Object> params,
                             List<IsuPageElement> elems1,
                             List<Object> params1,
                             List<IsuPageElement> elems2,
                             List<Object> params2);

    void delete(Integer id, TstUser user);
}
