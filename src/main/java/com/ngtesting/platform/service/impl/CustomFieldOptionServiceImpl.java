package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CustomFieldDao;
import com.ngtesting.platform.dao.CustomFieldOptionDao;
import com.ngtesting.platform.model.CustomField;
import com.ngtesting.platform.model.CustomFieldOption;
import com.ngtesting.platform.service.intf.IssueCustomFieldOptionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class CustomFieldOptionServiceImpl extends BaseServiceImpl implements IssueCustomFieldOptionService {
    @Autowired
    CustomFieldOptionDao customFieldOptionDao;
    @Autowired
    CustomFieldDao customFieldDao;

    @Override
    public List<CustomFieldOption> list(Integer fieldId, Integer orgId) {
        List<CustomFieldOption> ls = customFieldOptionDao.listByFieldId(fieldId, orgId);
        return ls;
    }

    @Override
    @Transactional
    public CustomFieldOption get(Integer id, Integer fieldId, Integer orgId) {
        return customFieldOptionDao.get(id, fieldId, orgId);
    }

    @Override
    public CustomFieldOption save(CustomFieldOption vo, Integer orgId) {
        vo.setOrgId(orgId);

        if (vo.getId() == null) {
            CustomField field = customFieldDao.get(vo.getFieldId(), orgId);
            if (field == null) {
                return null;
            }

            Integer maxOrder = customFieldOptionDao.getMaxOrder(vo.getFieldId(), orgId);
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);
            customFieldOptionDao.save(vo);
        } else {
            Integer count = customFieldOptionDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer fieldId, Integer orgId) {
        Integer count = customFieldOptionDao.delete(id, orgId);
        return count > 0;
    }

    @Override
    public Boolean changeOrder(Integer id, String act, Integer fieldId, Integer orgId) {
        CustomFieldOption curr = customFieldOptionDao.get(id, fieldId, orgId);
        if (curr == null) {
            return false;
        }

        CustomFieldOption neighbor = null;
        if ("up".equals(act)) {
            neighbor = customFieldOptionDao.getPrev(curr.getOrdr(), fieldId, orgId);
        } else if ("down".equals(act)) {
            neighbor = customFieldOptionDao.getNext(curr.getOrdr(), fieldId, orgId);
        }

        if (neighbor == null) {
            return false;
        }

        Integer currOrder = curr.getOrdr();
        Integer neighborOrder = neighbor.getOrdr();
        customFieldOptionDao.setOrder(id, neighborOrder);
        customFieldOptionDao.setOrder(neighbor.getId(), currOrder);

        return true;
    }

    @Override
    @Transactional
    public Boolean setDefault(Integer id, Integer fieldId, Integer orgId) {
        CustomFieldOption option = customFieldOptionDao.get(id, fieldId, orgId);
        if (option == null) {
            return false;
        }

        Integer count = customFieldOptionDao.removeDefault(fieldId, orgId);
        count = customFieldOptionDao.setDefault(id, fieldId, orgId);

        return count > 0;
    }

}
