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

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

		List<TstOrg> vos = orgService.list(user.getId(), keywords, disabled);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer orgId = json.getInteger("id");

        if (userNotInOrg(user.getId(), orgId)) {
            return authFail();
        }

		TstOrg po = orgService.get(orgId);

		ret.put("data", po);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}


	@RequestMapping(value = "view", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer orgId = json.getInteger("id");
        if (userNotInOrg(user.getId(), orgId)) {
            return authFail();
        }

		TstOrg po = orgService.get(orgId);

		List<TstPlan> planPos = planService.listByOrg(orgId);
		planService.genVos(planPos);

		List<TstHistory> historyPos = historyService.listByOrg(orgId);
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
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

        if (vo.getId() != null && authDao.userNotInOrg(user.getId(), vo.getId())) {
            return authFail();
        }

        TstOrg org = orgService.save(vo, user);
        if (user.getDefaultOrgId().intValue() == org.getId().intValue() &&
                !user.getDefaultOrgName().equals(org.getName())) {
            user.setDefaultOrgName(org.getName());
            pushSettingsService.pushOrgSettings(user);
        }

        pushSettingsService.pushMyOrgs(user);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
        TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer orgId = json.getInteger("id");
        if (userNotInOrg(user.getId(), orgId)) {
            return authFail();
        }

		Boolean result = orgService.delete(orgId, user);
        if (result && orgId.intValue() == user.getDefaultOrgId().intValue()) {
            userService.setEmptyOrg(user, orgId);

            pushSettingsService.pushMyOrgs(user);
            pushSettingsService.pushOrgSettings(user);
            pushSettingsService.pushRecentProjects(user);
        }
        ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "change", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> change(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);
		Integer orgId = json.getInteger("id");
        if (userNotInOrg(user.getId(), orgId)) {
            return authFail();
        }

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
		TstUser user = (TstUser) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_PROFILE);

		Integer orgId = json.getInteger("id");
		String keywords = json.getString("keywords");
		Boolean disabled = json.getBoolean("disabled");

        if (userNotInOrg(user.getId(), orgId)) {
            return authFail();
        }

		userService.setDefaultOrg(user, orgId);
		pushSettingsService.pushOrgSettings(user);
		pushSettingsService.pushRecentProjects(user);

        List<TstOrg> vos = orgService.list(user.getId(), keywords, disabled);

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
