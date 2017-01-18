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

import cn.linkr.events.constants.Constant;
import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtGuest;
import cn.linkr.events.entity.EvtQa;
import cn.linkr.events.service.QaService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.vo.GuestVo;
import cn.linkr.events.vo.QaVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "qa/")
public class QaAction extends BaseAction {
	@Autowired
	QaService qaService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");

		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		List<EvtQa> pos = qaService.list(Long.valueOf(eventId), client.getId());
		 List<QaVo> vos = qaService.genVos(pos);
		
		ret.put("qas", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		String content = req.getString("content");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		qaService.save(Long.valueOf(eventId), content);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
