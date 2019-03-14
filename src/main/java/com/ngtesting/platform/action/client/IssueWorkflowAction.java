package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuWorkflowSolution;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueWorkflowSolutionService;
import com.ngtesting.platform.servlet.PrivPrj;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.apache.shiro.SecurityUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@RestController
@RequestMapping(Constant.API_PATH_CLIENT + "issue_workflow/")
public class IssueWorkflowAction extends BaseAction {
	private static final Log log = LogFactory.getLog(IssueWorkflowAction.class);

	@Autowired
	IssueWorkflowSolutionService solutionService;

	@RequestMapping(value = "getByProject", method = RequestMethod.POST)
	@PrivPrj(perms = {"project:maintain"})
	public Map<String, Object> getByProject(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();
		Integer prjId = user.getDefaultPrjId();

		IsuWorkflowSolution solution = solutionService.getByProject(prjId, orgId);
		List<IsuWorkflowSolution> solutions = solutionService.list(orgId);

		Map itemMap = solutionService.getItemsMap(solution.getId(), orgId);

		ret.put("model", solution);
		ret.put("models", solutions);
		ret.put("itemMap", itemMap);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "setByProject", method = RequestMethod.POST)
	@PrivPrj(perms = {"project:maintain"})
	public Map<String, Object> setByProject(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) SecurityUtils.getSubject().getPrincipal();
		Integer orgId = user.getDefaultOrgId();
		Integer prjId = user.getDefaultPrjId();

		Integer solutionId = json.getInteger("solutionId");

		solutionService.setByProject(solutionId, prjId, orgId);

		IsuWorkflowSolution solution = solutionService.getByProject(prjId, orgId);
		List<IsuWorkflowSolution> solutions = solutionService.list(orgId);

		ret.put("model", solution);
		ret.put("models", solutions);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
