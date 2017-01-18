package cn.linkr.events.service;


import java.util.List;

import cn.linkr.events.entity.EvtRegisterRecord;

public interface RegisterService extends BaseService {

	void register(String eventId, String sessionIds);

	List<EvtRegisterRecord> listRegisterSession(Long eventId, Long clientId);

	long getRegisterNumb(Long eventId);

}
