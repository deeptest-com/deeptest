package com.ngtesting.platform.service;


import com.ngtesting.platform.entity.TestDocument;
import com.ngtesting.platform.entity.TestDocument.DocType;
import com.ngtesting.platform.vo.DocumentVo;
import com.ngtesting.platform.vo.Page;

import java.util.List;
public interface DocumentService extends BaseService {

	List<TestDocument> listByEvent(Long eventId, DocType type);
	Page listByPage(long eventId, int currentPage, int itemsPerPage, DocType type);

	List<DocumentVo> genVos(List<TestDocument> docPos);

	DocumentVo genVo(TestDocument po);
	boolean remove(Long id);
	TestDocument save(DocumentVo vo);

}
