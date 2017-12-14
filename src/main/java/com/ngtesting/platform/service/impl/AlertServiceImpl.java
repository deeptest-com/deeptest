package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.util.DateUtils;
import com.ngtesting.platform.vo.TestAlertVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

@Service
public class AlertServiceImpl extends BaseServiceImpl implements AlertService {
    @Override
    public List<TestAlertVo> list(Long userId, Boolean isRead) {
        List<TestRun> pos = scanTestPlan(userId);

        List<TestAlertVo> vos = genVos(pos);

        return vos;
    }

    @Override
    public List<TestRun> scanTestPlan(Long userId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestRun.class);
        dc.createAlias("plan", "plan");

        Date now = new Date();
        Date startTimeOfToday = DateUtils.GetStartTimeOfDay(now);
        Date endTimeOfToday = DateUtils.GetEndTimeOfDay(now);

        dc.add(
             Restrictions.or(
                     // 今天开始
                     Restrictions.and(
                             Restrictions.isNotNull("plan.startTime"),
                             Restrictions.ge("plan.startTime", startTimeOfToday),
                             Restrictions.le("plan.startTime", endTimeOfToday)),
                     // 今天结束
                     Restrictions.and(
                             Restrictions.isNotNull("plan.endTime"),
                             Restrictions.ge("plan.endTime", startTimeOfToday),
                             Restrictions.le("plan.endTime", endTimeOfToday))
             )
        );

        dc.add(Restrictions.eq("userId", userId));

        dc.add(Restrictions.ne("status", TestRun.RunStatus.end));
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        dc.addOrder(Order.asc("plan.startTime"));

        List<TestRun> pos = findAllByCriteria(dc);
        return pos;
    }

    @Override
    public List<TestAlertVo> genVos(List<TestRun> pos) {
        List<TestAlertVo> vos = new LinkedList<>();

        for (TestRun run: pos) {
            TestAlertVo vo = genVo(run);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestAlertVo genVo(TestRun po) {
        TestAlertVo vo = new TestAlertVo();
        BeanUtilEx.copyProperties(vo, po);
        vo.setAvatar(po.getUser().getAvatar());

        return vo;
    }

}

