package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.IssueFieldService;
import com.ngtesting.platform.service.intf.IssueService;
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
    IssueService issueService;
    @Autowired
    IssueFieldService fieldService;
	@Autowired
	IssueDynamicFormService dynamicFormService;

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        String opt = json.getString("opt");
		Integer id = json.getInteger("id");
		IsuIssue po;

		if (id == null) {
			po = new IsuIssue();
		} else {
			po = issueService.get(id, orgId);
		}

		if (po == null) { // 当对象不是默认org的，此处为空
			return authFail();
		}

        IsuPage page = issueService.getPage(orgId, prjId, opt);

        Map<String, Object> issuePropMap = dynamicFormService.fetchOrgField(orgId, prjId);

        ret.put("data", po);
        ret.put("page", page);
        ret.put("issuePropMap", issuePropMap);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer pageId = json.getInteger("pageId");
        JSONObject issue = json.getJSONObject("issue");

        IsuIssue po = issueService.save(issue, pageId, user);

		ret.put("data", po);
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

}
