package com.ngtesting.platform.service.impl;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.dao.IssueFieldDao;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.model.IsuFieldDefine;
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

    @Autowired
    IssueFieldDao isuFieldDefineDao;

    @Override
    public List<IsuJqlFilter> buildUiFilters(JsonRule rule, Integer orgId, Integer projectId) {
        List<IsuJqlFilter> filtes = new LinkedList<>();

        List<String> filterNameArr = new LinkedList<>();
        iterateRuleName(rule, filterNameArr);

        List<IsuFieldDefine> fields = isuFieldDefineDao.listFilters();
        int i = 0;
        for (IsuFieldDefine field : fields) {
            String code = field.getCode();

            Boolean filterEnable = filterNameArr.contains(code);
            if (filterEnable) {
                field.setDefaultShowInFilters(filterEnable);
            }

            IsuJqlFilter f = buildFilter(field, orgId, projectId);
            if (f != null) {
                filtes.add(f);
            }
        }

        return filtes;
    }

    @Override
    public IsuJqlFilter buildFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        switch(field.getCode()){
            case "projectId":
                return buildProjectFilter(field, orgId);
            case "typeId":
                return buildTypeFilter(field, orgId, projectId);
            case "statusId":
                return buildStatusFilter(field, orgId, projectId);
            case "priorityId":
                return buildPriorityFilter(field, orgId, projectId);
            case "assigneeId":
                return buildAssigneeFilter(field, orgId, projectId);

            case "creatorId":
                return buildCreatorFilter(field, orgId, projectId);
            case "reporterId":
                return buildReporterFilter(field, orgId, projectId);

            case "verId":
                return buildVerFilter(field, orgId, projectId);
            case "envId":
                return buildEnvFilter(field, orgId, projectId);
            case "resolutionId":
                return buildResolutionFilter(field, orgId, projectId);
            case "dueTime":
                return buildDueTimeFilter(field, orgId, projectId);
            case "resolveTime":
                return buildResolveTimeFilter(field, orgId, projectId);
            case "comments":
                return buildCommentsFilter(field, orgId, projectId);
            default:
                return null;
        }
    }

    @Override
    public IsuJqlFilter buildProjectFilter(IsuFieldDefine field, Integer orgId) {
        List<Map<String, String>> projects = projectDao.queryIdAndName(orgId);

        Map<String, String> values = new HashMap<>();
        for (Map<String, String> prj : projects) {
            values.put(prj.get("id"), prj.get("name"));
        }

        IsuJqlFilter f = new IsuJqlFilter(field, values);
        return f;
    }

    @Override
    public IsuJqlFilter buildTypeFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field, values);
        return f;
    }

    @Override
    public IsuJqlFilter buildStatusFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildPriorityFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildAssigneeFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildCreatorFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildReporterFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildVerFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildEnvFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolutionFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field,  values);
        return f;
    }

    @Override
    public IsuJqlFilter buildDueTimeFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field);
        return f;
    }

    @Override
    public IsuJqlFilter buildResolveTimeFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field);
        return f;
    }

    @Override
    public IsuJqlFilter buildCommentsFilter(IsuFieldDefine field, Integer orgId, Integer projectId) {
        Map<String, String> values = new LinkedHashMap<>();
        values.put("1", "issue"); // TODO: 从数据库获取
        values.put("2", "task");

        IsuJqlFilter f = new IsuJqlFilter(field);
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
