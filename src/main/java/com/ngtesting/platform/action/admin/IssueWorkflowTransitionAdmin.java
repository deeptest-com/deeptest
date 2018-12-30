package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuWorkflowTransition;
import com.ngtesting.platform.model.TstProjectRole;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueWorkflowService;
import com.ngtesting.platform.service.intf.IssueWorkflowTransitionService;
import com.ngtesting.platform.service.intf.ProjectRoleService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_workflow_transition/")
public class IssueWorkflowTransitionAdmin extends BaseAction {
	private static final Log log = LogFactory.getLog(IssueWorkflowTransitionAdmin.class);

	@Autowired
    IssueWorkflowService issueWorkflowService;
    @Autowired
    IssueWorkflowTransitionService transitionService;

    @Autowired
    ProjectRoleService projectRoleService;

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        Integer id = json.getInteger("id");
        Integer workflowId = json.getInteger("workflowId");
        Integer srcStatusId = json.getInteger("srcStatusId");
        Integer dictStatusId = json.getInteger("dictStatusId");

        IsuWorkflowTransition tran;
        if (id != null) {
            tran = transitionService.get(id, orgId);
        } else {
            tran = transitionService.emptyObject(workflowId, srcStatusId, dictStatusId, orgId);
        }

        List<TstProjectRole> projectRoles = transitionService.listProjectRoles(id, orgId);

        ret.put("data", tran);
        ret.put("projectRoles", projectRoles);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

        IsuWorkflowTransition tran = json.getObject("model", IsuWorkflowTransition.class);

        List<Integer> projectRoleIds = json.getObject("projectRoleIds", List.class);

        IsuWorkflowTransition po = transitionService.save(tran, projectRoleIds, orgId);

        ret.put("data", po);
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

        transitionService.delete(id, orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
