package com.ngtesting.platform.action;

import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.EvtSession;
import com.ngtesting.platform.service.EventService;
import com.ngtesting.platform.service.RegisterService;
import com.ngtesting.platform.service.SessionService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.SessionVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "session/")
public class SessionAction extends BaseAction {
	
	@Autowired
	EventService eventService;
	@Autowired
	SessionService sessionService;
	
	@Autowired
	RegisterService registerService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody SessionVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		EvtSession session = sessionService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		boolean success = sessionService.remove(to.getLong("id"), to.getString("type"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
