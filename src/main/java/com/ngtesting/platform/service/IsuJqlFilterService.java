package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IsuJqlFilterService extends BaseService {
    List<IsuJqlFilter> buildUiFilters(String jql, Integer orgId, Integer projectId);

    IsuJqlFilter buildFilter(String name, Integer orgId, Integer projectId);

    IsuJqlFilter buildProjectFilter(Integer orgId);

    IsuJqlFilter buildTypeFilter(Integer orgId, Integer projectId);

    void iterateRuleName(JsonRule rule, List<String> out);
}
