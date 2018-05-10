package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.SysNums;
import com.ngtesting.platform.service.WelcomeService;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class WelcomeServiceImpl extends BaseServiceImpl implements WelcomeService {

    @Override
    public List<SysNums> test() {
        DetachedCriteria dc = DetachedCriteria.forClass(SysNums.class);
        dc.add(Restrictions.eq("key", Long.valueOf(1)));
        List<SysNums> pos = findAllByCriteria(dc);

        return pos;
    }

}

