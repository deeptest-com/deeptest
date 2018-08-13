package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.WelcomeDao;
import com.ngtesting.platform.model.SysNums;
import com.ngtesting.platform.service.WelcomeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class WelcomeServiceImpl extends BaseServiceImpl implements WelcomeService {
    @Autowired
    WelcomeDao welcomeDao;

    @Override
    public List<SysNums> test() {
        List<SysNums> pos = welcomeDao.test();

        return pos;
    }

}

