package cn.mobiu.events.action.client;

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

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.entity.EvtBizcard;
import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.entity.EvtFeedback;
import cn.mobiu.events.service.BizcardService;
import cn.mobiu.events.service.FeedbackService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.BizcardVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "feedback/")
public class FeedbackAction extends BaseAction {
	@Autowired
	FeedbackService feedbackService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		String feedbackType = req.getString("feedbackType");
		String content = req.getString("content");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		EvtFeedback feedback = new EvtFeedback();
		feedback.setAuthorId(client.getId());
		feedback.setEventId(Long.valueOf(eventId));
		feedback.setFeedbackType(feedbackType);
		feedback.setContent(content);
		feedbackService.saveOrUpdate(feedback);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
//	@AuthPassport(validate = true)
//	@RequestMapping(value = "list", method = RequestMethod.POST)
//	@ResponseBody
//	public Map<String, Object> list(HttpServletRequest request) {
//		Map<String, Object> ret = new HashMap<String, Object>();
//		JSONObject req = reqData(request);
//		String eventId = req.getString("eventId");
//		
//		EvtClient client = (EvtClient) request.getAttribute(Constant.REQUEST_USER);
//
//		ret.put("code", Constant.RespCode.SUCCESS.getCode());
//		return ret;
//	}
//	
//	@AuthPassport(validate = true)
//	@RequestMapping(value = "get", method = RequestMethod.POST)
//	@ResponseBody
//	public Map<String, Object> getBizcard(HttpServletRequest request) {
//		Map<String, Object> ret = new HashMap<String, Object>();
//		JSONObject req = reqData(request);
//		String eventId = req.getString("eventId");
//		
//		EvtClient client = (EvtClient) request.getAttribute(Constant.REQUEST_USER);
//
//		ret.put("code", Constant.RespCode.SUCCESS.getCode());
//		return ret;
//	}

}
