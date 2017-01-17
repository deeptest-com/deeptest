package cn.mobiu.events.service;


import java.util.List;

import cn.mobiu.events.entity.EvtBanner;
import cn.mobiu.events.entity.EvtDocument;
import cn.mobiu.events.entity.EvtDocument.DocType;
import cn.mobiu.events.vo.BannerVo;
import cn.mobiu.events.vo.DocumentVo;
import cn.mobiu.events.vo.Page;
public interface BannerService extends BaseService {

	List<EvtBanner> listByEvent(Long eventId);
	Page listByPage(long eventId, int currentPage, int itemsPerPage);

	List<BannerVo> genVos(List<EvtBanner> docPos);

	BannerVo genVo(EvtBanner po);
	
}
