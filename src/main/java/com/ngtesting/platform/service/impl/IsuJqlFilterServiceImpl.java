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
//        iterateRuleName(rule, filterNameArr);

        int i = 0;
        for (String name : ConstantIssue.IssueFilters.keySet()) {
            IsuJqlFilter f = buildFilter(name, orgId, projectId, i < 5);
            filtes.add(f);
        }

//        for (String name : filterNameArr) {
//            if (ConstantIssue.IssueFilters.keySet().contains(name)) {
//                continue;
//            }
//            IsuJqlFilter f = buildFilter(name, orgId, projectId, false);
//            filtes.add(f);
//        }

        return filtes;
    }

    @Override
    public IsuJqlFilter buildFilter(String name, Integer orgId, Integer projectId, Boolean display) {
        switch(name){
            case "project":
                return buildProjectFilter(orgId, display);
            case "type":
                return buildTypeFilter(orgId, projectId, display);
            case "status":
                return buildStatusFilter(orgId, projectId, display);
            case "priority":
                return buildPriorityFilter(orgId, projectId, display);
            case "assignee":
                return buildAssigneeFilter(orgId, projectId, display);

            case "creator":
                return buildCreatorFilter(orgId, projectId, display);
            case "reporter":
                return buildReporterFilter(orgId, projectId, display);

            case "ver":
                return buildVerFilter(orgId, projectId, display);
            case "env":
                return buildEnvFilter(orgId, projectId, display);
            case "resolution":
                return buildResolutionFilter(orgId, projectId, display);
            case "dueTime":
                return buildDueTimeFilter(orgId, projectId, display);
            case "resolveTime":
                return buildResolveTimeFilter(orgId, projectId, display);
            case "comments":
                return buildCommentsFilter(orgId, projectId, display);
            default:
                return null;
        }
    }

    @Override
    public IsuJqlFilter buildProjectFilter(Integer orgId, Boolean display) {
        List<Map<String, String>> projects = projectDao.queryIdAndName(orgId);

        Map<String, String> values = new HashMap<>();
        for (Map<String, String> prj : projects) {
            values.put(prj.get("id"), prj.get("name"));
        }

        IsuJqlFilter f = new IsuJqlFilter("projectId", "项目",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildTypeFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("typeId", "类型",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildStatusFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("statusId", "状态",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildPriorityFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("priorityId", "优先级",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildAssigneeFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("assigneeId", "经办人",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildCreatorFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("creatorId", "创建人",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildReporterFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("reporterId", "报告人",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildVerFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("verId", "版本",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildEnvFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("envId", "环境",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolutionFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("resolutionId", "解决结果",
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildDueTimeFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("dueTime", "截止日期",
                ConstantIssue.IssueFilterType.date, ConstantIssue.IssueFilterInput.date, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolveTimeFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("resolveTime", "解决时间",
                ConstantIssue.IssueFilterType.date, ConstantIssue.IssueFilterInput.date, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildCommentsFilter(Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter("comments", "注释",
                ConstantIssue.IssueFilterType.string, ConstantIssue.IssueFilterInput.string, display);
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
