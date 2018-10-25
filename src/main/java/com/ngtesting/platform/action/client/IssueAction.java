package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "issue/")
public class IssueAction extends BaseAction {
	@Autowired
    ProjectService projectService;
	@Autowired
    IssueService issueService;
    @Autowired
    IssueFieldService fieldService;
    @Autowired
    IssuePropertyService propertyService;

    @Autowired
    IssueTypeService typeService;
    @Autowired
    IssuePriorityService priorityService;
    @Autowired
    IssueStatusService statusService;
    @Autowired
    IssueResolutionService resolutionService;
	@Autowired
    IssueCustomFieldService customFieldService;

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();
        Integer caseId = json.getInteger("id");

		IsuIssue vo = issueService.get(caseId, orgId);
//        List<IsuField> fields = fieldService.list(user);
        Map propMap = propertyService.getProps(orgId);

        ret.put("data", vo);
//        ret.put("fields", fields);
        ret.put("propMap", propMap);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }


	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
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

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

//		TstCase testCasePo = issueService.rename(json, userVo);
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

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

//		TstCase testCase = issueService.delete(id, userVo);
//		issueService.updateParentIfNeededPers(testCase.getpId());

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "move", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> move(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer srcId = json.getInteger("srcId");
//        Integer parentId = issueService.getWithCasesById(srcId).getpId();
//        Integer targetId = json.getInteger("targetId");
//        TstCaseVo vo = issueService.move(json, userVo);
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

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

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

//		TstCase po = issueService.changeContentType(id, contentType);

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
