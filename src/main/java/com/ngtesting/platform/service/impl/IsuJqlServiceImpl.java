package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.support.model.JsonRule;
import com.itfsw.query.builder.support.model.enums.EnumOperator;
import com.itfsw.query.builder.support.model.enums.EnumRuleType;
import com.itfsw.query.builder.support.model.result.SqlQueryResult;
import com.ngtesting.platform.dao.IssueTqlDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.service.intf.IssueService;
import com.ngtesting.platform.service.intf.IsuJqlBuildService;
import com.ngtesting.platform.service.intf.IsuJqlFilterService;
import com.ngtesting.platform.service.intf.IsuJqlService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;
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
    public List<IsuIssue> query(JsonRule rule, String columns, List<Map<String, String>> orderBy, Integer orgId, Integer projectId) {
        List<IsuIssue> result;

        String conditions;
        if (rule.getRules().size() > 0) {
            SqlQueryResult sqlQueryResult = isuJqlBuildService.buildSqlQuery(JSON.toJSONString(rule));
            conditions = sqlQueryResult.getQuery(true);
        } else {
            conditions = "projectId=" + projectId;
        }

        String reg = "[^,]*Id";
        Pattern r = Pattern.compile(reg);
        Matcher m = r.matcher(columns);
        columns = m.replaceAll("$0,$0Name").replaceAll("IdName","Name");

        result = isuTqlDao.query(conditions, columns, orderBy);

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
