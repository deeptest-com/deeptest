package cn.linkr.events.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.linkr.events.entity.EvtGuest;
import cn.linkr.events.entity.EvtQa;
import cn.linkr.events.service.QaService;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.vo.QaVo;

@Service
public class QaServiceImpl extends BaseServiceImpl implements QaService {

	@Override
	public List<EvtQa> list(Long eventId, Long clientId) {
		
        DetachedCriteria dc = DetachedCriteria.forClass(EvtQa.class);
        dc.add(Restrictions.eq("eventId", eventId));
        dc.add(Restrictions.eq("authorId", clientId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtQa> ls = findAllByCriteria(dc);
        
        return ls;
	}

	@Override
	public void save(Long eventId, String content) {
		EvtQa qa = new EvtQa();
		qa.setEventId(eventId);
		qa.setQuestion(content);
		
		saveOrUpdate(qa);
	}

	@Override
	public List<QaVo> genVos(List<EvtQa> pos) {
        List<QaVo> vos = new LinkedList<QaVo>();

        for (EvtQa po: pos) {
        	QaVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
	@Override
	public QaVo genVo(EvtQa po) {
    	QaVo vo = new QaVo();
    	BeanUtilEx.copyProperties(vo, po);
		return vo;
	}
	
}
