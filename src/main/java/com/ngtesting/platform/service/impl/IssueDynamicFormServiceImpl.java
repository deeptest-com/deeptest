package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDynamicFormDao;
import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import org.apache.ibatis.annotations.Param;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class IssueDynamicFormServiceImpl extends BaseServiceImpl implements IssueDynamicFormService {
    @Autowired
    IssueDynamicFormDao dynamicFormDao;

    @Override
    public List<IsuField> listTabNotUsedField(Integer orgId, Integer projectId, Integer tabId) {
        List<IsuField> fields = dynamicFormDao.listTabNotUsedField(tabId, projectId, orgId);

        return fields;
    }

    @Override // TODO: cached
    public Map<String, Object> fetchOrgField(Integer orgId, Integer projectId) {
        Map<String, Object> map = new HashMap<>();

        List<Map> fields = dynamicFormDao.fetchOrgField(orgId, projectId);

        for (Map field : fields) {
            map.put(field.get("code").toString(), field.get("options"));
        }

        return map;
    }

}
