package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IsuJqlFilterService extends BaseService {
    List<IsuJqlFilter> buildUiFilters(JsonRule jql, Integer orgId, Integer projectId);

    IsuJqlFilter buildFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildProjectFilter(String id, String label, Integer orgId, Boolean display);

    IsuJqlFilter buildTypeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildStatusFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildPriorityFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildAssigneeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildCreatorFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildReporterFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildEnvFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildVerFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildResolutionFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildDueTimeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildResolveTimeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    IsuJqlFilter buildCommentsFilter(String id, String label, Integer orgId, Integer projectId, Boolean display);

    void iterateRuleName(JsonRule rule, List<String> out);
}
