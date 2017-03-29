package com.ngtesting.platform.service;


import java.util.List;

import com.ngtesting.platform.entity.EvtRegisterRecord;

public interface RegisterService extends BaseService {

	void register(String eventId, String sessionIds);

	List<EvtRegisterRecord> listRegisterSession(Long eventId, Long clientId);

	long getRegisterNumb(Long eventId);

}
