package com.ngtesting.platform.service;

import java.util.List;

import com.ngtesting.platform.entity.EvtGuest;
import com.ngtesting.platform.vo.GuestVo;
import com.ngtesting.platform.vo.Page;

public interface GuestService extends BaseService {

	List<EvtGuest> list(Long valueOf);

	List<GuestVo> genVos(List<EvtGuest> pos);

	GuestVo genVo(EvtGuest po);

	Page list(Long orgId, int currentPage, int itemsPerPage);

	EvtGuest save(GuestVo vo);

	boolean remove(Long id);

	EvtGuest genPo(GuestVo vo);
	
}
