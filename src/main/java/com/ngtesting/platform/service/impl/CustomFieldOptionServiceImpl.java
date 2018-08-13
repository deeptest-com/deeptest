package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CustomFieldOptionDao;
import com.ngtesting.platform.model.TstCustomFieldOption;
import com.ngtesting.platform.service.CustomFieldOptionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CustomFieldOptionServiceImpl extends BaseServiceImpl implements CustomFieldOptionService {
    @Autowired
    CustomFieldOptionDao customFieldOptionDao;

    @Override
    public List<TstCustomFieldOption> listVos(Integer fieldId) {
        List<TstCustomFieldOption> ls = customFieldOptionDao.listByFieldId(fieldId);
        return ls;
    }

    @Override
    public TstCustomFieldOption save(TstCustomFieldOption vo) {
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
    public boolean delete(Integer id) {
        customFieldOptionDao.delete(id);
        return true;
    }

    @Override
    public boolean changeOrderPers(Integer id, String act, Integer fieldId) {
        TstCustomFieldOption curr = customFieldOptionDao.get(id);
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
