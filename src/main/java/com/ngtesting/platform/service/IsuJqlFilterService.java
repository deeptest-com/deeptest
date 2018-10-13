package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IsuJqlFilterService extends BaseService {
    List<IsuJqlFilter> buildUiFilters(String jql, Integer orgId, Integer projectId);

    IsuJqlFilter buildFilter(String name, Integer orgId, Integer projectId);

    IsuJqlFilter buildProjectFilter(Integer orgId);

    IsuJqlFilter buildTypeFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildStatusFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildPriorityFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildAssigneeFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildCreatorFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildReporterFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildEnvFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildVerFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildResolutionFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildDueTimeFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildResolveTimeFilter(Integer orgId, Integer projectId);

    IsuJqlFilter buildCommentsFilter(Integer orgId, Integer projectId);

    void iterateRuleName(JsonRule rule, List<String> out);
}
