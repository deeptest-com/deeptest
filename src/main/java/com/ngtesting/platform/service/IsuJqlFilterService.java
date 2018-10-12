package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.vo.IsuJqlFilter;

import java.util.List;

public interface IsuJqlFilterService extends BaseService {

    IsuJqlFilter buildFilter(String name);

    IsuJqlFilter buildProjectFilter();

    List<IsuJqlFilter> buildUiFilters(String jql);

    void iterateRuleName(JsonRule rule, List<String> out);
}
