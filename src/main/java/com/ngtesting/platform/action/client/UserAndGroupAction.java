package com.ngtesting.platform.action.client;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.BaseAction;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.OrgGroupService;
import com.ngtesting.platform.service.intf.UserService;
import com.ngtesting.platform.servlet.PrivOrg;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;


@Controller
@RequestMapping(Constant.API_PATH_CLIENT + "userAndGroup/")
public class UserAndGroupAction extends BaseAction {
	@Autowired
    UserService userService;
	@Autowired
    OrgGroupService orgGroupService;

	@PostMapping(value = "search")
	@ResponseBody
	@PrivOrg
	public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = user.getDefaultOrgId();

		String keywords = json.getString("keywords");
        List<Integer> exceptUserIds = json.getObject("exceptUserIds", List.class);
		List<Integer> exceptGroupIds = json.getObject("exceptGroupIds", List.class);

		List users = userService.searchOrgUser(orgId, keywords, exceptUserIds);
		List groups = orgGroupService.search(orgId, keywords, exceptGroupIds);

		List<Object> vos = new ArrayList<>();
		vos.addAll(users);
		vos.addAll(groups);

		ret.put("data", vos);
		ret.put("code", RespCode.SUCCESS.getCode());
		return ret;
	}

}
