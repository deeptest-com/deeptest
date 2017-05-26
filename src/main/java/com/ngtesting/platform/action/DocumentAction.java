package com.ngtesting.platform.action;

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
import com.ngtesting.platform.entity.EvtDocument;
import com.ngtesting.platform.service.DocumentService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.vo.DocumentVo;
import com.ngtesting.platform.vo.Page;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "document/")
public class DocumentAction extends BaseAction {
	@Autowired
	DocumentService documentService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		long eventId = json.getLong("eventId");
		int page = json.getInteger("page") == null? 0: json.getInteger("page") - 1;
		int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");
		
		Page pageData = documentService.listByPage(eventId, page, pageSize, null);
		List<DocumentVo> vos = documentService.genVos(pageData.getItems());
        
		ret.put("collectionSize", pageData.getTotal());
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody DocumentVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		EvtDocument doc = documentService.save(vo);
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@AuthPassport(validate = true)
	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		boolean success = documentService.remove(to.getLong("id"));
		
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
