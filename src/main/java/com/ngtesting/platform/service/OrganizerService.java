package com.ngtesting.platform.service;


import java.util.List;
import java.util.Map;

import com.ngtesting.platform.entity.EvtOrganizer;
import com.ngtesting.platform.vo.OrganizerVo;

public interface OrganizerService extends BaseService {

	List<EvtOrganizer> listByEvent(Long eventId);

	Map<String, List<OrganizerVo>> genOrganizerMap(
			List<EvtOrganizer> organizerPos);

	OrganizerVo genVo(EvtOrganizer po);

}
