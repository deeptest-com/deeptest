package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.dao.IssuePageSolutionDao;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuPageElement;
import com.ngtesting.platform.model.IsuPageTab;
import com.ngtesting.platform.service.IssuePageService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssuePageServiceImpl extends BaseServiceImpl implements IssuePageService {
    @Autowired
    UserService userService;

	@Autowired
    IssuePageDao pageDao;
    @Autowired
    IssuePageSolutionDao pageSolutionDao;

    @Override
    public List<IsuPage> list(Integer orgId) {
        return pageDao.list(orgId);
    }

    @Override
    public IsuPage get(Integer pageId, Integer orgId) {
        return pageDao.get(pageId, orgId);
    }

    @Override
    @Transactional
    public IsuPage save(IsuPage vo, Integer orgId) {
        if (vo.getId() == null) {

            vo.setOrgId(orgId);
            pageDao.save(vo);

            pageDao.saveDefaultTab(vo.getId(), orgId);
        } else {
            Integer count = pageDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        return vo;
    }

    @Override
    public boolean delete(Integer id, Integer orgId) {
        Integer count = pageDao.delete(id, orgId);
        if (count == 0) {
            return false;
        }

        return true;
    }

    @Override
    public void addTab(IsuPageTab tab) {
        pageDao.addTab(tab);
    }

    @Override
    public void addField(IsuPageElement element) {
        pageDao.addField(element);
    }

    @Override
    public IsuPageTab getTab(Integer tabId, Integer orgId) {
        return pageDao.getTab(tabId, orgId);
    }
}
