package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.support.model.JsonRule;
import com.itfsw.query.builder.support.model.enums.EnumOperator;
import com.itfsw.query.builder.support.model.enums.EnumRuleType;
import com.itfsw.query.builder.support.model.result.SqlQueryResult;
import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.dao.IsuTqlDao;
import com.ngtesting.platform.vo.IsuJqlColumn;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IssueService;
import com.ngtesting.platform.service.IsuJqlBuildService;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.IsuJqlService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

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
    IsuTqlDao isuTqlDao;

    @Override
    public List<IsuIssue> query(JsonRule rule, String columns, Integer orgId, Integer projectId) {
        List<IsuIssue> result;
        if (rule.getRules().size() > 0) {
            SqlQueryResult sqlQueryResult = isuJqlBuildService.buildSqlQuery(JSON.toJSONString(rule));
            result = issueService.queryByJql(sqlQueryResult.getQuery(true), columns);
        } else {
            result = issueService.queryByProject(projectId, columns);
        }

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
    public List<IsuJqlColumn> buildDefaultColumns(TstUser user) {
        String columnsStr = user.getIssueColumns();

        List<String> ls = new ArrayList<>(Arrays.asList(columnsStr.split(",")));
        List<IsuJqlColumn> cols = new LinkedList<>();

        int i = 0;
        for (String id : ConstantIssue.IssueColumns.keySet()) {
            Boolean enable;
            if (ls.size() > 0) {
                if (ls.contains(id)) {
                    enable = true;
                } else {
                    enable = false;
                }
            } else {
                enable = i++ < 5;
            }

            IsuJqlColumn col = new IsuJqlColumn();
            col.setLabel(ConstantIssue.IssueColumns.get(id));
            col.setId(id);
            col.setDisplay(enable);

            cols.add(col);
        }

        return cols;
    }

}
