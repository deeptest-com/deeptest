package cn.linkr.testspace.service;

import java.util.List;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.vo.GuestVo;
import cn.linkr.testspace.vo.Page;

public interface GuestService extends BaseService {

	List<EvtGuest> list(Long valueOf);

	List<GuestVo> genVos(List<EvtGuest> pos);

	GuestVo genVo(EvtGuest po);

	Page list(Long companyId, int currentPage, int itemsPerPage);

	EvtGuest save(GuestVo vo);

	boolean remove(Long id);

	EvtGuest genPo(GuestVo vo);
	
}
