package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.vo.IsuJqlColumn;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface IsuJqlService extends BaseService {
    List<IsuIssue> query(JsonRule jql, String columns, Integer orgId, Integer projectId);

    JsonRule buildDefaultJql(Integer orgId, Integer projectId);

    List<IsuJqlColumn> buildDefaultColumns(TstUser user);
}
