package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.IssueTqlDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.service.intf.IssueJqlBuildService;
import com.ngtesting.platform.service.intf.IssueJqlFilterService;
import com.ngtesting.platform.service.intf.IssueJqlService;
import com.ngtesting.platform.service.intf.IssueService;
import com.ngtesting.platform.tql.query.builder.support.model.JsonRule;
import com.ngtesting.platform.tql.query.builder.support.model.result.SqlQueryResult;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

@Service
public class IssueJqlServiceImpl extends BaseServiceImpl implements IssueJqlService {
    Log logger = LogFactory.getLog(IssueJqlServiceImpl.class);

    public static String[] defaultFilters = new String[] {"project", };

    @Autowired
    IssueJqlBuildService issueJqlBuildService;
    @Autowired
    IssueService issueService;

    @Autowired
    IssueJqlFilterService issueJqlFilterService;

    @Autowired
    IssueTqlDao isuTqlDao;

    @Override
    public List<IsuIssue> query(JsonRule rule, String columns, List<Map<String, String>> orderBy, Integer orgId, Integer projectId) {
        List<IsuIssue> result = new LinkedList<>();

        String conditions;
        if (rule.getRules().size() > 0) {
            List<JsonRule> rules = rule.getRules();
//            for (JsonRule rule1 : rules) {
//                if (rule1.getField().equals("projectId")) {
//                    rule1.setOperator(EnumOperator.EQUAL.value());
//                    rule1.setValue(projectId);
//                }
//            }
            SqlQueryResult sqlQueryResult = issueJqlBuildService.buildSqlQuery(JSON.toJSONString(rule));
            conditions = sqlQueryResult.getQuery(true);
            conditions += " AND \"projectId\" = " + projectId;
        } else {
            conditions = "\"projectId\"=" + projectId;
        }

        String reg = "[^,]*Id";
        Pattern r = Pattern.compile(reg);
        Matcher m = r.matcher(columns);

//        select id, "extProp" -> 'age', "extProp" -> 'nickName' from "Test"
//        where id = 5 and "extProp" @> '{"nickName":"aaron"}'::jsonb
        if (rule.getBuildIn() != null && rule.getBuildIn()) {
            columns = m.replaceAll("\"$0\"");
        } else {
            columns = m.replaceAll("\"extProp\" -> '$0'");
        }

        result = isuTqlDao.query(conditions, columns, orderBy);

        return result;
    }

    @Override
    public JsonRule buildEmptyJql() {
        JsonRule ret = issueJqlBuildService.genJsonRuleRoot();

//        JsonRule projectRule = issueJqlBuildService.genJsonRule(
//                "projectId", "projectId", "select", "-1",
//                EnumOperator.NOT_EQUAL, EnumRuleType.INTEGER);
//        ret.getRules().save(projectRule);

        return ret;
    }

    @Override
    public List<Map<String, String>> buildDefaultOrderBy() {
        List<Map<String, String>> ls = new LinkedList<>();
        Map map = new HashMap();
        map.put("key", "id");
        map.put("val", "asc");
        ls.add(map);

        map = new HashMap();
        map.put("key", "createTime");
        map.put("val", "asc");
        ls.add(map);

        return ls;
    }

}
