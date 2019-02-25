package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.IssueFieldService;
import com.ngtesting.platform.service.intf.IssuePageService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_page/")
public class IssuePageAdmin extends BaseAction {
	@Autowired
	IssuePageService pageService;

    @Autowired
    IssueFieldService fieldService;
    @Autowired
    IssueDynamicFormService dynamicFormService;

	@RequestMapping(value = "load", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> load(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		List<IsuPage> pages = pageService.listAll(orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("pages", pages);
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		Integer pageId = json.getInteger("id");
		IsuPage page = null;
		if (pageId == null) {
			page = new IsuPage();
		} else {
			page = pageService.get(pageId, orgId);
		}

		ret.put("page", page);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "getDetail", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> getDetail(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = user.getDefaultOrgId();
		Integer projectId = user.getDefaultPrjId();

        Integer pageId = json.getInteger("id");
        IsuPage page = pageService.get(pageId, orgId);

        List<IsuField> fields = dynamicFormService.listNotUsedField(orgId, projectId, pageId);
        Map issuePropMap = dynamicFormService.genIssuePropMap(orgId, projectId);

        ret.put("page", page);
        ret.put("fields", fields);
        ret.put("issuePropMap", issuePropMap);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();
		Integer projectId = user.getDefaultPrjId();

		IsuPage page = JSON.parseObject(JSON.toJSONString(json), IsuPage.class);
        pageService.save(page, orgId);

		page = pageService.get(page.getId(), orgId);

		List<IsuField> fields = dynamicFormService.listNotUsedField(orgId, projectId, page.getId());

		ret.put("page", page);
        ret.put("fields", fields);

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

		boolean success = pageService.delete(id, orgId);

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

		Boolean result = pageService.setDefault(id, orgId);
		if (!result) { // 当对象不是默认org的，结果会返回false
			return authFail();
		}

        List<IsuPage> pages = pageService.listAll(orgId);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        ret.put("pages", pages);

		return ret;
	}

}
