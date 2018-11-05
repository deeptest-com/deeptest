package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageElementDao;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.service.IssuePageElementService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IssuePageElementServiceImpl extends BaseServiceImpl implements IssuePageElementService {

	@Autowired
    IssuePageElementDao elementDao;

    @Override
    public void add(IsuPageElement element) {
        Integer maxOrder = elementDao.getMaxFieldOrdr(element.getTabId());
        maxOrder = maxOrder == null? 0: maxOrder;
        element.setOrdr(maxOrder + 1);
        elementDao.add(element);
    }

    @Override
    public boolean remove(Integer id, Integer orgId) {
        Integer count = elementDao.remove(id, orgId);
        if (count == 0) {
            return false;
        }

        return true;
    }

}
