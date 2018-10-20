package com.ngtesting.platform.service.impl;

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
    public List<IsuJqlFilter> buildUiFilters(JsonRule rule, Integer orgId, Integer projectId) {
        List<IsuJqlFilter> filtes = new LinkedList<>();

        List<String> filterNameArr = new LinkedList<>();
        iterateRuleName(rule, filterNameArr);

        int i = 0;
        for (String id : ConstantIssue.IssueFilters.keySet()) {
            Boolean filterEnable = i++ < 5 || filterNameArr.contains(id);

            String label = ConstantIssue.IssueFilters.get(id);
            IsuJqlFilter f = buildFilter(id, label, orgId, projectId, filterEnable);
            if (f != null) {
                filtes.add(f);
            }
        }

        return filtes;
    }

    @Override
    public IsuJqlFilter buildFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        switch(id){
            case "projectId":
                return buildProjectFilter(id, label, orgId, display);
            case "typeId":
                return buildTypeFilter(id, label, orgId, projectId, display);
            case "statusId":
                return buildStatusFilter(id, label, orgId, projectId, display);
            case "priorityId":
                return buildPriorityFilter(id, label, orgId, projectId, display);
            case "assigneeId":
                return buildAssigneeFilter(id, label, orgId, projectId, display);

            case "creatorId":
                return buildCreatorFilter(id, label, orgId, projectId, display);
            case "reporterId":
                return buildReporterFilter(id, label, orgId, projectId, display);

            case "verId":
                return buildVerFilter(id, label, orgId, projectId, display);
            case "envId":
                return buildEnvFilter(id, label, orgId, projectId, display);
            case "resolutionId":
                return buildResolutionFilter(id, label, orgId, projectId, display);
            case "dueTime":
                return buildDueTimeFilter(id, label, orgId, projectId, display);
            case "resolveTime":
                return buildResolveTimeFilter(id, label, orgId, projectId, display);
            case "comments":
                return buildCommentsFilter(id, label, orgId, projectId, display);
            default:
                return null;
        }
    }

    @Override
    public IsuJqlFilter buildProjectFilter(String id, String label, Integer orgId, Boolean display) {
        List<Map<String, String>> projects = projectDao.queryIdAndName(orgId);

        Map<String, String> values = new HashMap<>();
        for (Map<String, String> prj : projects) {
            values.put(prj.get("id"), prj.get("name"));
        }

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildTypeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildStatusFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildPriorityFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildAssigneeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildCreatorFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildReporterFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildVerFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildEnvFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolutionFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.integer, ConstantIssue.IssueFilterInput.select,  values, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildDueTimeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.date, ConstantIssue.IssueFilterInput.date, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolveTimeFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
                ConstantIssue.IssueFilterType.date, ConstantIssue.IssueFilterInput.date, display);
        return f;
    }

    @Override
    public IsuJqlFilter buildCommentsFilter(String id, String label, Integer orgId, Integer projectId, Boolean display) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(id, label,
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
