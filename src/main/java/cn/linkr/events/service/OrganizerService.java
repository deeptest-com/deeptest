package cn.linkr.events.service;


import java.util.List;
import java.util.Map;

import cn.linkr.events.entity.EvtOrganizer;
import cn.linkr.events.vo.OrganizerVo;

public interface OrganizerService extends BaseService {

	List<EvtOrganizer> listByEvent(Long eventId);

	Map<String, List<OrganizerVo>> genOrganizerMap(
			List<EvtOrganizer> organizerPos);

	OrganizerVo genVo(EvtOrganizer po);

}
