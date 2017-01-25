package cn.linkr.events.action.client;

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

import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtEvent;
import cn.linkr.events.entity.EvtScheduleItem;
import cn.linkr.events.service.EventService;
import cn.linkr.events.service.ScheduleService;
import cn.linkr.events.service.SessionService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.util.Constant;
import cn.linkr.events.util.DateUtils;
import cn.linkr.events.vo.EventVo;
import cn.linkr.events.vo.ScheduleItemVo;
import cn.linkr.events.vo.SessionVo;


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

}
