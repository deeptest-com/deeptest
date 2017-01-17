package cn.mobiu.events.action.client;

import java.io.BufferedReader;
import java.io.IOException;
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
import cn.mobiu.events.entity.EvtBanner;
import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.entity.EvtDocument;
import cn.mobiu.events.entity.EvtEvent;
import cn.mobiu.events.entity.EvtOrganizer;
import cn.mobiu.events.entity.EvtDocument.DocType;
import cn.mobiu.events.service.BannerService;
import cn.mobiu.events.service.DocumentService;
import cn.mobiu.events.service.EventService;
import cn.mobiu.events.service.OrganizerService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.BannerVo;
import cn.mobiu.events.vo.DocumentVo;
import cn.mobiu.events.vo.EventVo;
import cn.mobiu.events.vo.OrganizerVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "home/")
public class HomeAction extends BaseAction {
	@Autowired
	EventService eventService;
	@Autowired
	DocumentService documentService;
	@Autowired
	BannerService bannerService;
	@Autowired
	OrganizerService organizerService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "index", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> index(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		request.getSession(true).setAttribute("TEST", "123456");

		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		EvtEvent event = eventService.getDetail(Long.valueOf(eventId));
        EventVo eventVo = eventService.genVo(event);
		
		List<EvtDocument> docPos = documentService.listByEvent(Long.valueOf(eventId), DocType.banner);
        List<DocumentVo> docVos = documentService.genVos(docPos);
        
		List<EvtBanner> bannerPos = bannerService.listByEvent(Long.valueOf(eventId));
		List<BannerVo> bannerVos = bannerService.genVos(bannerPos);
        
        List<EvtOrganizer> organizerPos = organizerService.listByEvent(Long.valueOf(eventId));
        Map<String, List<OrganizerVo>> organizerMap = organizerService.genOrganizerMap(organizerPos);
        
        eventVo.setDocuments(docVos);
        eventVo.setBanners(bannerVos);
        eventVo.setOrganizers(organizerMap);
	
		ret.put("event", eventVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
