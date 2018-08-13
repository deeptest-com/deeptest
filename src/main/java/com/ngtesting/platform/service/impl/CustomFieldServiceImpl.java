package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CustomFieldDao;
import com.ngtesting.platform.dao.CustomFieldOptionDao;
import com.ngtesting.platform.model.TstCustomField;
import com.ngtesting.platform.service.CustomFieldOptionService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.service.ProjectService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.LinkedList;
import java.util.List;

@Service
public class CustomFieldServiceImpl extends BaseServiceImpl implements CustomFieldService {
    @Autowired
    CustomFieldDao customFieldDao;
    @Autowired
    CustomFieldOptionDao customFieldOptionDao;

    @Autowired
    ProjectService projectService;
    @Autowired
    CustomFieldOptionService customFieldOptionService;

    @Override
    public List<TstCustomField> list(Integer orgId) {
        List<TstCustomField> ls = customFieldDao.list(orgId);

        return ls;
    }

    @Override
    public List<TstCustomField> listForCaseByProject(Integer orgId, Integer projectId) {
        List<TstCustomField> ls = customFieldDao.listForCaseByProject(
                orgId, projectId, TstCustomField.FieldApplyTo.test_case.toString());

        return ls;
    }

    @Override
    public TstCustomField get(Integer customFieldId) {
        return customFieldDao.getDetail(customFieldId);
    }

    @Override
    @Transactional
    public TstCustomField save(TstCustomField vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            Integer maxOrder = customFieldDao.getMaxOrdrNumb(orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            customFieldDao.save(vo);
            if (vo.getType().equals(TstCustomField.FieldType.dropdown)) {
                customFieldOptionDao.saveAll(vo.getId(), vo.getOptions());
            }
        } else {
            customFieldDao.update(vo);
        }

        return vo;
    }

    @Override
    public boolean delete(Integer id) {
        customFieldDao.delete(id);

        return true;
    }

    @Override
    public boolean changeOrderPers(Integer id, String act, Integer orgId) {
        TstCustomField curr = customFieldDao.get(id);
        TstCustomField neighbor = null;
        if ("up".equals(act)) {
            neighbor = customFieldDao.getPrev(curr.getOrdr(), orgId);
        } else if ("down".equals(act)) {
            neighbor = customFieldDao.getNext(curr.getOrdr(), orgId);
        }
        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        customFieldDao.setOrder(id, neighborOrder);
        customFieldDao.setOrder(neighbor.getId(), currOrder);

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

    @Override
    public List<String> listApplyTo() {
        List<String> ls = new LinkedList();
        for (TstCustomField.FieldApplyTo item : TstCustomField.FieldApplyTo.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public List<String> listType() {
        List<String> ls = new LinkedList<String>();
        for (TstCustomField.FieldType item : TstCustomField.FieldType.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

    @Override
    public List<String> listFormat() {
        List<String> ls = new LinkedList();
        for (TstCustomField.FieldFormat item : TstCustomField.FieldFormat.values()) {
            ls.add(item.toString());
        }
        return ls;
    }

}
