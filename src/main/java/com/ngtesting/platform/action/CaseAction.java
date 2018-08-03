package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
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
    CustomFieldService customFieldService;


	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        Integer orgId = json.getInteger("orgId");
		Integer projectId = json.getInteger("projectId");

		List<TstCase> ls = caseService.query(projectId);

        List<TstCaseType> caseTypePos = caseTypeService.list(orgId);
        List<TstCasePriority> casePriorityPos = casePriorityService.list(orgId);
        List<TstCustomField> customFieldList = customFieldService.listForCaseByProject(orgId, projectId);

        ret.put("data", ls);
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

		Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
		Integer suiteId = json.getInteger("suiteId");

        List<TstCase> vos = caseService.queryForSuiteSelection(projectId, caseProjectId, suiteId);
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

		Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
		Integer taskId = json.getInteger("taskId");

		List<TstCase> vos = caseService.queryForTaskSelection(projectId, caseProjectId, taskId);
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

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
        Integer orgId = userVo.getDefaultOrgId();
        Integer caseId = json.getInteger("id");

        TstCase vo = caseService.getById(caseId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "rename", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        TstCase testCasePo = caseService.renamePers(json, userVo);

        ret.put("data", testCasePo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "move", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		TstCase vo = caseService.movePers(json, userVo);

		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		caseService.delete(id, userVo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "exportAll", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> exportAll(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer projectId = json.getInteger("projectId");

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

//		String excelPath = caseService.export(projectId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
//		ret.put("excelPath", excelPath);

		return ret;
	}

    @RequestMapping(value = "update", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> update(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        TstCase po = caseService.update(json, userVo);

        ret.put("data", po);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "saveField", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveField(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TstCase po = caseService.saveField(json, userVo);

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "changeContentType", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeContentType(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");
        String contentType = json.getString("contentType");

		caseService.changeContentTypePers(id, contentType);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "reviewResult", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> reviewResult(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");
		Boolean result = json.getBoolean("result");

		TstCase po = caseService.reviewResult(id, result);

        ret.put("reviewResult", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
