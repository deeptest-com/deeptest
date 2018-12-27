package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.action.admin.CaseTypeAdmin;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuTypeSolution;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueTypeSolutionService;
import com.ngtesting.platform.servlet.PrivOrg;
import com.ngtesting.platform.servlet.PrivPrj;
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
@RequestMapping(Constant.API_PATH_CLIENT + "issue_type/")
public class IssueTypeAction extends BaseAction {
	private static final Log log = LogFactory.getLog(CaseTypeAdmin.class);

	@Autowired
	IssueTypeSolutionService solutionService;

	@RequestMapping(value = "getByProject", method = RequestMethod.POST)
	@ResponseBody
	@PrivPrj
	public Map<String, Object> getByProject(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();
		Integer prjId = user.getDefaultPrjId();

		IsuTypeSolution solution = solutionService.getByProject(prjId, orgId);
        List<IsuTypeSolution> solutions = solutionService.list(orgId);

		ret.put("model", solution);
        ret.put("models", solutions);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "setByProject", method = RequestMethod.POST)
	@ResponseBody
	@PrivOrg(perms = {"project-admin"})
	public Map<String, Object> setByProject(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();
		Integer prjId = user.getDefaultPrjId();

        Integer solutionId = json.getInteger("solutionId");

        solutionService.setByProject(solutionId, prjId, orgId);

        IsuTypeSolution solution = solutionService.getByProject(prjId, orgId);
        List<IsuTypeSolution> solutions = solutionService.list(orgId);

        ret.put("model", solution);
        ret.put("models", solutions);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
