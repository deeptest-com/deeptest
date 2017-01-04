package cn.mobiu.events.action.client;

import java.util.Date;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONObject;

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.entity.EvtEvent;
import cn.mobiu.events.entity.EvtRegisterRecord;
import cn.mobiu.events.entity.EvtSession;
import cn.mobiu.events.service.EventService;
import cn.mobiu.events.service.RegisterService;
import cn.mobiu.events.service.SessionService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.util.DateUtils;
import cn.mobiu.events.vo.SessionVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "register/")
public class RegisterAction extends BaseAction {
	@Autowired
	EventService eventService;
	
	@Autowired
	RegisterService registerService;
	
	@Autowired
	SessionService sessionService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "getInfo", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getInfo(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		Map<String, Object> data = new HashMap<String, Object>();
		ret.put("data", data);

		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		
		EvtEvent event = (EvtEvent) eventService.get(EvtEvent.class, Long.valueOf(eventId));
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		List<EvtSession> allSessions = sessionService.listSessionForRegister(Long.valueOf(eventId));
		List<EvtRegisterRecord> registerSessions = registerService.listRegisterSession(Long.valueOf(eventId), client.getId());
		
		List<SessionVo> sessionVos = sessionService.genVos(allSessions, registerSessions);
		boolean alreadyRegister = sessionService.isRegister(allSessions, registerSessions);
        
        long registerNumb = registerService.getRegisterNumb(Long.valueOf(eventId));
        String registerCountdown = DateUtils.DiffStr(new Date(), event.getRegisterEndDatetime());
        
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        data.put("alreadyRegister", alreadyRegister);
        data.put("registerNumb", registerNumb);
        data.put("registerCountdown", registerCountdown);
		
		data.put("sessions", sessionVos);

		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "register", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> register(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		String sessionIds = req.getString("sessionIds");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		registerService.register(eventId, sessionIds);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
