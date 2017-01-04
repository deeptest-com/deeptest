package cn.mobiu.events.service;


import java.util.List;

import cn.mobiu.events.entity.EvtRegisterRecord;

public interface RegisterService extends BaseService {

	void register(String eventId, String sessionIds);

	List<EvtRegisterRecord> listRegisterSession(Long eventId, Long clientId);

	long getRegisterNumb(Long eventId);

}
