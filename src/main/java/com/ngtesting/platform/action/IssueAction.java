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
@RequestMapping(Constant.API_PATH_CLIENT + "issue/")
public class IssueAction extends BaseAction {
	@Autowired
    ProjectService projectService;
	@Autowired
    IssueService issueService;
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

		List<TstCase> ls = issueService.query(projectId);
//        List<TstCase> vos = issueService.genVos(ls, false);

        List<TstCaseType> caseTypePos = caseTypeService.list(orgId);

        List<TstCasePriority> casePriorityPos = casePriorityService.list(orgId);

        List<TstCustomField> customFieldList = customFieldService.listForCaseByProject(orgId, projectId);

//        ret.put("data", vos);
        ret.put("caseTypeList", caseTypePos);
        ret.put("casePriorityList", casePriorityPos);
		ret.put("customFields", customFieldList);
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

		TstCase vo = issueService.getById(caseId);

        ret.put("data", vo);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }


	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
//
//		IsuIssue po = issueService.save(json, userVo);
//		TstCase caseVo = issueService.genVo(po, true);
//
//		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "rename", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> rename(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

//		TstCase testCasePo = issueService.renamePers(json, userVo);
//        issueService.updateParentIfNeededPers(testCasePo.getpId());
//		TstCaseVo caseVo = issueService.genVo(testCasePo);
//
//		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

//		TstCase testCase = issueService.delete(id, userVo);
//		issueService.updateParentIfNeededPers(testCase.getpId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "move", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

        Integer srcId = json.getInteger("srcId");
//        Integer parentId = issueService.getWithCasesById(srcId).getpId();
//        Integer targetId = json.getInteger("targetId");
//        TstCaseVo vo = issueService.movePers(json, userVo);
//
//        issueService.updateParentIfNeededPers(parentId);
//        issueService.updateParentIfNeededPers(targetId);

//		ret.put("data", vo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "saveField", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveField(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

//		TstCase po = issueService.saveField(json, userVo);
//        TstCaseVo caseVo = issueService.genVo(po);
//
//		ret.put("data", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "changeContentType", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeContentType(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");
        String contentType = json.getString("contentType");

//		TstCase po = issueService.changeContentTypePers(id, contentType);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "reviewPass", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> reviewPass(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");
		Boolean pass = json.getBoolean("pass");

//		TstCase po = issueService.reviewResult(id, pass);
//        TstCaseVo caseVo = issueService.genVo(po);
//
//        ret.put("reviewResult", caseVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
