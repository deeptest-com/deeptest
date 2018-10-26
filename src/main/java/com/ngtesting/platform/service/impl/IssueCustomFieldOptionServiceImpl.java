package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueCustomFieldDao;
import com.ngtesting.platform.dao.IssueCustomFieldOptionDao;
import com.ngtesting.platform.model.IsuCustomField;
import com.ngtesting.platform.model.IsuCustomFieldOption;
import com.ngtesting.platform.service.IssueCustomFieldOptionService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueCustomFieldOptionServiceImpl extends BaseServiceImpl implements IssueCustomFieldOptionService {
    @Autowired
    IssueCustomFieldOptionDao customFieldOptionDao;
    @Autowired
    IssueCustomFieldDao customFieldDao;

    @Override
    public List<IsuCustomFieldOption> listVos(Integer fieldId) {
        List<IsuCustomFieldOption> ls = customFieldOptionDao.listByFieldId(fieldId);
        return ls;
    }

    @Override
    public IsuCustomFieldOption save(IsuCustomFieldOption vo, Integer orgId) {
        IsuCustomField field = customFieldDao.get(vo.getFieldId(), orgId);
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
        IsuCustomFieldOption option = customFieldOptionDao.get(id);
        if (option == null) {
            return false;
        }

        IsuCustomField field = customFieldDao.get(option.getFieldId(), orgId);
        if (field == null) {
            return false;
        }

        customFieldOptionDao.delete(id);
        return true;
    }

    @Override
    public Boolean changeOrder(Integer id, String act, Integer fieldId, Integer orgId) {
        IsuCustomFieldOption curr = customFieldOptionDao.get(id);
        if (curr == null) {
            return false;
        }

        IsuCustomField field = customFieldDao.get(curr.getFieldId(), orgId);
        if (field == null) {
            return false;
        }

        IsuCustomFieldOption neighbor = null;
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
