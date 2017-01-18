package cn.linkr.events.service.impl;

import java.util.Date;
import java.util.LinkedList;
import java.util.List;

import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtEvent;
import cn.linkr.events.entity.EvtRegisterRecord;
import cn.linkr.events.entity.EvtScheduleItem;
import cn.linkr.events.entity.EvtSession;
import cn.linkr.events.service.SessionService;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.util.DateUtils;
import cn.linkr.events.vo.SessionVo;

@Service
public class SessionServiceImpl extends BaseServiceImpl implements SessionService {
	
	@Override
	public List<EvtSession> listSessionsByEvent(Long eventId) {
		
        DetachedCriteria dc = DetachedCriteria.forClass(EvtSession.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtSession> pos = findAllByCriteria(dc);

		return pos;
	}

	@Override
	public List<EvtSession> listSessionForRegister(Long eventId) {
		DetachedCriteria dc = DetachedCriteria.forClass(EvtSession.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("startTime"));
        List<EvtSession> pos = findAllByCriteria(dc);
		
		return pos;
	}

	@Override
	public boolean isRegister(List<EvtSession> allSessions,
			List<EvtRegisterRecord> registerSessions) {

        boolean alreadyRegister = false;
        for (EvtSession po: allSessions) {
        	
        	for (EvtRegisterRecord po2: registerSessions) {
            	if (po.getId() == po2.getSessionId()) {
            		alreadyRegister = true;
            	}
        	}
        }
		return alreadyRegister;
	}
	
	@Override
	public List<SessionVo> genVos(List<EvtSession> allSessions,
			List<EvtRegisterRecord> registerSessions) {
        List<SessionVo> sessionVos = new LinkedList<SessionVo>();
        
        for (EvtSession po: allSessions) {
        	SessionVo vo = new SessionVo();
        	BeanUtilEx.copyProperties(vo, po);
        	
        	for (EvtRegisterRecord po2: registerSessions) {
            	if (po.getId() == po2.getSessionId()) {
            		vo.setIsRegister(true);
            	}
            	if (po.getStartTime() != null) {
            		vo.setStartTimeStr(DateUtils.formatDate(po.getStartTime(), "yyyy/MM/dd HH:mm"));
            	}
        	}
        	
        	sessionVos.add(vo);
        }
		return sessionVos;
	}
	@Override
	public SessionVo genVo(EvtSession po) {
    	SessionVo vo = new SessionVo();
    	BeanUtilEx.copyProperties(vo, po);
		return vo;
	}

	@Override
	public EvtSession save(SessionVo vo) {
		if (vo == null) {
			return null;
		}
		
		EvtSession po = new EvtSession();
		if (vo.getId() != null) {
			po = (EvtSession) get(EvtSession.class, vo.getId());
		}
		
		BeanUtilEx.copyProperties(po, vo);
		saveOrUpdate(po);
		return po;
	}
	
    @Override
    public EvtSession genPo(JSONObject vo) {
    	EvtSession po;
    	Long id = vo.getLong("id");
    	if (id == null) {
    		 po = new EvtSession();
    	} else {
    		 po = (EvtSession) get(EvtSession.class, id);
    	}
    	
    	po.setId(vo.getLong("id"));
    	po.setName(vo.getString("name"));
    	po.setHost(vo.getString("host"));
    	po.setAddress(vo.getString("address"));
    	
        return po;
    }

	@Override
	public boolean remove(Long id, String type) {
		if ("session".equals(type)) {
			EvtSession po = (EvtSession) get(EvtSession.class, id);
			po.setDeleted(true);
			saveOrUpdate(po);
		} else {
			EvtScheduleItem po = (EvtScheduleItem) get(EvtScheduleItem.class, id);
			po.setDeleted(true);
			saveOrUpdate(po);
		}
		return true;
	}

}
