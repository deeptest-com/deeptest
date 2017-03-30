package com.ngtesting.platform.action;

import java.util.HashMap;
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
import com.ngtesting.platform.entity.EvtScheduleItem;
import com.ngtesting.platform.service.EventService;
import com.ngtesting.platform.service.ScheduleService;
import com.ngtesting.platform.service.SessionService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.ScheduleItemVo;


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
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request) {
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
		
		
		EvtScheduleItem event = scheduleService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
