package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.itfsw.query.builder.support.model.JsonRule;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IsuJqlColumnService;
import com.ngtesting.platform.service.IsuJqlFilterService;
import com.ngtesting.platform.service.IsuJqlService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.vo.IsuJqlColumn;
import com.ngtesting.platform.vo.IsuJqlFilter;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.util.StringUtils;
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
    @Autowired
    IsuJqlColumnService isuJqlColumnService;
    @Autowired
    UserService userService;

	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();
        Integer projectId = user.getDefaultPrjId();

        Integer pageNum = json.getInteger("page");
        Integer pageSize = json.getInteger("pageSize");
        Boolean init = json.getBoolean("init");

        JsonRule rule;
        if (!json.getJSONObject("rule").containsKey("condition")) {
            rule = isuJqlService.buildDefaultJql(orgId, projectId);
        } else {
            rule = json.getObject("rule", JsonRule.class);
        }

        if (StringUtils.isEmpty(user.getIssueColumns())) {
            isuJqlColumnService.buildDefaultColStr(user);
        }

        com.github.pagehelper.Page page = PageHelper.startPage(pageNum, pageSize);
        List<IsuIssue> data = isuJqlService.query(rule, user.getIssueColumns(), orgId, projectId);

        if (init) {
            List<IsuJqlFilter> filters = isuJqlFilterService.buildUiFilters(rule, orgId, projectId);
            List<IsuJqlColumn> columns = isuJqlColumnService.loadColumns(user);

            ret.put("rule", rule);
            ret.put("filters", filters);
            ret.put("columns", columns);
        }

        ret.put("total", page.getTotal());
        ret.put("data", data);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

    @RequestMapping(value = "changeColumns", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> changeColumns(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        String columnsStr = json.getString("columns");

        userService.saveIssueColumns(columnsStr, user);

        ret.put("code", Constant.RespCode.SUCCESS.getCode());
        return ret;
    }

}
