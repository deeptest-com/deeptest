package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestHistory;
import com.ngtesting.platform.entity.TestOrg;
import com.ngtesting.platform.entity.TestPlan;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.service.PlanService;
import com.ngtesting.platform.service.PushSettingsService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.TestHistoryVo;
import com.ngtesting.platform.vo.TestPlanVo;
import com.ngtesting.platform.vo.UserVo;
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
	PlanService planService;
	@Autowired
	HistoryService historyService;

	@Autowired
	PushSettingsService pushSettingsService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");

		List<OrgVo> vos = orgService.listVo(keywords, disabled, userVo.getId());

        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		Long id = json.getLong("id");

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		if (id != null) {
			TestOrg po = (TestOrg) orgService.get(TestOrg.class, id);
			OrgVo vo = orgService.genVo(po);

			TestUser user = (TestUser)orgService.get(TestUser.class, userVo.getId());
			if (po.getId().longValue() == user.getDefaultOrgId().longValue()) {
				vo.setDefaultOrg(true);
			}

	        ret.put("data", vo);
		}

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "view", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> view(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long id = json.getLong("id");

		TestOrg po = (TestOrg) orgService.get(TestOrg.class, id);
		OrgVo vo = orgService.genVo(po);

		List<TestPlan> planPos = planService.listByOrg(id);
		List<TestPlanVo> planVos = planService.genVos(planPos);

		List<TestHistory> historyPos = historyService.listByOrg(id);
		Map<String, List<TestHistoryVo>> historyVos = historyService.genVosByDate(historyPos);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		ret.put("org", vo);
		ret.put("plans", planVos);
		ret.put("histories", historyVos);

		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody OrgVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		orgService.save(vo, userVo.getId());
		List<OrgVo> vos = orgService.listVo(null, "false", userVo.getId());

        pushSettingsService.pushMyOrgs(userVo);

		ret.put("myOrgs", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long id = json.getLong("id");

		boolean success = orgService.delete(id);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "change", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> change(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = json.getLong("id");

		orgService.setDefaultPers(orgId, userVo);

		pushSettingsService.pushOrgSettings(userVo);
		pushSettingsService.pushRecentProjects(userVo);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "setDefault", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setDefault(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = json.getLong("id");
		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");

		orgService.setDefaultPers(orgId, userVo);
		pushSettingsService.pushOrgSettings(userVo);
		pushSettingsService.pushRecentProjects(userVo);

		List<OrgVo> vos = orgService.listVo(keywords, disabled, userVo.getId());

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

}
