package cn.linkr.testspace.service;


import java.util.List;

import cn.linkr.testspace.entity.EvtRegisterRecord;

public interface RegisterService extends BaseService {

	void register(String eventId, String sessionIds);

	List<EvtRegisterRecord> listRegisterSession(Long eventId, Long clientId);

	long getRegisterNumb(Long eventId);

}
