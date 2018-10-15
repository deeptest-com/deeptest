package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.IsuJqlService;
import com.ngtesting.platform.vo.IsuJqlFilter;
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
@RequestMapping(Constant.API_PATH_CLIENT + "tql/")
public class TplAction extends BaseAction {

	@Autowired
	private WsFacade optFacade;

	@Autowired
    IsuJqlService isuJqlService;
    @Autowired
    IsuJqlFilterService isuJqlFilterService;

	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();
        Integer projectId = user.getDefaultPrjId();

        Boolean init = json.getBoolean("init");

        JsonRule rule;
        if (!json.getJSONObject("jql").containsKey("condition")) {
            rule = isuJqlService.buildDefaultJql(orgId, projectId);
        } else {
            rule = json.getObject("jql", JsonRule.class);
        }

        List<IsuIssue> data = isuJqlService.query(rule, orgId, projectId);

        if (init) {
            List<IsuJqlFilter> filters = isuJqlFilterService.buildUiFilters(rule, orgId, projectId);
            ret.put("jql", rule);
            ret.put("filters", filters);
        }

        ret.put("data", data);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
