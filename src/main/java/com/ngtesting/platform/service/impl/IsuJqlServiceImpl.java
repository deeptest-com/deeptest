package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.support.model.JsonRule;
import com.itfsw.query.builder.support.model.enums.EnumOperator;
import com.itfsw.query.builder.support.model.enums.EnumRuleType;
import com.itfsw.query.builder.support.model.result.SqlQueryResult;
import com.ngtesting.platform.dao.IssueTqlDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.service.IssueService;
import com.ngtesting.platform.service.IsuJqlBuildService;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.IsuJqlService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@Service
public class IsuJqlServiceImpl extends BaseServiceImpl implements IsuJqlService {
    Log logger = LogFactory.getLog(IsuJqlServiceImpl.class);

    public static String[] defaultFilters = new String[] {"project", };

    @Autowired
    IsuJqlBuildService isuJqlBuildService;
    @Autowired
    IssueService issueService;

    @Autowired
    IsuJqlFilterService isuJqlFilterService;

    @Autowired
    IssueTqlDao isuTqlDao;

    @Override
    public List<IsuIssue> query(JsonRule rule, String columns, Integer orgId, Integer projectId) {
        List<IsuIssue> result;

        String conditions;
        if (rule.getRules().size() > 0) {
            SqlQueryResult sqlQueryResult = isuJqlBuildService.buildSqlQuery(JSON.toJSONString(rule));
            conditions = sqlQueryResult.getQuery(true);
        } else {
            conditions = "projectId = " + projectId;
        }

        String reg = "[^,]*Id";

        Pattern r = Pattern.compile(reg);
        Matcher m = r.matcher(columns);
        columns = m.replaceAll("$0,$0Name").replaceAll("IdName","Name");
        logger.info("ReplaceAll: " + columns);

        result = isuTqlDao.query(conditions, columns);

        return result;
    }

    @Override
    public JsonRule buildDefaultJql(Integer orgId, Integer projectId) {
        JsonRule ret = isuJqlBuildService.genJsonRuleRoot();

        JsonRule projectRule = isuJqlBuildService.genJsonRule(
                "projectId", "projectId", "select", projectId.toString(),
                EnumOperator.EQUAL, EnumRuleType.INTEGER);
        ret.getRules().add(projectRule);

        return ret;
    }

}
