package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueFilterDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IsuFilterService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IsuFilterServiceImpl extends BaseServiceImpl implements IsuFilterService {
    UserDao userDao;
    @Autowired
    IssueFilterDao isuFilterDao;

    @Override
    public Boolean save(Integer caseId, String name, String path, TstUser user) {
        return null;
    }

    @Override
    public Boolean delete(Integer id, TstUser user) {
        return null;
    }
}
