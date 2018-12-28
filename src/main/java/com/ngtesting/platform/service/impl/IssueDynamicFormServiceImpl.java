package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDynamicFormDao;
import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssueDynamicFormServiceImpl extends BaseServiceImpl implements IssueDynamicFormService {
    @Autowired
    IssueDynamicFormDao dynamicFormDao;

    @Override
    public List<IsuField> listNotUsedField(Integer orgId, Integer projectId, Integer pageId) {
        List<IsuField> fields = dynamicFormDao.listNotUsedField(orgId, projectId, pageId, "elem");

        return fields;
    }

    @Override
    public Map<String, List<Map>> genIssuePropMap(Integer orgId, Integer projectId) {
        Map<String, List<Map>> map = new LinkedHashMap<>();

        List<Map> fields = fetchOrgField(orgId, projectId, "elem");

        for (Map field : fields) {
            map.put(field.get("colCode").toString(), (List)field.get("options"));
        }

        return map;
    }

    @Override
    public Map<String, Object> genIssueBuldInPropValMap(Integer orgId, Integer projectId) {
        Map<String, Object> map = new LinkedHashMap<>();

        List<Map> fields = fetchOrgField(orgId, projectId, "elem");

		for (Map field : fields) {
		    if (!"1".equals(field.get("buildIn").toString()) || field.get("options") == null) {
                continue;
            }

            Map optionMap = new LinkedHashMap();
            for (Map option: (List<Map>)field.get("options")) {
                    optionMap.put(option.get("id"), option.get("label"));
            }
			map.put(field.get("colCode").toString(), optionMap);
		}

        return map;
    }

    @Override // TODO: cached
    public List<Map> fetchOrgField(Integer orgId, Integer projectId, String sort) {
        List<Map> fields = dynamicFormDao.fetchOrgField(orgId, projectId, sort);
        return fields;
    }

}
