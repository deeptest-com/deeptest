package cn.linkr.events.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.linkr.events.entity.EvtBizcard;
import cn.linkr.events.service.BizcardService;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.vo.BizcardVo;

@Service
public class BizCardServiceImpl extends BaseServiceImpl implements BizcardService {

	@Override
	public EvtBizcard getMine(Long clientId) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtBizcard.class);
        dc.add(Restrictions.eq("clientId", clientId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtBizcard> ls = findAllByCriteria(dc);
        
        if (ls.size() > 0) {
        	return ls.get(0);
        }
        return null;
	}

	@Override
	public List<EvtBizcard> listByEvent(Long eventId, Long clientId) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtBizcard.class);
        dc.add(Restrictions.eq("eventId", eventId));
        dc.add(Restrictions.eq("clientId", clientId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtBizcard> ls = findAllByCriteria(dc);
        
        return ls;
	}

	@Override
	public EvtBizcard getDetail(Long bizcardId, Long eventId, Long clientId) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtBizcard.class);
        dc.add(Restrictions.eq("id", bizcardId));
        dc.add(Restrictions.eq("eventId", eventId));
        dc.add(Restrictions.eq("clientId", clientId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtBizcard> ls = findAllByCriteria(dc);
        
        if (ls.size() > 0) {
        	return ls.get(0);
        }
        return null;
	}

	@Override
	public List<BizcardVo> genVos(List<EvtBizcard> pos) {
        List<BizcardVo> vos = new LinkedList<BizcardVo>();
        for (EvtBizcard po: pos) {
        	BizcardVo vo = new BizcardVo();
        	BeanUtilEx.copyProperties(vo, po);
        	vos.add(vo);
        }
		return vos;
	}
	@Override
	public BizcardVo genVo(EvtBizcard po) {
    	BizcardVo vo = new BizcardVo();
    	BeanUtilEx.copyProperties(vo, po);
		return vo;
	}
    
}
