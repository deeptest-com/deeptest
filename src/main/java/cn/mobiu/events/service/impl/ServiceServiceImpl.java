package cn.mobiu.events.service.impl;

import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.entity.EvtService;
import cn.mobiu.events.entity.EvtService.ServiceType;
import cn.mobiu.events.service.ServiceService;
import cn.mobiu.events.vo.GuestVo;
import cn.mobiu.events.vo.ServiceVo;

@Service
public class ServiceServiceImpl extends BaseServiceImpl implements ServiceService {

    @Override
    public List<EvtService> list(Long eventId, ServiceType type) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtService.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        if (type != null) {
        	dc.add(Restrictions.eq("type", type));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtService> ls = findAllByCriteria(dc);
        
        return ls;
    }
    
    @Override
    public List<EvtService> listForEdit(Long eventId, ServiceType type) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtService.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        if (type != null) {
        	dc.add(Restrictions.eq("type", type));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.addOrder(Order.asc("disabled"));
        dc.addOrder(Order.asc("id"));
        List<EvtService> ls = findAllByCriteria(dc);
        
        return ls;
    }
    
	@Override
	public EvtService save(ServiceVo vo) {
		if (vo == null) {
			return null;
		}
		
		EvtService po = new EvtService();
		if (vo.getId() != null) {
			po = (EvtService) get(EvtService.class, vo.getId());
		}
		
		po.setEventId(vo.getEventId());
		po.setSubject(vo.getSubject());
		po.setDescr(vo.getDescr());
		
		saveOrUpdate(po);
		return po;
	}
	
	@Override
	public boolean disablePers(Long id, String action) {
		EvtService po = (EvtService) get(EvtService.class, id);
		if ("disable".equals(action)) {
			po.setDisabled(true);
		} else {
			po.setDisabled(false);
		}
		saveOrUpdate(po);
		
		return true;
	}
    
}
