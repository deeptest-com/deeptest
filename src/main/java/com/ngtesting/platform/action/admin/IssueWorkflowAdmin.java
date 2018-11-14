package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuWorkflow;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IssueStatusService;
import com.ngtesting.platform.service.IssueWorkflowService;
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

	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

		List<IsuWorkflow> vos = issueWorkflowService.list(orgId);

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
		if (id == null) {
			ret.put("data", new IsuWorkflow());
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}

		IsuWorkflow vo = issueWorkflowService.get(id, orgId);

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

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = userVo.getDefaultOrgId();

        Integer id = json.getInteger("id");
        if (id == null) {
            ret.put("data", new IsuWorkflow());
            ret.put("code", Constant.RespCode.SUCCESS.getCode());
            return ret;
        }

        IsuWorkflow vo = issueWorkflowService.get(id, orgId);
        List<IsuStatus> statuses = issueWorkflowService.listStatusForDesign(id);

        ret.put("data", vo);
        ret.put("statuses", statuses);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

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

		Integer id = json.getInteger("id");

//		issueWorkflowService.delete(id);

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
//		boolean success = issueWorkflowService.setDefault(id, orgId);
//
//		List<CaseTypeVo> vos = issueWorkflowService.listVos(orgId);
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
//		boolean success = issueWorkflowService.changeOrder(id, act, orgId);
//
//		List<CaseTypeVo> vos = issueWorkflowService.listVos(orgId);
//
//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
