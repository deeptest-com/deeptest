package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.service.AlertService;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class AlertServiceImpl extends BaseServiceImpl implements AlertService {

    @Override
    public List<TestAlert> list(Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestAlert.class);

        dc.add(Restrictions.eq("read", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("createTime"));

        List<TestAlert> ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public void scanTestPlan() {
        DetachedCriteria dc = DetachedCriteria.forClass(TestAlert.class);

        dc.add(Restrictions.eq("sent", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("createTime"));

        List<TestAlert> ls = findAllByCriteria(dc);
    }


}

