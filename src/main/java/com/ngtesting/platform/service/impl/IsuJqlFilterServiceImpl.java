package com.ngtesting.platform.service.impl;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.dao.IssueFieldDao;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.model.IsuFieldDefine;
import com.ngtesting.platform.service.intf.IsuJqlFilterService;
import com.ngtesting.platform.service.intf.ProjectService;
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

        List<IsuFieldDefine> fields = isuFieldDefineDao.listDefaultFilter();
        int i = 0;
        for (IsuFieldDefine field : fields) {
            String code = field.getColCode();

            Boolean filterEnable = filterNameArr.contains(code);
            if (filterEnable) {
                field.setDefaultShowInFilters(true);
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
        switch(field.getColCode()){
            case "projectId":
                return buildProjectFilter(field, orgId);
            default:
                return new IsuJqlFilter(field);
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
