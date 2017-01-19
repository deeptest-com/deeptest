package cn.linkr.events.action.admin;

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

import cn.linkr.events.action.client.BaseAction;
import cn.linkr.events.constants.Constant;
import cn.linkr.events.entity.EvtBanner;
import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtDocument;
import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.EvtDocument.DocType;
import cn.linkr.events.entity.EvtEvent;
import cn.linkr.events.entity.EvtOrganizer;
import cn.linkr.events.service.BannerService;
import cn.linkr.events.service.DocumentService;
import cn.linkr.events.service.EventService;
import cn.linkr.events.service.OrganizerService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.vo.BannerVo;
import cn.linkr.events.vo.DocumentVo;
import cn.linkr.events.vo.EventVo;
import cn.linkr.events.vo.OrganizerVo;
import cn.linkr.events.vo.Page;
import cn.linkr.events.vo.SessionVo;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "event/")
public class EventAdmin extends BaseAction {
	@Autowired
	EventService eventService;
	@Autowired
	DocumentService documentService;
	@Autowired
	BannerService bannerService;
	@Autowired
	OrganizerService organizerService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject req) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		SysUser user = (SysUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		int currentPage = req.getString("currentPage") == null? 0: Integer.valueOf(req.getString("currentPage")) - 1;
		int itemsPerPage = req.getString("itemsPerPage") == null? Constant.PAGE_SIZE: Integer.valueOf(req.getString("itemsPerPage"));
		String status = req.getString("status");
		
		Page page = eventService.list(user, status, currentPage, itemsPerPage);
		List<EventVo> vos = eventService.genVos(page.getItems());
        
		ret.put("totalItems", page.getTotal());
        ret.put("events", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody EventVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		SysUser user = (SysUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		EvtEvent event = eventService.save(vo, user);
        EventVo eventVo = eventService.genVo(event);
        
        ret.put("event", eventVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
}
