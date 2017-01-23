package cn.linkr.events.service;


import java.util.List;

import cn.linkr.events.entity.EvtGuest;
import cn.linkr.events.entity.EvtService;
import cn.linkr.events.entity.EvtService.ServiceType;
import cn.linkr.events.vo.ServiceVo;

public interface ServiceService extends BaseService {

	List<EvtService> list(Long eventId, ServiceType type);
	List<EvtService> listForEdit(Long eventId, ServiceType type);

	EvtService save(ServiceVo vo);

	boolean disablePers(Long id);

}
