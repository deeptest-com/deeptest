package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.WelcomeDao;
import com.ngtesting.platform.model.CustomField;
import com.ngtesting.platform.service.intf.WelcomeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class WelcomeServiceImpl extends BaseServiceImpl implements WelcomeService {
    @Autowired
    WelcomeDao welcomeDao;

    @Override
    public List<CustomField> test() {
        List<CustomField> pos = welcomeDao.test();

        return pos;
    }

}

