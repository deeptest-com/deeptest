package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.IssuePageService;
import com.ngtesting.platform.service.intf.IssueStatusService;
import com.ngtesting.platform.service.intf.IssueWorkflowService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_workflow/")
public class IssueWorkflowAdmin extends BaseAction {
	private static final Log log = LogFactory.getLog(CaseTypeAdmin.class);

	@Autowired
    IssueWorkflowService issueWorkflowService;
	@Autowired
    IssueStatusService statusService;

    @Autowired
    IssuePageService pageService;
    @Autowired
    ProjectRoleService projectRoleService;

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		List<IsuWorkflow> vos = issueWorkflowService.list(orgId);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer id = json.getInteger("id");

		IsuWorkflow vo;
		if (id == null) {
			vo = new IsuWorkflow();
		} else {
			vo = issueWorkflowService.get(id, orgId);
		}

		List<IsuStatus> statuses = issueWorkflowService.listStatusForEdit(vo.getId(), orgId);

		ret.put("data", vo);
        ret.put("statuses", statuses);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "design", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> design(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();

        Integer id = json.getInteger("id");
        if (id == null) {
            ret.put("data", new IsuWorkflow());
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
            return ret;
        }

        IsuWorkflow vo = issueWorkflowService.get(id, orgId);
        List<IsuStatus> statuses = issueWorkflowService.listStatusForDesign(id);
        Map<String, IsuWorkflowTransition> tranMap = issueWorkflowService.getTransitionMap(id);

        List<IsuPage> pages = pageService.listForWorkflowTran(orgId);

        ret.put("data", vo);
        ret.put("statuses", statuses);
        ret.put("tranMap", tranMap);

        ret.put("pages", pages);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

        IsuWorkflow vo = json.getObject("model", IsuWorkflow.class);
        List<Integer> statusIds = json.getObject("statusIds", List.class);

        IsuWorkflow po = issueWorkflowService.save(vo, statusIds, orgId);

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

		issueWorkflowService.delete(id, orgId);

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

        Boolean result = issueWorkflowService.setDefault(id, orgId);
        if (!result) { // 当对象不是默认org的，结果会返回false
            return authFail();
        }

        List<IsuWorkflow> vos = issueWorkflowService.list(orgId);

        ret.put("data", vos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());

        return ret;
    }

}
