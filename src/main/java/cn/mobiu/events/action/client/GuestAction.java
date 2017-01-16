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
import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.entity.EvtScheduleItem;
import cn.mobiu.events.entity.EvtSession;
import cn.mobiu.events.service.GuestService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.EventVo;
import cn.mobiu.events.vo.GuestVo;
import cn.mobiu.events.vo.Page;
import cn.mobiu.events.vo.ScheduleItemVo;
import cn.mobiu.events.vo.SessionVo;


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
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "listByPage", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> listByPage(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		long eventId = json.getLong("eventId");
		int currentPage = json.getInteger("currentPage") == null? 0: json.getInteger("currentPage") - 1;
		int itemsPerPage = json.getInteger("itemsPerPage") == null? Constant.PAGE_SIZE: json.getInteger("itemsPerPage");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		Page page = guestService.list(eventId, currentPage, itemsPerPage);
		List<GuestVo> vos = guestService.genVos(page.getItems());
        
		ret.put("totalItems", page.getTotal());
        ret.put("guests", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String guestId = req.getString("guestId");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		EvtGuest po = (EvtGuest) guestService.get(EvtGuest.class, Long.valueOf(guestId));
		GuestVo vo = guestService.genVo(po);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("guest", vo);
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody GuestVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		EvtGuest guest = guestService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		boolean success = guestService.remove(to.getLong("id"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
