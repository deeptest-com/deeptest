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

import cn.linkr.events.entity.EvtBanner;
import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtDocument;
import cn.linkr.events.entity.EvtDocument.DocType;
import cn.linkr.events.entity.EvtEvent;
import cn.linkr.events.entity.EvtOrganizer;
import cn.linkr.events.service.BannerService;
import cn.linkr.events.service.DocumentService;
import cn.linkr.events.service.EventService;
import cn.linkr.events.service.OrganizerService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.util.Constant;
import cn.linkr.events.vo.BannerVo;
import cn.linkr.events.vo.DocumentVo;
import cn.linkr.events.vo.EventVo;
import cn.linkr.events.vo.OrganizerVo;
import cn.linkr.events.vo.Page;
import cn.linkr.events.vo.SessionVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "event/")
public class EventAction extends BaseAction {
	@Autowired
	EventService eventService;
	@Autowired
	DocumentService documentService;
	@Autowired
	BannerService bannerService;
	@Autowired
	OrganizerService organizerService;
	
	
}
