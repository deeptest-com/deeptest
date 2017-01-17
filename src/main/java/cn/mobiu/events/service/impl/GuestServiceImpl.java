package cn.mobiu.events.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.alibaba.fastjson.JSONObject;

import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.entity.EvtScheduleItem;
import cn.mobiu.events.entity.EvtSession;
import cn.mobiu.events.service.GuestService;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.GuestVo;
import cn.mobiu.events.vo.Page;
import cn.mobiu.events.vo.SessionVo;

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
	public Page list(Long eventId, int currentPage, int itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtGuest.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
		
		return page;
	}

	@Override
	public EvtGuest save(GuestVo vo) {
		if (vo == null) {
			return null;
		}
		
		EvtGuest po = new EvtGuest();
		if (vo.getId() != null) {
			po = (EvtGuest) get(EvtGuest.class, vo.getId());
		}
		
		po.setEventId(vo.getEventId());
		po.setName(vo.getName());
		po.setTitle(vo.getTitle());
		po.setDescr(vo.getDescr());
		po.setAvatar(vo.getAvatar());
		
		saveOrUpdate(po);
		return po;
	}
	
	@Override
	public boolean remove(Long id) {
		EvtGuest po = (EvtGuest) get(EvtGuest.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
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
	
    @Override
    public EvtGuest genPo(GuestVo vo) {
    	EvtGuest po;
    	Long id = vo.getId();
    	if (id == null) {
    		 po = new EvtGuest();
    	} else {
    		 po = (EvtGuest) get(EvtGuest.class, id);
    	}
    	BeanUtilEx.copyProperties(po, vo);
    	
        return po;
    }

}
