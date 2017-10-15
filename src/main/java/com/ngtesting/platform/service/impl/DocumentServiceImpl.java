package com.ngtesting.platform.service.impl;

import java.util.LinkedList;
import java.util.List;

import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import com.ngtesting.platform.entity.TestDocument;
import com.ngtesting.platform.entity.TestDocument.DocType;
import com.ngtesting.platform.service.DocumentService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.DocumentVo;
import com.ngtesting.platform.vo.Page;

@Service
public class DocumentServiceImpl extends BaseServiceImpl implements DocumentService {

    @Override
    public List<TestDocument> listByEvent(Long eventId, DocType type) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestDocument.class);
        dc.add(Restrictions.eq("eventId", eventId));

        if (type != null) {
        	dc.add(Restrictions.eq("docType", type));
        }

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("id"));
        List<TestDocument> docPos = findAllByCriteria(dc);

        return docPos;
    }

	@Override
	public Page listByPage(long eventId, int currentPage, int itemsPerPage, DocType type) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestDocument.class);
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
	public TestDocument save(DocumentVo vo) {
		if (vo == null) {
			return null;
		}

		TestDocument po = new TestDocument();
		if (vo.getId() != null) {
			po = (TestDocument) get(TestDocument.class, vo.getId());
		}

		po.setEventId(vo.getEventId());
		po.setTitle(vo.getTitle());
		po.setUri(vo.getUri());

		saveOrUpdate(po);
		return po;
	}

	@Override
	public boolean remove(Long id) {
		TestDocument po = (TestDocument) get(TestDocument.class, id);
		po.setDeleted(true);
		saveOrUpdate(po);

		return true;
	}

	@Override
	public List<DocumentVo> genVos(List<TestDocument> pos) {
        List<DocumentVo> vos = new LinkedList<DocumentVo>();
        for (TestDocument po: pos) {
        	DocumentVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}
	@Override
	public DocumentVo genVo(TestDocument po) {

    	DocumentVo vo = new DocumentVo();
    	BeanUtilEx.copyProperties(vo, po);

		return vo;
	}

}
