package com.ngtesting.platform.service.impl;


import com.ngtesting.platform.dao.IssueTypeDao;
import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.service.IssueTypeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class IssueTypeServiceImpl extends BaseServiceImpl implements IssueTypeService {
    @Autowired
    IssueTypeDao issueTypeDao;

    @Override
    public List<IsuType> list(Integer orgId) {
        List<IsuType> ls = issueTypeDao.list(orgId);

        return ls;
    }
}
