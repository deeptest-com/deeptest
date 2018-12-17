package com.ngtesting.platform.service.intf;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.model.IsuFieldDefine;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IsuJqlFilterService extends BaseService {
    List<IsuJqlFilter> buildUiFilters(JsonRule jql, Integer orgId, Integer projectId);

    IsuJqlFilter buildFilter(IsuFieldDefine field, Integer orgId, Integer projectId);

    IsuJqlFilter buildProjectFilter(IsuFieldDefine field, Integer orgId);

    void iterateRuleName(JsonRule rule, List<String> out);
}
