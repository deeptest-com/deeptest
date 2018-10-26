package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssuePageDao;
import com.ngtesting.platform.dao.IssuePageSolutionDao;
import com.ngtesting.platform.model.IsuPageSolution;
import com.ngtesting.platform.service.IssuePageSolutionService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssuePageSolutionServiceImpl extends BaseServiceImpl implements IssuePageSolutionService {
    @Autowired
    UserService userService;

	@Autowired
    IssuePageDao pageDao;
    @Autowired
    IssuePageSolutionDao pageSolutionDao;

    @Override
    public List<IsuPageSolution> list(Integer orgId) {
        return pageSolutionDao.list(orgId);
    }
}
