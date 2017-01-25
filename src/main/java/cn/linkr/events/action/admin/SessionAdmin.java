package cn.linkr.events.action.admin;

import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import cn.linkr.events.action.client.BaseAction;
import cn.linkr.events.entity.EvtSession;
import cn.linkr.events.entity.SysUser;
import cn.linkr.events.service.EventService;
import cn.linkr.events.service.RegisterService;
import cn.linkr.events.service.SessionService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.Constant;
import cn.linkr.events.vo.SessionVo;

import com.alibaba.fastjson.JSONObject;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "session/")
public class SessionAdmin extends BaseAction {
	
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
