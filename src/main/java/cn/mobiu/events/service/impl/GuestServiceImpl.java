package cn.mobiu.events.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.service.GuestService;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.GuestVo;

@Service
public class GuestServiceImpl extends BaseServiceImpl implements GuestService {

	@Override
	public List<EvtGuest> list(Long eventId) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtGuest.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtGuest> ls = findAllByCriteria(dc);
        
        return ls;
	}

	@Override
	public List<GuestVo> genVos(List<EvtGuest> pos) {
        List<GuestVo> vos = new LinkedList<GuestVo>();

        for (EvtGuest po: pos) {
        	GuestVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public GuestVo genVo(EvtGuest po) {
		GuestVo vo = new GuestVo();
		BeanUtilEx.copyProperties(vo, po);
		return vo;
	}
}
