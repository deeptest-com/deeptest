package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.TstCustomFieldDao;
import com.ngtesting.platform.dao.TstCustomFieldOptionDao;
import com.ngtesting.platform.model.TstCustomField;
import com.ngtesting.platform.model.TstCustomFieldOption;
import com.ngtesting.platform.service.TestCustomFieldOptionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TestCustomFieldOptionServiceImpl extends BaseServiceImpl implements TestCustomFieldOptionService {
    @Autowired
    TstCustomFieldOptionDao customFieldOptionDao;
    @Autowired
    TstCustomFieldDao customFieldDao;

    @Override
    public List<TstCustomFieldOption> listVos(Integer fieldId) {
        List<TstCustomFieldOption> ls = customFieldOptionDao.listByFieldId(fieldId);
        return ls;
    }

    @Override
    public TstCustomFieldOption save(TstCustomFieldOption vo, Integer orgId) {
        TstCustomField field = customFieldDao.get(vo.getFieldId(), orgId);
        if (field == null) {
            return null;
        }

        if (vo.getId() == null) {
            Integer maxOrder = customFieldOptionDao.getMaxOrder(vo.getFieldId());
            if (maxOrder == null) {
                maxOrder = 0;
            }
            vo.setOrdr(maxOrder + 10);

            customFieldOptionDao.save(vo);
        } else {
            customFieldOptionDao.update(vo);
        }

        return vo;
    }

    @Override
    public Boolean delete(Integer id, Integer orgId) {
        TstCustomFieldOption option = customFieldOptionDao.get(id);
        if (option == null) {
            return false;
        }

        TstCustomField field = customFieldDao.get(option.getFieldId(), orgId);
        if (field == null) {
            return false;
        }

        customFieldOptionDao.delete(id);
        return true;
    }

    @Override
    public Boolean changeOrder(Integer id, String act, Integer fieldId, Integer orgId) {
        TstCustomFieldOption curr = customFieldOptionDao.get(id);
        if (curr == null) {
            return false;
        }

        TstCustomField field = customFieldDao.get(curr.getFieldId(), orgId);
        if (field == null) {
            return false;
        }

        TstCustomFieldOption neighbor = null;
        if ("up".equals(act)) {
            neighbor = customFieldOptionDao.getPrev(curr.getOrdr(), fieldId);
        } else if ("down".equals(act)) {
            neighbor = customFieldOptionDao.getNext(curr.getOrdr(), fieldId);
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
}
