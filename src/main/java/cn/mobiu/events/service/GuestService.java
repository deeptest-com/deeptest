package cn.mobiu.events.service;

import java.util.List;

import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.vo.GuestVo;
import cn.mobiu.events.vo.Page;

public interface GuestService extends BaseService {

	List<EvtGuest> list(Long valueOf);

	List<GuestVo> genVos(List<EvtGuest> pos);

	GuestVo genVo(EvtGuest po);

	Page list(Long companyId, int currentPage, int itemsPerPage);

	EvtGuest save(GuestVo vo);

	boolean remove(Long id);

	EvtGuest genPo(GuestVo vo);
	
}
