package cn.mobiu.events.service;


import java.util.List;

import cn.mobiu.events.entity.EvtService;
import cn.mobiu.events.entity.EvtService.ServiceType;

public interface ServiceService extends BaseService {

	List<EvtService> list(Long eventId, ServiceType type);

}
