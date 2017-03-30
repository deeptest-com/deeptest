package com.ngtesting.platform.service;


import java.util.List;

import com.ngtesting.platform.entity.EvtBanner;
import com.ngtesting.platform.vo.BannerVo;
import com.ngtesting.platform.vo.Page;
public interface BannerService extends BaseService {

	List<EvtBanner> listByEvent(Long eventId);
	Page listByPage(long eventId, int currentPage, int itemsPerPage);

	List<BannerVo> genVos(List<EvtBanner> docPos);

	BannerVo genVo(EvtBanner po);
	boolean remove(Long id);
	EvtBanner save(BannerVo vo);
	
}
