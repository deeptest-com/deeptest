package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.model.IsuIssue;

import java.util.List;

public interface IsuJqlService extends BaseService {
    List<IsuIssue> query(JsonRule jql, String columns, Integer orgId, Integer projectId);

    JsonRule buildDefaultJql(Integer orgId, Integer projectId);
}
