package cn.linkr.events.action.client;

import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.service.NoticeService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.Constant;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "notice/")
public class NoticeAction extends BaseAction {
	@Autowired
	NoticeService noticeService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getEvent(HttpServletRequest request, @RequestBody Map<String, Object> json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String eventId = json.get("eventId").toString();
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


}
