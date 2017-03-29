package com.ngtesting.platform.service;


import java.util.List;

import com.ngtesting.platform.entity.EvtDocument;
import com.ngtesting.platform.entity.EvtDocument.DocType;
import com.ngtesting.platform.vo.DocumentVo;
import com.ngtesting.platform.vo.Page;
public interface DocumentService extends BaseService {

	List<EvtDocument> listByEvent(Long eventId, DocType type);
	Page listByPage(long eventId, int currentPage, int itemsPerPage, DocType type);

	List<DocumentVo> genVos(List<EvtDocument> docPos);

	DocumentVo genVo(EvtDocument po);
	boolean remove(Long id);
	EvtDocument save(DocumentVo vo);

}
