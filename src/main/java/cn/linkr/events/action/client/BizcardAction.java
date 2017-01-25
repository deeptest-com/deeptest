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

import cn.linkr.events.entity.EvtBizcard;
import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtDocument;
import cn.linkr.events.entity.EvtDocument.DocType;
import cn.linkr.events.service.BizcardService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.util.Constant;
import cn.linkr.events.vo.BizcardVo;
import cn.linkr.events.vo.DocumentVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "bizcard/")
public class BizcardAction extends BaseAction {
	@Autowired
	BizcardService bizcardService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		List<EvtBizcard> pos = bizcardService.listByEvent(Long.valueOf(eventId), client.getId());
		List<BizcardVo> vos = bizcardService.genVos(pos);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("bizcards", vos);
		
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		String bizcardId = req.getString("bizcardId");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		BizcardVo vo = null;
		if (bizcardId != null) {
			EvtBizcard po = (EvtBizcard) bizcardService.getDetail(Long.valueOf(bizcardId), 
					Long.valueOf(eventId), client.getId());
	        if (po != null) {
	        	vo = new BizcardVo();
		        BeanUtilEx.copyProperties(vo, po);
	        }
		}

		ret.put("bizcard", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		
		return ret;
	}


}
