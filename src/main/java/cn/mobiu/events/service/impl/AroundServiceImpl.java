package cn.mobiu.events.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.mobiu.events.entity.EvtAround;
import cn.mobiu.events.entity.EvtAround.AroundType;
import cn.mobiu.events.service.AroundService;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.AroundVo;

@Service
public class AroundServiceImpl extends BaseServiceImpl implements AroundService {
	
    @Override
    public List<EvtAround> list(Long eventId, AroundType type) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtAround.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        if (type != null) {
        	dc.add(Restrictions.eq("type", type));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtAround> ls = findAllByCriteria(dc);
        
        return ls;
    }

	@Override
	public List<AroundVo> genVos(List<EvtAround> pos) {
        List<AroundVo> vos = new LinkedList<AroundVo>();
        for (EvtAround po: pos) {
        	vos.add(genVo(po));
        }
		return vos;
	}
	@Override
	public AroundVo genVo(EvtAround po) {
    	AroundVo vo = new AroundVo();
    	BeanUtilEx.copyProperties(vo, po);
		return vo;
	}
    
}
