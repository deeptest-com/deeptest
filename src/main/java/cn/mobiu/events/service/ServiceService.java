package cn.mobiu.events.service;


import java.util.List;

import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.entity.EvtService;
import cn.mobiu.events.entity.EvtService.ServiceType;
import cn.mobiu.events.vo.ServiceVo;

public interface ServiceService extends BaseService {

	List<EvtService> list(Long eventId, ServiceType type);
	List<EvtService> listForEdit(Long eventId, ServiceType type);

	EvtService save(ServiceVo vo);

	boolean disable(Long id);

}
