package cn.mobiu.events.service;


import java.util.List;

import com.alibaba.fastjson.JSONObject;

import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.entity.EvtEvent;
import cn.mobiu.events.entity.EvtRegisterRecord;
import cn.mobiu.events.entity.EvtSession;
import cn.mobiu.events.vo.SessionVo;

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
