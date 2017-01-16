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
import cn.mobiu.events.entity.EvtGuest;
import cn.mobiu.events.entity.EvtService;
import cn.mobiu.events.service.ServiceService;
import cn.mobiu.events.util.AuthPassport;
import cn.mobiu.events.util.BeanUtilEx;
import cn.mobiu.events.vo.GuestVo;
import cn.mobiu.events.vo.Page;
import cn.mobiu.events.vo.ServiceVo;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "service/")
public class ServiceAction extends BaseAction {
	@Autowired
	ServiceService serviceService;
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		List<EvtService> pos = serviceService.list(Long.valueOf(eventId), null);
        List<ServiceVo> vos = new LinkedList<ServiceVo>();
        for (EvtService po: pos) {
        	ServiceVo vo = new ServiceVo();
        	BeanUtilEx.copyProperties(vo, po);
        	vo.setTypeName(po.getType().getName());
        	vos.add(vo);
        }

		ret.put("services", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "listForEdit", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> listForEdit(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String eventId = req.getString("eventId");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		List<EvtService> pos = serviceService.listForEdit(Long.valueOf(eventId), null);
        List<ServiceVo> vos = new LinkedList<ServiceVo>();
        for (EvtService po: pos) {
        	ServiceVo vo = new ServiceVo();
        	BeanUtilEx.copyProperties(vo, po);
        	vo.setTypeName(po.getType().getName());
        	vos.add(vo);
        }

		ret.put("services", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request) {
		Map<String, Object> ret = new HashMap<String, Object>();
		JSONObject req = reqJson(request);
		String serviceId = req.getString("serviceId");
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		
		EvtService po = (EvtService) serviceService.get(EvtService.class, Long.valueOf(serviceId));

    	ServiceVo vo = new ServiceVo();
    	BeanUtilEx.copyProperties(vo, po);
    	vo.setTypeName(po.getType().getName());

		ret.put("service", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody ServiceVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		EvtService service = serviceService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "disable", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> disable(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		EvtClient client = (EvtClient) request.getSession().getAttribute(Constant.HTTP_SESSION_CLIENT_KEY);
		boolean success = serviceService.disable(to.getLong("id"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
