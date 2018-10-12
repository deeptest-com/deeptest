package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.vo.IsuJqlFilter;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

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

        for (String name : filterNameArr) {
            IsuJqlFilter f = buildFilter(name, orgId, projectId);
            filtes.add(f);
        }

        return filtes;
    }

    @Override
    public IsuJqlFilter buildFilter(String name, Integer orgId, Integer projectId) {
        IsuJqlFilter f = null;
        if (name.equals("project")) {
            f = buildProjectFilter(orgId);
        }

        return f;
    }

    @Override
    public IsuJqlFilter buildProjectFilter(Integer orgId) {
        List<Map<String, String>> projects = projectDao.queryIdAndName(orgId);

        Map<String, String> values = new HashMap<>();
        for (Map<String, String> prj : projects) {
            values.put(prj.get("id"), prj.get("name"));
        }

        IsuJqlFilter f = new IsuJqlFilter("project", "项目", values);
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
