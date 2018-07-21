package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
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
        List<TstCase> vos = caseService.genVos(ls, false);

        List<TstCaseType> caseTypePos = caseTypeService.list(orgId);
        List<TstCaseType> caseTypeList = caseTypeService.genVos(caseTypePos);

        List<TstCasePriority> casePriorityPos = casePriorityService.list(orgId);
        List<TstCasePriority> casePriorityList = casePriorityService.genVos(casePriorityPos);

        List<TstCustomField> customFieldList = customFieldService.listForCaseByProject(orgId, projectId);

        ret.put("data", vos);
        ret.put("caseTypeList", caseTypeList);
        ret.put("casePriorityList", casePriorityList);
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
		List<TstProject> projects = projectService.listBrothers(projectId);

		ret.put("data", vos);
		ret.put("brotherProjects", projects);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "queryForRunSelection", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> queryForRunSelection(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
		Integer runId = json.getInteger("runId");

		List<TstCase> vos = caseService.queryForRunSelection(projectId, caseProjectId, runId);
		List<TstProject> projects = projectService.listBrothers(projectId);

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
        TstCase caseVo = caseService.genVo(testCasePo);

        ret.put("data", caseVo);
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

		TstCase testCase = caseService.delete(id, userVo);

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

    @RequestMapping(value = "save", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        TstCase po = caseService.save(json, userVo);
        TstCase caseVo = caseService.genVo(po, true);

        ret.put("data", caseVo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "saveField", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveField(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		TstCase po = caseService.saveField(json, userVo);
        TstCase caseVo = caseService.genVo(po);

		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "changeContentType", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeContentType(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");
        String contentType = json.getString("contentType");

		TstCase po = caseService.changeContentTypePers(id, contentType);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "reviewPass", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> reviewPass(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");
		Boolean pass = json.getBoolean("pass");

		TstCase po = caseService.reviewPassPers(id, pass);
        TstCase caseVo = caseService.genVo(po);

        ret.put("reviewResult", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
