package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuPriority;
import com.ngtesting.platform.model.IsuPrioritySolution;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IssuePriorityService;
import com.ngtesting.platform.service.IssuePrioritySolutionService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_priority_solution/")
public class IssuePrioritySolutionAdmin extends BaseAction {
	private static final Log log = LogFactory.getLog(CasePriorityAdmin.class);

	@Autowired
	IssuePrioritySolutionService solutionService;

	@Autowired
	IssuePriorityService priorityService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		List<IsuPrioritySolution> vos = solutionService.list(orgId);

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
		IsuPrioritySolution po;
		if (id == null) {
			po = new IsuPrioritySolution();
		} else {
			po = solutionService.getDetail(id, orgId);
		}

		if (po == null) { // 当对象不是默认org的，此处为空
			return authFail();
		}
        List<IsuPriority> otherItems = priorityService.listNotInSolution(id, orgId);

        ret.put("data", po);
        ret.put("otherItems", otherItems);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		IsuPrioritySolution vo = json.getObject("model", IsuPrioritySolution.class);

		IsuPrioritySolution po = solutionService.save(vo, orgId);
		if (po == null) {    // 当对象不是默认org的，update的结果会返回空
			return authFail();
		}

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

		Boolean result = solutionService.delete(id, orgId);
		if (!result) { // 当对象不是默认org的，结果会返回false
			return authFail();
		}

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "addPriority", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> addPriority(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer priorityId = json.getInteger("priorityId");
		Integer solutionId = json.getInteger("solutionId");

		solutionService.addPriority(priorityId, solutionId, orgId);

		IsuPrioritySolution po = solutionService.getDetail(solutionId, orgId);
		List<IsuPriority> otherItems = priorityService.listNotInSolution(solutionId, orgId);

		ret.put("data", po);
		ret.put("otherItems", otherItems);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "removePriority", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> removePriority(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer priorityId = json.getInteger("priorityId");
		Integer solutionId = json.getInteger("solutionId");

		solutionService.removePriority(priorityId, solutionId, orgId);

        IsuPrioritySolution po = solutionService.getDetail(solutionId, orgId);
        List<IsuPriority> otherItems = priorityService.listNotInSolution(solutionId, orgId);

		ret.put("data", po);
		ret.put("otherItems", otherItems);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "addAll", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> addAll(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer solutionId = json.getInteger("solutionId");

		solutionService.addAll(solutionId, orgId);

		IsuPrioritySolution po = solutionService.getDetail(solutionId, orgId);
		List<IsuPriority> otherItems = priorityService.listNotInSolution(solutionId, orgId);

		ret.put("data", po);
		ret.put("otherItems", otherItems);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "removeAll", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> removeAll(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		Integer solutionId = json.getInteger("solutionId");

		solutionService.removeAll(solutionId, orgId);

		IsuPrioritySolution po = solutionService.getDetail(solutionId, orgId);
		List<IsuPriority> otherItems = priorityService.listNotInSolution(solutionId, orgId);

		ret.put("data", po);
		ret.put("otherItems", otherItems);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
