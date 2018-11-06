package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageTabDao;
import com.ngtesting.platform.model.IsuPageTab;
import com.ngtesting.platform.service.IssuePageTabService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IssuePageTabServiceImpl extends BaseServiceImpl implements IssuePageTabService {

	@Autowired
    IssuePageTabDao tabDao;

    @Override
    public void add(IsuPageTab tab) {
        tabDao.add(tab);
    }

    @Override
    public IsuPageTab get(Integer tabId, Integer orgId) {
        return tabDao.get(tabId, orgId);
    }

    @Override
    public boolean remove(Integer id, Integer pageId, Integer orgId) {
        Integer count = tabDao.countByPageId(pageId);
        if (count == 1) {
            return false;
        }

        count = tabDao.remove(id, orgId);

        return count > 0;
    }

    @Override
    public void updateName(IsuPageTab tab) {
        tabDao.updateName(tab);
    }
}
