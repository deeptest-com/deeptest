package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.service.OrgGroupService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.vo.OrgGroupVo;
import com.ngtesting.platform.vo.UserVo;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
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

	@AuthPassport(validate = true)
	@RequestMapping(value = "search", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		Long orgId = json.getLong("orgId");
		String keywords = json.getString("keywords");
        JSONArray exceptIds = json.getJSONArray("exceptIds");

		List userPos = userService.search(orgId, keywords, exceptIds);
		List<UserVo> userVos = userService.genVos(userPos);

		List groupPos = orgGroupService.search(orgId, keywords, exceptIds);
		List<OrgGroupVo> groupVos = orgGroupService.genVos(groupPos);

		List<Object> vos = new ArrayList<>();
		vos.addAll(groupVos);
		vos.addAll(userVos);

		ret.put("data", vos);
		ret.put("code", RespCode.SUCCESS.getCode());
		return ret;
	}

}
