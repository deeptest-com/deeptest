package com.ngtesting.platform.service.intf;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.model.IsuIssue;

import java.util.List;
import java.util.Map;

public interface IssueJqlService extends BaseService {
    List<IsuIssue> query(JsonRule jql, String columns, List<Map<String, String>> orderBy, Integer orgId, Integer projectId);

    JsonRule buildEmptyJql();

    List<Map<String, String>> buildDefaultOrderBy();
}
