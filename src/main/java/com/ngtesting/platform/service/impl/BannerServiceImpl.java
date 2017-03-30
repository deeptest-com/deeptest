package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.EvtBanner;
import com.ngtesting.platform.service.BannerService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.BannerVo;
import com.ngtesting.platform.vo.Page;

@Service
public class BannerServiceImpl extends BaseServiceImpl implements BannerService {

    @Override
    public List<EvtBanner> listByEvent(Long eventId) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtBanner.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtBanner> docPos = findAllByCriteria(dc);
        
        return docPos;
    }
    
	@Override
	public Page listByPage(long eventId, int currentPage, int itemsPerPage) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtBanner.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
        
        return page;
	}
	
	@Override
	public EvtBanner save(BannerVo vo) {
		if (vo == null) {
			return null;
		}
		
		EvtBanner po = new EvtBanner();
		if (vo.getId() != null) {
			po = (EvtBanner) get(EvtBanner.class, vo.getId());
		}
		
		po.setEventId(vo.getEventId());
		po.setTitle(vo.getTitle());
		po.setUri(vo.getUri());
		
		saveOrUpdate(po);
		return po;
	}
	
	@Override
	public boolean remove(Long id) {
		EvtBanner po = (EvtBanner) get(EvtBanner.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);
		
		return true;
	}

	@Override
	public List<BannerVo> genVos(List<EvtBanner> pos) {
        List<BannerVo> vos = new LinkedList<BannerVo>();
        for (EvtBanner po: pos) {
        	BannerVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
	@Override
	public BannerVo genVo(EvtBanner po) {

		BannerVo vo = new BannerVo();
    	BeanUtilEx.copyProperties(vo, po);

		return vo;
	}
    
}
