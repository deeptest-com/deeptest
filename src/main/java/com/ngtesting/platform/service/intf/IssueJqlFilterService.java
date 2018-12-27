package com.ngtesting.platform.service.intf;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IssueJqlFilterService extends BaseService {
    List<IsuJqlFilter> buildUiFilters(JsonRule jql, Integer orgId, Integer projectId);

    void iterateRuleName(JsonRule rule, List<String> out);
}
