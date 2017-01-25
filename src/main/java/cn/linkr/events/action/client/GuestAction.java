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
import cn.linkr.events.entity.EvtGuest;
import cn.linkr.events.entity.EvtScheduleItem;
import cn.linkr.events.entity.EvtSession;
import cn.linkr.events.service.GuestService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.util.Constant;
import cn.linkr.events.vo.EventVo;
import cn.linkr.events.vo.GuestVo;
import cn.linkr.events.vo.Page;
import cn.linkr.events.vo.ScheduleItemVo;
import cn.linkr.events.vo.SessionVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "guest/")
public class GuestAction extends BaseAction {
	@Autowired
	GuestService guestService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		List<EvtGuest> pos = guestService.list(Long.valueOf(eventId));
		List<GuestVo> vos = guestService.genVos(pos);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("guests", vos);
		return ret;
	}
	
}
