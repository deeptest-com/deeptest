package com.ngtesting.platform.service.intf;


import com.ngtesting.platform.model.Document;
import com.ngtesting.platform.vo.Page;

import java.util.List;

public interface IssueDocumentService extends BaseService {

	Page listByPage(Integer eventId, Integer currentPage, Integer itemsPerPage, String type);

	List<Document> genVos(List<Document> docPos);

	Document genVo(Document po);
	boolean remove(Integer id);
	Document save(Document vo);

}
