package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.model.TstAlert;
import com.ngtesting.platform.model.TstRun;
import com.ngtesting.platform.service.AlertService;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class AlertServiceImpl extends BaseServiceImpl implements AlertService {
    @Override
    public List<TstAlert> list(Integer userId, Boolean isRead) {
        List<TstAlert> pos = scanTestAlert(userId);
        List<TstAlert> vos = genVos(pos);

        return vos;
    }

    @Override
    public List<TstAlert> scanTestAlert(Integer userId) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstAlert.class);
//
//        Date now = new Date();
//        Date startTimeOfToday = DateUtils.GetStartTimeOfDay(now);
//        Date endTimeOfToday = DateUtils.GetEndTimeOfDay(now);
//
//        dc.add(
//            Restrictions.or(
//                // 今天开始
//                Restrictions.and(
//                        Restrictions.isNotNull("startTime"),
//                        Restrictions.ge("startTime", startTimeOfToday),
//                        Restrictions.le("startTime", endTimeOfToday)),
//                // 今天结束
//                Restrictions.and(
//                        Restrictions.isNotNull("endTime"),
//                        Restrictions.ge("endTime", startTimeOfToday),
//                        Restrictions.le("endTime", endTimeOfToday))
//            )
//        );
//
//        dc.add(Restrictions.eq("userId", userId));
////        dc.add(Restrictions.eq("isRead", false));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("startTime"));
//
//        List<TstAlert> pos = findAllByCriteria(dc);
        return null;
    }

    @Override
    public List<TstAlert> genVos(List<TstAlert> pos) {
        List<TstAlert> vos = new LinkedList<>();

        for (TstAlert run: pos) {
            TstAlert vo = genVo(run);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TstAlert genVo(TstAlert po) {
        TstAlert vo = new TstAlert();
//        BeanUtilEx.copyProperties(vo, po);
//        vo.setName(po.getEntityName());
//
//        TestUser user = (TestUser)get(TestUser.class, po.getUserId());
//        TestUser assignee = (TestUser)get(TestUser.class, po.getAssigneeId());
//        vo.setUserName(user.getName());
//        vo.setUserAvatar(user.getAvatar());
//
//        vo.setAssigneeName(assignee.getName());
//        vo.setAssigneeAvatar(assignee.getAvatar());
//
//        Date now = new Date();
//        Long startTimeOfToday = DateUtils.GetStartTimeOfDay(now).getTime();
//        Long endTimeOfToday = DateUtils.GetEndTimeOfDay(now).getTime();
//
//        Date startTime = po.getStartTime();
//        Date endTime = po.getEndTime();
//
//        if (endTime != null && endTime.getTime() >= startTimeOfToday && endTime.getTime() <= endTimeOfToday) {
//            vo.setTitle("测试集" + StringUtil.highlightDict(vo.getName()) + "完成");
//        } else {
//            vo.setTitle("测试集" + StringUtil.highlightDict(vo.getName()) + "开始");
//        }

        return vo;
    }
    @Override
    public void saveAlert(TstRun run) {

//        for (TestUser user : run.getAssignees()) {
//            TstAlert po = getByRun(run.getId());;
//            if (po == null) {
//                po = new TstAlert();
//            }
//
//            po.setType("run");
//            po.setDescr(run.getDescr());
//            po.setEntityId(run.getId());
//            po.setEntityName(run.getName());
//            po.setStatus(run.getStatus().toString());
//            po.setRead(false);
//            po.setUserId(run.getUserId());
//            po.setAssigneeId(user.getId());
//
//            TestPlan plan = run.getPlan();
//            if (plan == null || plan.getId() == null) {
//                plan= (TestPlan)get(TestPlan.class, run.getPlanId());
//            }
//            po.setStartTime(plan.getStartTime());
//            po.setEndTime(plan.getEndTime());
//
//            saveOrUpdate(po);
//        }
    }

    @Override
    public void markAllReadPers(String idStr) {
//        String hql = "update TstAlert alert set alert.isRead=true where alert.id IN (?) " +
//                "AND alert.isRead != true AND alert.deleted != true AND alert.disabled != true";
//
//        List<Long> ids = new LinkedList();
//        for (String str : idStr.split(",")) {
//            ids.add(Long.valueOf(str));
//        }
//        getDao().executeByHql(hql, ids.toArray());
    }

    @Override
    public TstAlert getByRun(Integer id) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstAlert.class);
//
//        dc.add(Restrictions.eq("type", "run"));
//        dc.add(Restrictions.eq("entityId", id));
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.addOrder(Order.asc("id"));
//
//        List<TstAlert> pos = findAllByCriteria(dc);
//        if (pos.size() > 0) {
//            return pos.get(0);
//        } else {
//            return null;
//        }
        return null;
    }

}

