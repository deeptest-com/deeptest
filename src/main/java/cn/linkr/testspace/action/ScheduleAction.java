package cn.linkr.testspace.action;

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


import cn.linkr.testspace.entity.EvtScheduleItem;
import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.service.EventService;
import cn.linkr.testspace.service.ScheduleService;
import cn.linkr.testspace.service.SessionService;
import cn.linkr.testspace.util.AuthPassport;
import cn.linkr.testspace.util.Constant;
import cn.linkr.testspace.vo.ScheduleItemVo;

import com.alibaba.fastjson.JSONObject;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "schedule/")
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
