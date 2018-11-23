package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageTabDao;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.model.IsuPageTab;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.IssuePageTabService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Map;

@Service
public class IssuePageTabServiceImpl extends BaseServiceImpl implements IssuePageTabService {

	@Autowired
    IssuePageTabDao tabDao;

    @Autowired
    IssueDynamicFormService dynamicFormService;

    @Override
    public void add(IsuPageTab tab) {
        tabDao.add(tab);
    }

    @Override
    public IsuPageTab get(Integer tabId, Integer orgId) {
        IsuPageTab tab = tabDao.get(tabId, orgId);

        return tab;
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
