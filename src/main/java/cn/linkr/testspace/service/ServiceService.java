package cn.linkr.testspace.service;


import java.util.List;

import cn.linkr.testspace.entity.EvtGuest;
import cn.linkr.testspace.entity.EvtService;
import cn.linkr.testspace.entity.EvtService.ServiceType;
import cn.linkr.testspace.vo.ServiceVo;

public interface ServiceService extends BaseService {

	List<EvtService> list(Long eventId, ServiceType type);
	List<EvtService> listForEdit(Long eventId, ServiceType type);

	EvtService save(ServiceVo vo);

	boolean disablePers(Long id);

}
