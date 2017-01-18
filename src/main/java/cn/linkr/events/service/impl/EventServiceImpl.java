package cn.linkr.events.service.impl;

import java.util.Date;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtDocument;
import cn.linkr.events.entity.EvtQa;
import cn.linkr.events.entity.EvtDocument.DocType;
import cn.linkr.events.entity.EvtEvent;
import cn.linkr.events.entity.EvtEvent.EventStatus;
import cn.linkr.events.entity.EvtOrganizer;
import cn.linkr.events.service.DocumentService;
import cn.linkr.events.service.EventService;
import cn.linkr.events.service.OrganizerService;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.util.DateUtils;
import cn.linkr.events.vo.DocumentVo;
import cn.linkr.events.vo.EventVo;
import cn.linkr.events.vo.OrganizerVo;
import cn.linkr.events.vo.Page;
import cn.linkr.events.vo.QaVo;

@Service
public class EventServiceImpl extends BaseServiceImpl implements EventService {
	
	@Autowired
	OrganizerService organizerService;
	
	@Autowired
	DocumentService documentService;
    
    @Override
    public EvtEvent getDetail(Long eventId) {
    	
    	EvtEvent event = (EvtEvent) get(EvtEvent.class, eventId);
//		this.updateStatus(event);
		
		return event;
    }
	
	@Override
	public EvtEvent save(EventVo vo, EvtClient client) {
		EvtEvent po = genPo(vo);
		po.setCreatorId(client.getId());
		po.setCompanyId(client.getCompanyId());
		saveOrUpdate(po);
		return po;
	}
	
//	@Override
//	public void updateStatus(EvtEvent event) {
//		long now = new Date().getTime();
//		EventStatus status = null;
//		if (now > event.getEndTime().getTime()){
//			status = EventStatus.end;
//		} else if (now >= event.getStartTime().getTime() 
//				&& now <= event.getEndTime().getTime()) {
//			status = EventStatus.in_progress;
//		} else if (event.getSignStartTime() != null && now >= event.getSignStartTime().getTime() 
//				&& event.getSignEndTime() !=null && now <= event.getSignEndTime().getTime()) {
//			status = EventStatus.sign;
//		} else if (event.getRegisterStartTime() != null && now >= event.getRegisterStartTime().getTime() 
//				&& event.getRegisterEndTime() != null && now <= event.getRegisterEndTime().getTime()) {
//			status = EventStatus.register;
//		} else {
//			status = EventStatus.not_start;
//		}
//		
//		if (!status.equals(event.getStatus())) {
//			event.setStatus(status);
//			saveOrUpdate(event);
//		}
//	}
    
	@Override
	public List<EventVo> genVos(List<EvtEvent> pos) {
        List<EventVo> vos = new LinkedList<EventVo>();

        for (EvtEvent po: pos) {
        	EventVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
    
    @Override
    public EventVo genVo(EvtEvent po) {
//    	this.updateStatus(po);
        EventVo vo = new EventVo();
		BeanUtilEx.copyProperties(vo, po);
    	
        return vo;
    }
    
    @Override
    public EvtEvent genPo(EventVo vo) {
    	EvtEvent po;
    	Long id = vo.getId();
    	if (id == null) {
    		 po = new EvtEvent();
    	} else {
    		 po = (EvtEvent) get(EvtEvent.class, id);
    	}
    	
    	BeanUtilEx.copyProperties(po, vo);
    	
    	Integer signBefore = vo.getSignBefore();
    	if (signBefore != null) {
    		po.setSignBefore(Integer.valueOf(signBefore));
    		po.setSignStartDatetime(new Date(po.getStartDatetime().getTime() - signBefore * 60 * 60 * 1000));
    		po.setSignEndDatetime(po.getEndDatetime());
    	}
    	
        return po;
    }

	@Override
	public Page list(Long companyId, String statusStr, int currentPage, int itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtEvent.class);
        dc.add(Restrictions.eq("companyId", companyId));
        if (StringUtils.isNotEmpty(statusStr)) {
        	Date now = new Date();
    		EventStatus status = EventStatus.getValue(statusStr);
    		if (status.equals(EventStatus.end)) {
    			dc.add(Restrictions.lt("endDatetime", now));
    		} else if (status.equals(EventStatus.in_progress)) {
    			dc.add(Restrictions.le("startDatetime", now));
    			dc.add(Restrictions.ge("endDatetime", now));
    		} else if (status.equals(EventStatus.sign)) {
    			dc.add(Restrictions.ne("signStartDatetime", null));
    			dc.add(Restrictions.ne("signEndDatetime", null));
    			
    			dc.add(Restrictions.le("signStartDatetime", now));
    			dc.add(Restrictions.ge("signEndDatetime", now));
    		} else if (status.equals(EventStatus.register)) {
    			dc.add(Restrictions.ne("registerStartDatetime", null));
    			dc.add(Restrictions.ne("registerEndDatetime", null));
    			
    			dc.add(Restrictions.le("registerStartDatetime", now));
    			dc.add(Restrictions.ge("registerEndDatetime", now));
    		} else if (status.equals(EventStatus.not_start)) {
    			dc.add(Restrictions.gt("startDatetime", now));
    		} else if (status.equals(EventStatus.cancel)) {
    			dc.add(Restrictions.eq("status", EventStatus.cancel));
    		}
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.desc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);

        return page;
	}
    
}
