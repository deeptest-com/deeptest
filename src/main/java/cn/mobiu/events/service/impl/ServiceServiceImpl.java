package cn.mobiu.events.service.impl;

import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.mobiu.events.entity.EvtService;
import cn.mobiu.events.entity.EvtService.ServiceType;
import cn.mobiu.events.service.ServiceService;

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
    
}
