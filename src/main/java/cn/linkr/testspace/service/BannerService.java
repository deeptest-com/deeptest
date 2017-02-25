package cn.linkr.testspace.service;


import java.util.List;

import cn.linkr.testspace.entity.EvtBanner;
import cn.linkr.testspace.entity.EvtDocument;
import cn.linkr.testspace.entity.EvtDocument.DocType;
import cn.linkr.testspace.vo.BannerVo;
import cn.linkr.testspace.vo.DocumentVo;
import cn.linkr.testspace.vo.Page;
public interface BannerService extends BaseService {

	List<EvtBanner> listByEvent(Long eventId);
	Page listByPage(long eventId, int currentPage, int itemsPerPage);

	List<BannerVo> genVos(List<EvtBanner> docPos);

	BannerVo genVo(EvtBanner po);
	boolean remove(Long id);
	EvtBanner save(BannerVo vo);
	
}
