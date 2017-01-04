package cn.mobiu.events.service;


import java.util.List;
import java.util.Map;

import cn.mobiu.events.entity.EvtOrganizer;
import cn.mobiu.events.vo.OrganizerVo;

public interface OrganizerService extends BaseService {

	List<EvtOrganizer> listByEvent(Long eventId);

	Map<String, List<OrganizerVo>> genOrganizerMap(
			List<EvtOrganizer> organizerPos);

	OrganizerVo genVo(EvtOrganizer po);

}
