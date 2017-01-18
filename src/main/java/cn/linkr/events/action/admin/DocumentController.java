package cn.linkr.events.action.admin;

import java.util.HashMap;
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
import cn.linkr.events.entity.EvtBizcard;
import cn.linkr.events.entity.EvtClient;
import cn.linkr.events.entity.EvtDocument;
import cn.linkr.events.entity.EvtGuest;
import cn.linkr.events.entity.SysUser;
import cn.linkr.events.service.BizcardService;
import cn.linkr.events.service.DocumentService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.vo.BizcardVo;
import cn.linkr.events.vo.DocumentVo;
import cn.linkr.events.vo.GuestVo;
import cn.linkr.events.vo.Page;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "document/")
public class DocumentController extends BaseAction {
	@Autowired
	DocumentService documentService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		SysUser user = (SysUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		long eventId = json.getLong("eventId");
		int currentPage = json.getInteger("currentPage") == null? 0: json.getInteger("currentPage") - 1;
		int itemsPerPage = json.getInteger("itemsPerPage") == null? Constant.PAGE_SIZE: json.getInteger("itemsPerPage");
		
		Page page = documentService.listByPage(eventId, currentPage, itemsPerPage, null);
		List<DocumentVo> vos = documentService.genVos(page.getItems());
        
		ret.put("totalItems", page.getTotal());
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody DocumentVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		SysUser user = (SysUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		EvtDocument doc = documentService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		SysUser user = (SysUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		boolean success = documentService.remove(to.getLong("id"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
