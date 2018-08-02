package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.service.IssuePriorityService;
import com.ngtesting.platform.service.IssuePropertyService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_query/")
public class IssueQueryAction extends BaseAction {
	private static final Log log = LogFactory.getLog(CasePriorityAction.class);

	@Autowired
    IssuePriorityService casePriorityService;

	@Autowired
    IssuePropertyService propertyService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

//		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//		Integer orgId = userVo.getDefaultOrgId();
//
//		List<CasePriorityVo> vos = casePriorityService.listVos(orgId);

//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

//		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//
//		Integer id = json.getInteger("id");
//		if (id == null) {
//			ret.put("data", new CasePriorityVo());
//			ret.put("code", Constant.RespCode.SUCCESS.getCode());
//			return ret;
//		}
//
//		TstCasePriority po = (TstCasePriority) casePriorityService.getDetail(TstCasePriority.class, id);
//		CasePriorityVo vo = casePriorityService.genVo(po);
//		ret.put("data", vo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

//		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//		Integer orgId = userVo.getDefaultOrgId();
//
//		CasePriorityVo vo = json.getObject("model", CasePriorityVo.class);
//
//		TstCasePriority po = casePriorityService.save(vo, orgId);
//		CasePriorityVo projectVo = casePriorityService.genVo(po);

//		Map<String,Map<String,String>> casePropertyMap = propertyService.getMap(orgId);
//		ret.put("casePropertyMap", casePropertyMap);

//        ret.put("data", projectVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");

		casePriorityService.delete(id);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "setDefault", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

//		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//		Integer orgId = userVo.getDefaultOrgId();
//		Integer id = json.getInteger("id");
//
//		boolean success = casePriorityService.setDefaultPers(id, orgId);
//		List<CasePriorityVo> vos = casePriorityService.listVos(orgId);
//
//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

//		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//		Integer orgId = userVo.getDefaultOrgId();
//		Integer id = json.getInteger("id");
//		String act = json.getString("act");
//
//		boolean success = casePriorityService.changeOrder(id, act, orgId);
//
//		List<CasePriorityVo> vos = casePriorityService.listVos(orgId);
//
//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}
}
