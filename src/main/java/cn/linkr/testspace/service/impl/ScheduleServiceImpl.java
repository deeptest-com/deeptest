package cn.linkr.testspace.service.impl;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.testspace.entity.EvtClient;
import cn.linkr.testspace.entity.EvtEvent;
import cn.linkr.testspace.entity.EvtScheduleItem;
import cn.linkr.testspace.entity.EvtSession;
import cn.linkr.testspace.service.ScheduleService;
import cn.linkr.testspace.service.SessionService;
import cn.linkr.testspace.util.BeanUtilEx;
import cn.linkr.testspace.util.DateUtils;
import cn.linkr.testspace.vo.ScheduleItemVo;

@Service
public class ScheduleServiceImpl extends BaseServiceImpl implements
		ScheduleService {
	@Autowired
	SessionService sessionService;

	@Override
	public List<EvtScheduleItem> listScheduleItemsByDate(Long eventId) {

		EvtEvent event = (EvtEvent) get(EvtEvent.class, eventId);

		DetachedCriteria dc = DetachedCriteria.forClass(EvtScheduleItem.class);
		dc.add(Restrictions.eq("eventId", eventId));

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
		
		dc.createAlias("session", "session");  
        dc.add(Restrictions.ne("session.deleted", true));  
		
		dc.addOrder(Order.asc("startDatetime"));
		List<EvtScheduleItem> pos = findAllByCriteria(dc);

		return pos;
	}

	@Override
	public List<EvtScheduleItem> listScheduleItemsBySession(Long eventId) {

		DetachedCriteria dc = DetachedCriteria.forClass(EvtScheduleItem.class);
		dc.add(Restrictions.eq("eventId", eventId));

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
		
		dc.createAlias("session", "session");  
        dc.add(Restrictions.ne("session.deleted", true));  
        
		dc.addOrder(Order.asc("sessionId"));
		dc.addOrder(Order.asc("startDatetime"));
		List<EvtScheduleItem> pos = findAllByCriteria(dc);

		return pos;
	}

	@Override
	public List<ScheduleItemVo> genVosBySession(List<EvtScheduleItem> scheduleItemsBySession, boolean isNest) {

		List<ScheduleItemVo> vosBySession = new LinkedList<ScheduleItemVo>();

		Long currSessionId = null;
		ScheduleItemVo group = new ScheduleItemVo();;
		for (EvtScheduleItem po : scheduleItemsBySession) {
			EvtSession session = po.getSession();
			Long poSessionId = session.getId();

			if (currSessionId != poSessionId) {
				currSessionId = poSessionId;

				group = new ScheduleItemVo();
				group.setId(session.getId());
				group.setAddress(session.getAddress());
				group.setHost(session.getHost());
				group.setName(session.getName());
				group.setItemType("for-group");
				group.setEventId(session.getEventId());
				vosBySession.add(group);
			}

			ScheduleItemVo vo = new ScheduleItemVo();
			
			BeanUtilEx.copyProperties(vo, po);
			vo.setStartDatetimeStr(DateUtils.formatDate(po.getStartDatetime(), "MM/dd HH:mm"));
			vo.setEndDatetimeStr(DateUtils.formatDate(po.getEndDatetime(), "MM-dd HH:mm"));
//			vo.setSubject(po.getSubject());

			// vo.setAddress(po.getSession().getAddress());
			vo.setItemType("for-item");
			
			if (isNest) {
				group.getChildren().add(vo);
			} else {
				vosBySession.add(vo);
			}
		}

		return vosBySession;
	}

	@Override
	public List<ScheduleItemVo> genVosByDate(
			List<EvtScheduleItem> scheduleItemsByDate) {

		List<ScheduleItemVo> vosByDate = new LinkedList<ScheduleItemVo>();

		String dt = "";
		for (EvtScheduleItem po : scheduleItemsByDate) {
			String poDt = DateUtils.formatDate(po.getStartDatetime(),
					"yyyy年MM月dd号 ") + DateUtils.getWeekDay(po.getStartDatetime());

			if (!dt.equals(poDt)) {
				dt = poDt;

				ScheduleItemVo vo = new ScheduleItemVo();
				vo.setName(dt);
				vo.setItemType("for-group");
				vosByDate.add(vo);
			}

			ScheduleItemVo vo = new ScheduleItemVo();
			BeanUtilEx.copyProperties(vo, po);
			vo.setStartDatetimeStr(DateUtils.formatDate(po.getStartDatetime(), "HH:mm"));
			vo.setEndDatetimeStr(DateUtils.formatDate(po.getEndDatetime(), "HH:mm"));
			
			vo.setAddress(po.getSession().getAddress());
			vo.setItemType("for-item");
			vosByDate.add(vo);
		}

		return vosByDate;
	}

	@Override
	public EvtScheduleItem save(ScheduleItemVo vo) {
		if (vo == null) {
			return null;
		}
		
		EvtScheduleItem po = new EvtScheduleItem();
		if (vo.getId() != null) {
			po = (EvtScheduleItem) get(EvtScheduleItem.class, vo.getId());
		}
		
		BeanUtilEx.copyProperties(po, vo);
		
		saveOrUpdate(po);
		return po;
	}

    @Override
    public EvtScheduleItem genPo(JSONObject vo) {
    	EvtScheduleItem po;
    	Long id = vo.getLong("id");
    	if (id == null) {
    		 po = new EvtScheduleItem();
    	} else {
    		 po = (EvtScheduleItem) get(EvtScheduleItem.class, id);
    	}
    	
    	po.setId(vo.getLong("id"));
//    	po.setGuestId(vo.getLong("GuestId"));
    	po.setSubject(vo.getString("subject"));
    	po.setDescr(vo.getString("descr"));
    	String startDate = vo.getString("startDate");
    	String startTime = vo.getString("startTime");
    	String endDate = vo.getString("endDate");
    	String endTime = vo.getString("endTime");
    	
    	if (StringUtils.isNotEmpty(startDate)) {
    		po.setStartDatetime(DateUtils.str2Date(startDate + " " + startTime, "yyyy-MM-dd hh:mm"));
    	}
    	if (StringUtils.isNotEmpty(endDate)) {
    		po.setEndDatetime(DateUtils.str2Date(endDate + " " + endTime, "yyyy-MM-dd hh:mm"));
    	}
    	
        return po;
    }
}
