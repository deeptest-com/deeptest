package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.bean.websocket.WsFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.IsuFilterService;
import com.ngtesting.platform.service.IsuJqlService;
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
@RequestMapping(Constant.API_PATH_CLIENT + "tql/")
public class TplAction extends BaseAction {

	@Autowired
	private WsFacade optFacade;

	@Autowired
    IsuJqlService isuJqlService;
	@Autowired
	IsuFilterService isuFilterService;

	@RequestMapping(value = "query", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> query(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();
        Integer projectId = user.getDefaultPrjId();

		String jqlStr = json.getString("jql");
		if (jqlStr == null || "all".equals(jqlStr)) {
            jqlStr = "";
        }

        Map<String, Object> data = isuJqlService.query(jqlStr, orgId, projectId);

        ret.put("data", data);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

}
