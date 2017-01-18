package cn.linkr.events.service;


import java.util.List;

import cn.linkr.events.entity.EvtBanner;
import cn.linkr.events.entity.EvtDocument;
import cn.linkr.events.entity.EvtDocument.DocType;
import cn.linkr.events.vo.BannerVo;
import cn.linkr.events.vo.DocumentVo;
import cn.linkr.events.vo.Page;
public interface BannerService extends BaseService {

	List<EvtBanner> listByEvent(Long eventId);
	Page listByPage(long eventId, int currentPage, int itemsPerPage);

	List<BannerVo> genVos(List<EvtBanner> docPos);

	BannerVo genVo(EvtBanner po);
	boolean remove(Long id);
	EvtBanner save(BannerVo vo);
	
}
