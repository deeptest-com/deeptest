package com.ngtesting.platform.service;

import com.itfsw.query.builder.support.model.JsonRule;
import com.itfsw.query.builder.support.model.enums.EnumCondition;
import com.itfsw.query.builder.support.model.enums.EnumOperator;
import com.itfsw.query.builder.support.model.enums.EnumRuleType;
import com.itfsw.query.builder.support.model.result.SqlQueryResult;

public interface IsuJqlBuildService extends BaseService {

    JsonRule genJsonRuleGroup(String id, String field, String input, String val,
                              EnumCondition condition, EnumOperator operator, EnumRuleType type);

    JsonRule genJsonRule(String id, String field, String input, String val,
                         EnumOperator operator, EnumRuleType type);

    JsonRule genJsonRuleRoot();

    SqlQueryResult buildSqlQuery(String jql);
}
