package com.ngtesting.platform.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.EvtRegisterRecord;
import com.ngtesting.platform.entity.EvtSession;
import com.ngtesting.platform.vo.SessionVo;

public interface SessionService extends BaseService {

	List<EvtSession> listSessionsByEvent(Long eventId);

	List<EvtSession> listSessionForRegister(Long eventId);

	List<SessionVo> genVos(List<EvtSession> allSessions,
			List<EvtRegisterRecord> registerSessions);

	boolean isRegister(List<EvtSession> allSessions,
			List<EvtRegisterRecord> registerSessions);

	SessionVo genVo(EvtSession po);

	EvtSession save(SessionVo vo);

	EvtSession genPo(JSONObject vo);

	boolean remove(Long id, String type);

}
