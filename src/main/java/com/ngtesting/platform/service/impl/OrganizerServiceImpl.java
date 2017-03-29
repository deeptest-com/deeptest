package com.ngtesting.platform.service.impl;

import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.EvtOrganizer;
import com.ngtesting.platform.service.OrganizerService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.OrganizerVo;

@Service
public class OrganizerServiceImpl extends BaseServiceImpl implements OrganizerService {

    @Override
    public List<EvtOrganizer> listByEvent(Long eventId) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtOrganizer.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtOrganizer> organizerPos = findAllByCriteria(dc);
        
        return organizerPos;
    }

	@Override
	public Map<String, List<OrganizerVo>> genOrganizerMap(
			List<EvtOrganizer> organizerPos) {
        Map<String, List<OrganizerVo>> organizerMap = new HashMap<String, List<OrganizerVo>>();
        for (EvtOrganizer organizerPo: organizerPos) {
        	OrganizerVo organizerVo = genVo(organizerPo);
        	
        	String key = organizerPo.getType().toString();
        	if (!organizerMap.containsKey(key)) {
        		organizerMap.put(key, new LinkedList<OrganizerVo>());
        	}
        	organizerMap.get(key).add(organizerVo);
        }
		return organizerMap;
	}
	@Override
	public OrganizerVo genVo(EvtOrganizer po) {
    	OrganizerVo vo = new OrganizerVo();
    	BeanUtilEx.copyProperties(vo, po);
		return vo;
	}
    
}
