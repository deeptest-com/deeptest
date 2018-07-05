package com.ngtesting.platform.service;

import com.ngtesting.platform.model.Document;
import com.ngtesting.platform.vo.Page;

import java.util.List;

public interface DocumentService extends BaseService {

	Page listByPage(Integer eventId, Integer currentPage, Integer itemsPerPage, Document.DocType type);

	List<Document> genVos(List<Document> docPos);

	Document genVo(Document po);
	boolean remove(Long id);
	Document save(Document vo);

}
