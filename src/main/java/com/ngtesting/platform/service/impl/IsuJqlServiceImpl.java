package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.support.model.JsonRule;
import com.itfsw.query.builder.support.model.enums.EnumOperator;
import com.itfsw.query.builder.support.model.enums.EnumRuleType;
import com.itfsw.query.builder.support.model.result.SqlQueryResult;
import com.ngtesting.platform.dao.IsuTqlDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.service.IsuJqlBuildService;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.IsuJqlService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class IsuJqlServiceImpl extends BaseServiceImpl implements IsuJqlService {
    Log logger = LogFactory.getLog(IsuJqlServiceImpl.class);

    public static String[] defaultFilters = new String[] {"project", };

    @Autowired
    IsuJqlBuildService isuJqlBuildService;

    @Autowired
    IsuJqlFilterService isuJqlFilterService;

    @Autowired
    IsuTqlDao isuTqlDao;

    @Override
    public List<IsuIssue> query(String jql, Integer orgId, Integer projectId) {
        List<IsuIssue> result = new LinkedList<>();

        SqlQueryResult sqlQueryResult = isuJqlBuildService.buildSqlQuery(jql);
        // TODO: 执行查询

        return result;
    }

    @Override
    public String buildDefaultJql(Integer orgId, Integer projectId) {
        JsonRule ret = isuJqlBuildService.genJsonRuleRoot();

        JsonRule projectRule = isuJqlBuildService.genJsonRule(
                "project", "projectId", "select", projectId.toString(),
                EnumOperator.EQUAL, EnumRuleType.INTEGER);
        ret.getRules().add(projectRule);

        return JSON.toJSONString(ret);
    }

}
