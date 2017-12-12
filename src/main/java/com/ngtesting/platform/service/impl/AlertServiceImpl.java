package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestAlertVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AlertServiceImpl extends BaseServiceImpl implements AlertService {

    @Override
    public List<TestAlert> list(Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestAlert.class);

        if (userId != null) {
            dc.add(Restrictions.eq("runId", userId));
        }

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("pId"));
        dc.addOrder(Order.asc("ordr"));

        List<TestAlert> ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public TestAlertVo getById(Long id) {
        TestAlert po = (TestAlert) get(TestAlert.class, id);
        TestAlertVo vo = genVo(po);

        return vo;
    }

    @Override
    public TestAlert save(JSONObject json) {
        String title = json.getString("title");
        String descr = json.getString("descr");
        Long alertId = json.getLong("id");
        Long userId = json.getLong("userId");

        TestAlert alert;
        if (alertId != null) {
            alert = (TestAlert) get(TestAlert.class, alertId);
        } else {
            alert = new TestAlert();
            alert.setUserId(userId);
        }
        alert.setTitle(title);
        alert.setDescr(descr);
        alert.setUserId(userId);
        saveOrUpdate(alert);

        return alert;
    }

    @Override
    public List<TestAlertVo> genVos(List<TestAlert> pos) {
        List<TestAlertVo> vos = new LinkedList<>();

        for (TestAlert po: pos) {
            TestAlertVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestAlertVo genVo(TestAlert po) {
        TestAlertVo vo = new TestAlertVo();
        BeanUtilEx.copyProperties(vo, po);

        return vo;
    }

}

