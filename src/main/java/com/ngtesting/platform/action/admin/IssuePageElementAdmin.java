package com.ngtesting.platform.action.admin;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuField;
import com.ngtesting.platform.model.IsuPageTab;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueDynamicFormService;
import com.ngtesting.platform.service.intf.IssueFieldService;
import com.ngtesting.platform.service.intf.IssuePageElementService;
import com.ngtesting.platform.service.intf.IssuePageTabService;
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
@RequestMapping(Constant.API_PATH_ADMIN + "issue_page_elem/")
public class IssuePageElementAdmin extends BaseAction {

	@Autowired
	IssuePageTabService tabService;
	@Autowired
	IssuePageElementService elementService;

    @Autowired
	IssueFieldService fieldService;
    @Autowired
    IssueDynamicFormService dynamicFormService;

	@RequestMapping(value = "saveAll", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveAll(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = userVo.getDefaultOrgId();

        Integer pageId = json.getInteger("pageId");
        Integer tabId = json.getInteger("tabId");
        List<Map> maps = JSON.parseArray(json.getJSONArray("elems").toJSONString(), Map.class) ;

        elementService.saveAll(orgId, pageId, tabId, maps);

		IsuPageTab tab = tabService.get(tabId, orgId);
		List<IsuField> fields = dynamicFormService.listTabNotUsedField(orgId, tabId);

		ret.put("tab", tab);
		ret.put("fields", fields);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "updateProp", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> updateProp(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
        Integer orgId = userVo.getDefaultOrgId();

        String id = json.getString("id");
        String prop = json.getString("prop");
        String val = json.getString("val");

        elementService.updateProp(id, prop, val, orgId);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
