package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.ngtesting.platform.dao.IssueTqlDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.service.intf.*;
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
    IssueDynamicFormService issueDynamicFormService;

    @Autowired
    IssueTqlDao isuTqlDao;

    @Override
    public List<IsuIssue> query(JsonRule rule, String columns, List<Map<String, String>> orders, Integer orgId, Integer projectId) {
        List<IsuIssue> result = new LinkedList<>();

        String conditions;
        if (rule.getRules().size() > 0) {
            SqlQueryResult sqlQueryResult = issueJqlBuildService.buildSqlQuery(JSON.toJSONString(rule));
            conditions = sqlQueryResult.getQuery(true);
            conditions += " AND \"projectId\" = " + projectId;
        } else {
            conditions = "\"projectId\"=" + projectId;
        }

        String[] cols = columns.split(",");
        List<String> customaFields = issueDynamicFormService.listCustomaField(orgId, projectId);
        List<String> newCols = new LinkedList<>();
        for (String col : cols) {
            // select id, "extProp" -> 'age', "extProp" -> 'nickName' from "Test"
            // where id = 5 and "extProp" @> '{"nickName":"aaron"}'::jsonb
            String col1;
            if (!customaFields.contains(col)) {
                col1 = "\"" + col + "\"";
            } else {
                col1 = "\"extProp\" -> '" + col + "'";
            }
            newCols.add(col1);
        }

        result = isuTqlDao.query(conditions, String.join(",", newCols), orders);

        return result;
    }

    @Override
    public JsonRule buildEmptyJql() {
        JsonRule ret = issueJqlBuildService.genJsonRuleRoot();

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
