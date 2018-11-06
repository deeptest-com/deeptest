package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuPage;
import com.ngtesting.platform.model.IsuPageTab;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IssueFieldService;
import com.ngtesting.platform.service.IssuePageService;
import com.ngtesting.platform.service.IssuePageTabService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_page_tab/")
public class IssuePageTabAdmin extends BaseAction {
    @Autowired
    IssuePageService pageService;
	@Autowired
    IssuePageTabService tabService;

    @Autowired
    IssueFieldService fieldService;

    @RequestMapping(value = "add", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> add(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = userVo.getDefaultOrgId();

        IsuPageTab tab = JSON.parseObject(JSON.toJSONString(json), IsuPageTab.class);
        tab.setOrgId(orgId);

        tabService.add(tab);

        List<IsuField> fields = fieldService.listOrgField(orgId, tab.getId());

        ret.put("tab", tab);
        ret.put("fields", fields);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "get", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = userVo.getDefaultOrgId();

        Integer tabId = json.getInteger("id");

        IsuPageTab tab = tabService.get(tabId, orgId);

        List<IsuField> fields = fieldService.listOrgField(orgId, tab.getId());

        ret.put("tab", tab);
        ret.put("fields", fields);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "remove", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> remove(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = userVo.getDefaultOrgId();

        Integer id = json.getInteger("id");
        Integer pageId = json.getInteger("pageId");
        Integer currTabId = json.getInteger("currTabId");

        boolean success = tabService.remove(id, pageId, orgId);
        if (!success) {
            ret.put("code", Constant.RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "不能删除最后一个标签");

            return ret;
        }

        IsuPage page = pageService.get(pageId, orgId);

        IsuPageTab tab;
        if (currTabId.intValue() == id.intValue()) {
            tab = page.getTabs().get(0);
            List<IsuField> fields = fieldService.listOrgField(orgId, tab.getId());
            ret.put("fields", fields);
        }

        ret.put("page", page);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

    @RequestMapping(value = "updateName", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> updateName(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = userVo.getDefaultOrgId();

        IsuPageTab tab = JSON.parseObject(JSON.toJSONString(json), IsuPageTab.class);

        tabService.updateName(tab);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

	@RequestMapping(value = "changeOrder", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changeOrder(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();
		Integer id = json.getInteger("id");
		String act = json.getString("act");

//		boolean success = customFieldService.changeOrderPers(id, act, orgId);
//
//		List<IsuCustomField> vos = customFieldService.list(orgId);

//        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
