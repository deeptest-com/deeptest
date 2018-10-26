package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuPriority;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IssuePriorityService;
import com.ngtesting.platform.service.IssuePropertyService;
import com.ngtesting.platform.service.IssueTypeService;
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
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "issue_priority/")
public class IssuePriorityAdmin extends BaseAction {
	private static final Log log = LogFactory.getLog(CasePriorityAdmin.class);

	@Autowired
	IssuePriorityService issuePriorityService;

	@Autowired
	IssuePropertyService issuePropertyService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		List<IsuPriority> vos = issuePriorityService.list(orgId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}


	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer id = json.getInteger("id");
		IsuPriority po;
		if (id == null) {
			po = new IsuPriority();
		} else {
			po = issuePriorityService.get(id, orgId);
		}

		if (po == null) { // 当对象不是默认org的，此处为空
			return authFail();
		}

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		IsuPriority vo = json.getObject("model", IsuPriority.class);

		IsuPriority po = issuePriorityService.save(vo, orgId);
		if (po == null) {    // 当对象不是默认org的，update的结果会返回空
			return authFail();
		}

		Map<String, Map<String, String>> casePropertyMap = issuePropertyService.getMap(orgId);
		ret.put("casePropertyMap", casePropertyMap);

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer id = json.getInteger("id");

		Boolean result = issuePriorityService.delete(id, orgId);
		if (!result) { // 当对象不是默认org的，结果会返回false
			return authFail();
		}

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "setDefault", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer id = json.getInteger("id");

		Boolean result = issuePriorityService.setDefault(id, orgId);
		if (!result) { // 当对象不是默认org的，结果会返回false
			return authFail();
		}

		List<IsuPriority> vos = issuePriorityService.list(orgId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer id = json.getInteger("id");
		String act = json.getString("act");

		Boolean result = issuePriorityService.changeOrder(id, act, orgId);
		if (!result) { // 当对象不是默认org的，结果会返回false
			return authFail();
		}

		List<IsuPriority> vos = issuePriorityService.list(orgId);

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}
}
