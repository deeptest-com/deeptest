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

import com.alibaba.fastjson.JSONObject;

import cn.linkr.events.constants.Constant;
import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.service.ExchangeService;
import cn.linkr.events.service.NewsService;
import cn.linkr.events.service.WifiService;
import cn.linkr.events.util.AuthPassport;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "exchange/")
public class ExchangeAction extends BaseAction {
	@Autowired
	ExchangeService exchangeService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "getData", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getData(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		Map<String, Object> data = new HashMap<String, Object>();
		
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
