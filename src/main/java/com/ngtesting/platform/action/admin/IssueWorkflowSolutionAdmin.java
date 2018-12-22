package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuWorkflow;
import com.ngtesting.platform.model.IsuWorkflowSolution;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueWorkflowService;
import com.ngtesting.platform.service.intf.IssueWorkflowSolutionService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_workflow_solution/")
public class IssueWorkflowSolutionAdmin extends BaseAction {
    @Autowired
    IssueWorkflowService workflowService;
	@Autowired
    IssueWorkflowSolutionService workflowSolutionService;

	@RequestMapping(value = "load", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> load(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		List<IsuWorkflowSolution> vos = workflowSolutionService.list(orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("solutions", vos);
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer solutionId = json.getInteger("id");
		IsuWorkflowSolution solution = null;
		if (solutionId == null) {
			solution = new IsuWorkflowSolution();
		} else {
			solution = workflowSolutionService.get(solutionId, orgId);
		}

		ret.put("solution", solution);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "getConfig", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getConfig(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer solutionId = json.getInteger("id");

        IsuWorkflowSolution solution = workflowSolutionService.get(solutionId, orgId);
        Map itemMap = workflowSolutionService.getItemsMap(solutionId, orgId);

        List<IsuWorkflow> workflows = workflowService.list(orgId);

		ret.put("solution", solution);
        ret.put("itemMap", itemMap);
        ret.put("workflows", workflows);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		IsuWorkflowSolution vo = JSON.parseObject(JSON.toJSONString(json), IsuWorkflowSolution.class);
		workflowSolutionService.save(vo, orgId);

		ret.put("solution", vo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer id = json.getInteger("id");

		boolean success = workflowSolutionService.delete(id, orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "changeItem", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeItem(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer solutionId = json.getInteger("solutionId");
		Integer typeId = json.getInteger("type");
		Integer workflowId = json.getInteger("workflow");

		boolean success = workflowSolutionService.changeItem(typeId, workflowId, solutionId, orgId);

        IsuWorkflowSolution solution = workflowSolutionService.get(solutionId, orgId);
        Map itemMap = workflowSolutionService.getItemsMap(solutionId, orgId);

        List<IsuWorkflow> workflows = workflowService.list(orgId);

        ret.put("solution", solution);
        ret.put("itemMap", itemMap);
        ret.put("workflows", workflows);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "setDefault", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        Integer id = json.getInteger("id");

        Boolean result = workflowSolutionService.setDefault(id, orgId);
        if (!result) { // 当对象不是默认org的，结果会返回false
            return authFail();
        }

        List<IsuWorkflowSolution> vos = workflowSolutionService.list(orgId);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("solutions", vos);

        return ret;
    }

}
