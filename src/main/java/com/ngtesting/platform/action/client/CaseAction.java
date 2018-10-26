package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.*;
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
@RequestMapping(Constant.API_PATH_CLIENT + "case/")
public class CaseAction extends BaseAction {
	@Autowired
    ProjectService projectService;
	@Autowired
	ProjectDao projectDao;

	@Autowired
    CaseService caseService;
    @Autowired
    CaseTypeService caseTypeService;
    @Autowired
    CasePriorityService casePriorityService;
	@Autowired
    TestCustomFieldService customFieldService;


	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer orgId = user.getDefaultOrgId();
		Integer projectId = user.getDefaultPrjId();

		List<TstCase> ls = caseService.query(projectId);

        List<TstCaseType> caseTypePos = caseTypeService.list(orgId);
        List<TstCasePriority> casePriorityPos = casePriorityService.list(orgId);
        List<TstCustomField> customFieldList = customFieldService.listForCaseByProject(orgId, projectId);

        ret.put("data", ls);
        ret.put("prjId", projectId);
        ret.put("caseTypeList", caseTypePos);
        ret.put("casePriorityList", casePriorityPos);
		ret.put("customFields", customFieldList);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "queryForSuiteSelection", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> queryForSuiteSelection(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
		Integer suiteId = json.getInteger("suiteId");

        Integer prjId = caseProjectId == null? projectId: caseProjectId;
        if (userNotInProject(user.getId(), prjId)) {
            return authFail();
        }

        List<TstCase> vos = caseService.queryForSuiteSelection(prjId, suiteId);
		List<TstProject> projects = projectDao.listBrothers(projectId);

		ret.put("data", vos);
		ret.put("brotherProjects", projects);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "queryForTaskSelection", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> queryForTaskSelection(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
		Integer taskId = json.getInteger("taskId");

        Integer prjId = caseProjectId == null? projectId: caseProjectId;
        if (userNotInProject(user.getId(), prjId)) {
            return authFail();
        }

		List<TstCase> vos = caseService.queryForTaskSelection(prjId, taskId);
		List<TstProject> projects = projectDao.listBrothers(projectId);

		ret.put("data", vos);
		ret.put("brotherProjects", projects);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer prjId = user.getDefaultPrjId();
        Integer caseId = json.getInteger("id");

        TstCase testCase = caseService.getDetail(caseId, prjId);

        ret.put("data", testCase);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "rename", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        TstCase testCase = caseService.rename(json, user);
        if (testCase == null) {
            return authFail();
        }

        ret.put("data", testCase);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "move", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		TstCase testCase = caseService.move(json, user);
        if (testCase == null) {
            return authFail();
        }

		ret.put("data", testCase);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer count = caseService.delete(id, user);
        if (count == 0) {
            return authFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "update", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> update(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        TstCase testCase = caseService.update(json, user);
        if (testCase == null) {
            return authFail();
        }

        ret.put("data", testCase);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "saveField", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveField(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		TstCase testCase = caseService.saveField(json, user);
        if (testCase == null) {
            return authFail();
        }

        ret.put("data", testCase);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "changeContentType", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeContentType(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer id = json.getInteger("id");
        String contentType = json.getString("contentType");

        TstCase testCase = caseService.changeContentType(id, contentType, user);
        if (testCase == null) {
            return authFail();
        }

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "reviewResult", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> reviewResult(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer id = json.getInteger("id");
		Boolean result = json.getBoolean("result");

		TstCase testCase = caseService.reviewResult(id, result, user);
        if (testCase == null) {
            return authFail();
        }

        ret.put("reviewResult", testCase);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "exportAll", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> exportAll(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer projectId = json.getInteger("projectId");

        if (userNotInProject(user.getId(), projectId)) {
            return authFail();
        }

		String excelPath = caseService.export(projectId);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("excelPath", excelPath);

        return ret;
    }

}
