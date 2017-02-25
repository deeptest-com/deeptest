package cn.linkr.testspace.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import cn.linkr.testspace.entity.EvtClient;
import cn.linkr.testspace.entity.EvtEvent;
import cn.linkr.testspace.entity.EvtRegisterRecord;
import cn.linkr.testspace.entity.EvtSession;
import cn.linkr.testspace.vo.SessionVo;

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
