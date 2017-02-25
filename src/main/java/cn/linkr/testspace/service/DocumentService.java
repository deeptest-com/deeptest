package cn.linkr.testspace.service;


import java.util.List;

import cn.linkr.testspace.entity.EvtDocument;
import cn.linkr.testspace.entity.EvtDocument.DocType;
import cn.linkr.testspace.vo.DocumentVo;
import cn.linkr.testspace.vo.Page;
public interface DocumentService extends BaseService {

	List<EvtDocument> listByEvent(Long eventId, DocType type);
	Page listByPage(long eventId, int currentPage, int itemsPerPage, DocType type);

	List<DocumentVo> genVos(List<EvtDocument> docPos);

	DocumentVo genVo(EvtDocument po);
	boolean remove(Long id);
	EvtDocument save(DocumentVo vo);

}
