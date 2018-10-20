package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.model.IsuFieldDefine;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IsuJqlFilterService extends BaseService {
    List<IsuJqlFilter> buildUiFilters(JsonRule jql, Integer orgId, Integer projectId);

    IsuJqlFilter buildFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildProjectFilter(IsuFieldDefine field, Integer orgId);

    IsuJqlFilter buildTypeFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildStatusFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildPriorityFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildAssigneeFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildCreatorFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildReporterFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildVerFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildEnvFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildResolutionFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildDueTimeFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildResolveTimeFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildCommentsFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    void iterateRuleName(JsonRule rule, List<String> out);
}
