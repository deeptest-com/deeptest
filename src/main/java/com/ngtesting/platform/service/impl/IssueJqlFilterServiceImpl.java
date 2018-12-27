package com.ngtesting.platform.service.impl;

import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.dao.IssueFieldDao;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.IssueJqlFilterService;
import com.ngtesting.platform.service.intf.ProjectService;
import com.ngtesting.platform.vo.IsuJqlFilter;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class IssueJqlFilterServiceImpl extends BaseServiceImpl implements IssueJqlFilterService {
    Log logger = LogFactory.getLog(IssueJqlFilterServiceImpl.class);

    @Autowired
    ProjectService projectService;

    @Autowired
    IssueDynamicFormService dynamicFormService;

    @Autowired
    ProjectDao projectDao;

    @Autowired
    IssueFieldDao isuFieldDefineDao;

    @Override
    public List<IsuJqlFilter> buildUiFilters(JsonRule rule, Integer orgId, Integer projectId) {
        List<IsuJqlFilter> filtes = new LinkedList<>();

        List<String> filterNameArr = new LinkedList<>();
        iterateRuleName(rule, filterNameArr);

        List<Map> fields = dynamicFormService.fetchOrgField(orgId, projectId, "filter");

        for (Map field : fields) {
            String code = field.get("colCode").toString();

            Boolean filterEnable = filterNameArr.contains(code);
            if (filterEnable) {
                field.put("defaultShowInFilters", true);
            }

            IsuJqlFilter f = new IsuJqlFilter(field);
            filtes.add(f);
        }

        return filtes;
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
