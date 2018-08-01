package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.Document;
import com.ngtesting.platform.service.DocumentService;
import com.ngtesting.platform.vo.Page;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "issue_document/")
public class IssueDocumentAction extends BaseAction {
	@Autowired
    DocumentService documentService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer eventId = json.getInteger("eventId");
		int page = json.getInteger("listByPage") == null? 0: json.getInteger("listByPage") - 1;
		int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

		Page pageData = documentService.listByPage(eventId, page, pageSize, null);
//		List<DocumentVo> vos = documentService.genVos(pageData.getItems());

		ret.put("collectionSize", pageData.getTotal());
//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody Document vo) {
		Map<String, Object> ret = new HashMap<String, Object>();

//		TestDocument doc = documentService.save(vo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "remove", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject to) {
		Map<String, Object> ret = new HashMap<String, Object>();

//		boolean success = documentService.remove(to.getInteger("id"));

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}
}
