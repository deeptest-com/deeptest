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

import cn.mobiu.events.constants.Constant;
import cn.mobiu.events.entity.EvtAround;
import cn.mobiu.events.entity.EvtClient;
import cn.mobiu.events.service.AroundService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.AroundVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "around/")
public class AroundAction extends BaseAction {
	@Autowired
	AroundService arroundService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody Map<String, Object> json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String eventId = json.get("eventId").toString();
		
		List<EvtAround> pos = arroundService.list(Long.valueOf(eventId), null);
		List<AroundVo> vos = arroundService.genVos(pos);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


}
