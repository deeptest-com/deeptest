package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.*;
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
@RequestMapping(Constant.API_PATH_CLIENT + "org/")
public class OrgAction extends BaseAction {
	@Autowired
    OrgService orgService;
    @Autowired
    UserService userService;

	@Autowired
	TestPlanService planService;
	@Autowired
    HistoryService historyService;

	@Autowired
    PushSettingsService pushSettingsService;


	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		List<TstOrg> vos = orgService.list(userVo.getId(), keywords, disabled);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		Integer id = json.getInteger("id");

		TstOrg po = orgService.get(id);

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "view", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer id = json.getInteger("id");

		TstOrg po = orgService.get(id);

		List<TstPlan> planPos = planService.listByOrg(id);
		planService.genVos(planPos);

		List<TstHistory> historyPos = historyService.listByOrg(id);
		Map<String, List<TstHistory>> historyVos = historyService.genVosByDate(historyPos);

		ret.put("org", po);
		ret.put("plans", planPos);
		ret.put("histories", historyVos);
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody TstOrg vo) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser userVo = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		orgService.save(vo, userVo.getId());
        List<TstOrg> vos = orgService.list(userVo.getId(), "false", null);

        pushSettingsService.pushMyOrgs(userVo);

		ret.put("myOrgs", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Integer id = json.getInteger("id");

		boolean success = orgService.delete(id);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "change", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> change(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = json.getInteger("id");

		userService.setDefaultOrg(user, orgId);

		pushSettingsService.pushOrgSettings(user);
		pushSettingsService.pushRecentProjects(user);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}


	@RequestMapping(value = "setDefault", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Integer orgId = json.getInteger("id");
		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		userService.setDefaultOrg(user, orgId);
		pushSettingsService.pushOrgSettings(user);
		pushSettingsService.pushRecentProjects(user);

        List<TstOrg> vos = orgService.list(user.getId(), keywords, disabled);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
