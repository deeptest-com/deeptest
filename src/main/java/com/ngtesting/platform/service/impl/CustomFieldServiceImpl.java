package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.ConstantIssue;
import com.ngtesting.platform.dao.*;
import com.ngtesting.platform.model.CustomField;
import com.ngtesting.platform.model.TstCasePriority;
import com.ngtesting.platform.model.TstCaseType;
import com.ngtesting.platform.service.intf.CustomFieldService;
import com.ngtesting.platform.service.intf.ProjectService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.*;

@Service
public class CustomFieldServiceImpl extends BaseServiceImpl implements CustomFieldService {

    @Autowired
    CustomFieldDao customFieldDao;
    @Autowired
    CustomFieldOptionDao customFieldOptionDao;

    @Autowired
    IssueFieldDao issueFieldDao;

    @Autowired
    CaseTypeDao caseTypeDao;
    @Autowired
    CasePriorityDao casePriorityDao;

    @Autowired
    IssuePageElementDao elementDao;

    @Autowired
    ProjectService projectService;

    @Override
    public List<CustomField> list(Integer orgId, String applyTo, String keywords) {
        List<CustomField> ls = customFieldDao.list(orgId, applyTo, keywords);

        return ls;
    }

    @Override
    public CustomField getDetail(Integer id, Integer orgId) {
        return customFieldDao.getDetail(id, orgId);
    }

    @Override
    @Transactional
    public CustomField save(CustomField vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            Integer maxOrder = customFieldDao.getMaxOrdrNumb(orgId, vo.getApplyTo().toString());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            customFieldDao.save(vo);
        } else {
            Integer count = customFieldDao.update(vo);
            if (count == 0) {
                return null;
            }

            elementDao.updateFromCustomField(vo);
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        Integer count = customFieldDao.delete(id, orgId);
        return count > 0;
    }

    @Override
    public Boolean changeOrderPers(Integer id, String act, Integer orgId, String applyTo) {
        CustomField curr = customFieldDao.get(id, orgId);
        if (curr == null) {
            return false;
        }

        CustomField neighbor = null;
        if ("up".equals(act)) {
            neighbor = customFieldDao.getPrev(curr.getOrdr(), orgId, applyTo);
        } else if ("down".equals(act)) {
            neighbor = customFieldDao.getNext(curr.getOrdr(), orgId, applyTo);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        customFieldDao.setOrder(id, neighborOrder, orgId);
        customFieldDao.setOrder(neighbor.getId(), currOrder, orgId);

        return true;
    }

    @Override
    public String getLastUnusedColumn(Integer orgId) {
        List<String> ls = customFieldDao.getLastUnusedColumn(orgId);

        String ret = null;
        for (int i = 1; i <= 20; i++) {
            String prop = "prop" + String.format("%02d", i);
            if (!ls.contains(prop)) {
                ret = prop;
                break;
            }
        }

        return ret;
    }

    @Override // TODO: cached
    public Map<String, Object> fetchProjectFieldForCase(Integer orgId, Integer projectId) {
        List<TstCaseType> caseTypes = caseTypeDao.list(orgId);
        List<TstCasePriority> casePriorities = casePriorityDao.list(orgId);

        Map<String, Object> map = new HashMap<>();
        map.put("type", caseTypes);
        map.put("priority", casePriorities);

        List<Map> fields = customFieldDao.listForCase(orgId);
        for (Map field : fields) {
            map.put(field.get("colCode").toString(), field.get("options"));
        }

        Map<String, Object> ret = new HashMap<>();
        ret.put("fields", fields);
        ret.put("props", map);

        return ret;
    }

    // 获取input及其对应的type，用于表单联动
    @Override
    public Map<String, Map> fetchInputMap() {
        List<Map> inputs = customFieldDao.fetchInputMap();

        Map<String, Map> ret = new LinkedHashMap<>();
        for (Map<String, Object> input : inputs) {
            ret.put(input.get("value").toString(), input);
        }

        return ret;
    }

    @Override
    public Map inputMap() {
        List<Map> ls = customFieldDao.listInput();

        Map<String, String> ret = new LinkedHashMap<>();
        for (Map<String, String> input : ls) {
            ret.put(input.get("value"), input.get("label"));
        }
        return ret;
    }

    @Override
    public Map typeMap() {
        List<Map> ls = customFieldDao.listType();

        Map<String, String> ret = new LinkedHashMap<>();
        for (Map<String, String> input : ls) {
            ret.put(input.get("value"), input.get("label"));
        }
        return ret;
    }

    @Override
    public List<String> listApplyTo() {
        List<String> ls = new LinkedList();
        for (CustomField.FieldApplyTo item : CustomField.FieldApplyTo.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public List<String> listFormat() {
        List<String> ls = new LinkedList();
        for (ConstantIssue.TextFormat item : ConstantIssue.TextFormat.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

}
