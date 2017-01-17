package cn.mobiu.events.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import cn.mobiu.events.entity.EvtDocument;
import cn.mobiu.events.entity.EvtDocument.DocType;
import cn.mobiu.events.service.DocumentService;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.DocumentVo;
import cn.mobiu.events.vo.Page;

@Service
public class DocumentServiceImpl extends BaseServiceImpl implements DocumentService {

    @Override
    public List<EvtDocument> listByEvent(Long eventId, DocType type) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtDocument.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        if (type != null) {
        	dc.add(Restrictions.eq("docType", type));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<EvtDocument> docPos = findAllByCriteria(dc);
        
        return docPos;
    }
    
	@Override
	public Page listByPage(long eventId, int currentPage, int itemsPerPage, DocType type) {
        DetachedCriteria dc = DetachedCriteria.forClass(EvtDocument.class);
        dc.add(Restrictions.eq("eventId", eventId));
        
        if (type != null) {
        	dc.add(Restrictions.eq("docType", type));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
        
        return page;
	}

	@Override
	public List<DocumentVo> genVos(List<EvtDocument> pos) {
        List<DocumentVo> vos = new LinkedList<DocumentVo>();
        for (EvtDocument po: pos) {
        	DocumentVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
	@Override
	public DocumentVo genVo(EvtDocument po) {

    	DocumentVo vo = new DocumentVo();
    	BeanUtilEx.copyProperties(vo, po);

		return vo;
	}
    
}
