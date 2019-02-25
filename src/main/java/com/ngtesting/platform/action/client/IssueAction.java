package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.*;
import com.ngtesting.platform.servlet.PrivPrj;
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
    @Autowired
    IssueMiscService issueMiscService;

    @Autowired
    IssueLinkService issueLinkService;
    @Autowired
    IssueTagService issueTagService;
    @Autowired
    IssueWatchService issueWatchService;


    @PrivPrj(perms = {"issue-view"})
    @RequestMapping(value = "view", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");
        IsuIssue po = issueService.getDetail(id, user.getId(), prjId);

        if (po == null) { // 当对象不是默认org的，此处为空
            return authFail();
        }

        IsuPage page = issueService.getPage(orgId, prjId, "view");

        ret.put("data", po);
        ret.put("page", page);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PrivPrj(perms = {"issue-view"})
    @RequestMapping(value = "getData", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> getData(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");
        IsuIssue po = issueService.getData(id, user.getId(), prjId);

        ret.put("data", po);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "create", method = RequestMethod.POST)
    @ResponseBody
    @PrivPrj(perms = {"issue-maintain"})
    public Map<String, Object> create(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

		IsuIssue po = new IsuIssue();
        po.setReporterId(user.getId());

        IsuPage page = issueService.getPage(orgId, prjId, "create");

        ret.put("data", po);
        ret.put("page", page);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PrivPrj(perms = {"issue-maintain"})
    @RequestMapping(value = "edit", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> edit(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
        Integer prjId = user.getDefaultPrjId();

        Integer id = json.getInteger("id");
        IsuIssue po = issueService.get(id, user.getId(), prjId);

        if (po == null) { // 当对象不是默认org的，此处为空
            return authFail();
        }

        IsuPage page = issueService.getPage(orgId, prjId, "edit");

        ret.put("data", po);
        ret.put("page", page);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PrivPrj(perms = {"issue-maintain"})
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        Integer pageId = json.getInteger("pageId");
        JSONObject issue = json.getJSONObject("issue");

        IsuIssue po = issueService.save(issue, pageId, user);

		ret.put("id", po.getId());
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @PrivPrj(perms = {"issue-maintain"})
    @RequestMapping(value = "update", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> update(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer prjId = user.getDefaultPrjId();

        Integer pageId = json.getInteger("pageId");
        JSONObject issue = json.getJSONObject("issue");
        Integer id = issue.getInteger("id");

        issueService.update(issue, pageId, user);

        ret.put("id", issue.getInteger("id"));
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PrivPrj(perms = {"issue-maintain"})
	@RequestMapping(value = "delete", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        Integer id = json.getInteger("id");

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        issueService.delete(id, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @PrivPrj(perms = {"issue-maintain"})
	@RequestMapping(value = "updateField", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> updateField(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		IsuIssue po = issueService.updateField(json, user);

        ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
