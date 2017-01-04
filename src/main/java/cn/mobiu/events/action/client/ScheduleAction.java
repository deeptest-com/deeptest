package cn.mobiu.events.action.client;

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
import cn.mobiu.events.entity.EvtScheduleItem;
import cn.mobiu.events.service.EventService;
import cn.mobiu.events.service.ScheduleService;
import cn.mobiu.events.service.SessionService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.util.DateUtils;
import cn.mobiu.events.vo.EventVo;
import cn.mobiu.events.vo.ScheduleItemVo;
import cn.mobiu.events.vo.SessionVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "schedule/")
public class ScheduleAction extends BaseAction {
	
	@Autowired
	EventService eventService;
	@Autowired
	SessionService sessionService;
	
	@Autowired
	ScheduleService scheduleService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "listByEvent", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> listByEvent(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		String isNest = req.getString("isNest") != null? req.getString("isNest"): "false";
		
		List<EvtScheduleItem> scheduleItemsBySession = scheduleService.listScheduleItemsBySession(Long.valueOf(eventId));
		List<EvtScheduleItem> scheduleItemsByDate = scheduleService.listScheduleItemsByDate(Long.valueOf(eventId));
		
        List<ScheduleItemVo> vosBySession = scheduleService.genVosBySession(scheduleItemsBySession, Boolean.valueOf(isNest));
        List<ScheduleItemVo> vosByDate = scheduleService.genVosByDate(scheduleItemsByDate);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		
		ret.put("bySession", vosBySession);
		ret.put("byDate", vosByDate);
		
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody ScheduleItemVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		EvtScheduleItem event = scheduleService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
