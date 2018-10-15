package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IsuJqlFilterService extends BaseService {
    List<IsuJqlFilter> buildUiFilters(String jql, Integer orgId, Integer projectId);

    IsuJqlFilter buildFilter(String name, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildProjectFilter(Integer orgId, Boolean display);

    IsuJqlFilter buildTypeFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildStatusFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildPriorityFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildAssigneeFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildCreatorFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildReporterFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildEnvFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildVerFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildResolutionFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildDueTimeFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildResolveTimeFilter(Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildCommentsFilter(Integer orgId, Integer projectId, Boolean display);

    void iterateRuleName(JsonRule rule, List<String> out);
}
