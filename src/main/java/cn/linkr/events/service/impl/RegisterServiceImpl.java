package cn.linkr.events.service.impl;

import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.linkr.events.entity.EvtRegisterRecord;
import cn.linkr.events.service.RegisterService;

@Service
public class RegisterServiceImpl extends BaseServiceImpl implements RegisterService {

	@Override
	public void register(String eventId, String sessionIds) {
		
	}
	
	@Override
	public List<EvtRegisterRecord> listRegisterSession(Long eventId, Long clientId) {
		
        DetachedCriteria dc = DetachedCriteria.forClass(EvtRegisterRecord.class);
        dc.add(Restrictions.eq("eventId", eventId));
        dc.add(Restrictions.eq("clientId", clientId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtRegisterRecord> pos = findAllByCriteria(dc);
		
        return pos;
	}

	@Override
	public long getRegisterNumb(Long eventId) {
		String hql = "select count(rcd) from EvtRegisterRecord rcd where eventId = ?"
				+ " and rcd.deleted != true and rcd.deleted != true";
		
		long count = (Long) getByHQL(hql, eventId);
		return count;
	}

}
