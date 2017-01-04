package cn.mobiu.events.service;


import java.util.List;

import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.vo.GuestVo;

public interface GuestService extends BaseService {

	List<EvtGuest> list(Long valueOf);

	List<GuestVo> genVos(List<EvtGuest> pos);

	GuestVo genVo(EvtGuest po);

}
