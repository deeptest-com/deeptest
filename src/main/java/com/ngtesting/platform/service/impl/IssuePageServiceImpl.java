package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.dao.IssuePageSolutionDao;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.service.IssuePageService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

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
}
