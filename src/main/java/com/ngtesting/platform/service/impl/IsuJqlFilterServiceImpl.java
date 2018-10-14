package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.vo.IsuJqlFilter;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

@Service
public class IsuJqlFilterServiceImpl extends BaseServiceImpl implements IsuJqlFilterService {
    Log logger = LogFactory.getLog(IsuJqlFilterServiceImpl.class);

    @Autowired
    ProjectService projectService;

    @Autowired
    ProjectDao projectDao;

    @Override
    public List<IsuJqlFilter> buildUiFilters(String jql, Integer orgId, Integer projectId) {
        List<IsuJqlFilter> filtes = new LinkedList<>();

        JsonRule rule = JSON.parseObject(jql, JsonRule.class);

        List<String> filterNameArr = new LinkedList<>();
        iterateRuleName(rule, filterNameArr);

        for (String name : ConstantIssue.IssueFilters.keySet()) {
            IsuJqlFilter f = buildFilter(name, orgId, projectId);
            filtes.add(f);
        }

        for (String name : filterNameArr) {
            if (ConstantIssue.IssueFilters.keySet().contains(name)) {
                continue;
            }
            IsuJqlFilter f = buildFilter(name, orgId, projectId);
            filtes.add(f);
        }

        return filtes;
    }

    @Override
    public IsuJqlFilter buildFilter(String name, Integer orgId, Integer projectId) {
        switch(name){
            case "project":
                return buildProjectFilter(orgId);
            case "type":
                return buildTypeFilter(orgId, projectId);
            case "status":
                return buildStatusFilter(orgId, projectId);
            case "priority":
                return buildPriorityFilter(orgId, projectId);
            case "assignee":
                return buildAssigneeFilter(orgId, projectId);

            case "creator":
                return buildCreatorFilter(orgId, projectId);
            case "reporter":
                return buildReporterFilter(orgId, projectId);

            case "ver":
                return buildVerFilter(orgId, projectId);
            case "env":
                return buildEnvFilter(orgId, projectId);
            case "resolution":
                return buildResolutionFilter(orgId, projectId);
            case "dueTime":
                return buildDueTimeFilter(orgId, projectId);
            case "resolveTime":
                return buildResolveTimeFilter(orgId, projectId);
            case "comments":
                return buildCommentsFilter(orgId, projectId);
            default:
                return null;
        }
    }

    @Override
    public IsuJqlFilter buildProjectFilter(Integer orgId) {
        List<Map<String, String>> projects = projectDao.queryIdAndName(orgId);

        Map<String, String> values = new HashMap<>();
        for (Map<String, String> prj : projects) {
            values.put(prj.get("id"), prj.get("name"));
        }

        IsuJqlFilter f = new IsuJqlFilter("project", "项目",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildTypeFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("type", "类型",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildStatusFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("status", "状态",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildPriorityFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("priority", "优先级",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildAssigneeFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("assignee", "经办人",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildCreatorFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("creator", "创建人",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildReporterFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("reporter", "报告人",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildVerFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("ver", "版本",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildEnvFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("env", "环境",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolutionFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("resolution", "解决结果",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildDueTimeFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("dueTime", "截止日期",
                ConstantIssue.IssueFilterType.date, ConstantIssue.IssueFilterInput.date);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolveTimeFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("resolveTime", "解决时间",
                ConstantIssue.IssueFilterType.date, ConstantIssue.IssueFilterInput.date);
        return f;
    }

    @Override
    public IsuJqlFilter buildCommentsFilter(Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("comments", "注释",
                ConstantIssue.IssueFilterType.string, ConstantIssue.IssueFilterInput.string);
        return f;
    }

    @Override
    public void iterateRuleName(JsonRule rule, List<String> names) {
        if (rule.getId() != null) {
            names.add(rule.getId());
        }

        if (rule.isGroup()) {
            for (JsonRule child : rule.getRules()) {
                iterateRuleName(child, names);
            }
        }
    }
}
