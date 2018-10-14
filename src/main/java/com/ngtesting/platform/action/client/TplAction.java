package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
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
import org.springframework.util.StringUtils;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.io.UnsupportedEncodingException;
import java.net.URLDecoder;
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

		String jql = json.getString("jql");
		if (StringUtils.isEmpty(jql) || "all".equals(jql)) {
            jql = isuJqlService.buildDefaultJql(orgId, projectId);
        } else {
            try {
                jql = URLDecoder.decode(jql, "UTF-8");
            } catch (UnsupportedEncodingException e) {
                e.printStackTrace();
            }
        }

        List<IsuJqlFilter> filters = isuJqlFilterService.buildUiFilters(jql, orgId, projectId);

        List<IsuIssue> data = isuJqlService.query(jql, orgId, projectId);

        ret.put("jql", JSON.parseObject(jql));
        ret.put("filters", filters);
        ret.put("data", data);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
