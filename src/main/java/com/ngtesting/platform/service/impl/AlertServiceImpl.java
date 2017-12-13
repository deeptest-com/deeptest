package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestAlertVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AlertServiceImpl extends BaseServiceImpl implements AlertService {

    @Override
    public List<TestAlert> list() {
        DetachedCriteria dc = DetachedCriteria.forClass(TestAlert.class);

        dc.add(Restrictions.eq("sent", Boolean.FALSE));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("createTime"));

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
    public void create(TestRun run, UserVo optUser) {
        if(run.getStartTime() != null) {
            TestAlert alert1 = new TestAlert();
            alert1.setTitle("测试集\"" + run.getName() + "\"计划在" +
                    TestAlert.AlertType.run_start.remindDay + "天后开始");
            alert1.setDescr(run.getDescr());
            alert1.setUserId(run.getUserId());
            alert1.setOptUserId(optUser.getId());
            alert1.setStartTime(run.getStartTime());
            saveOrUpdate(alert1);
        }

        if(run.getEndTime() != null) {
            TestAlert alert2 = new TestAlert();
            alert2.setTitle("测试集\"" + run.getName() + "\"计划在" +
                    TestAlert.AlertType.run_end.remindDay + "天后完成");
            alert2.setDescr(run.getDescr());
            alert2.setUserId(run.getUserId());
            alert2.setOptUserId(optUser.getId());
            alert2.setDueTime(run.getEndTime());
            saveOrUpdate(alert2);
        }
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

