package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.*;
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
@RequestMapping(Constant.API_PATH_CLIENT + "user/")
public class UserAction extends BaseAction {
	@Autowired
	UserService userService;
    @Autowired
    AccountService accountService;
	@Autowired
	RelationOrgGroupUserService orgGroupUserService;
	@Autowired
	PushSettingsService pushSettingsService;

    @Autowired
    OrgService orgService;
    @Autowired
    ProjectService projectService;
    @Autowired
    SysPrivilegeService sysPrivilegeService;
    @Autowired
    OrgRolePrivilegeService orgRolePrivilegeService;
    @Autowired
    ProjectPrivilegeService projectPrivilegeService;
    @Autowired
    CasePropertyService casePropertyService;

	@AuthPassport(validate = true)
	@RequestMapping(value = "list", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> list(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		String keywords = json.getString("keywords");
		String disabled = json.getString("disabled");
		int page = json.getInteger("page") == null? 0: json.getInteger("page") - 1;
		int pageSize = json.getInteger("pageSize") == null? Constant.PAGE_SIZE: json.getInteger("pageSize");

		Page pageDate = userService.listByPage(orgId, keywords, disabled, page, pageSize);
		List<UserVo> vos = userService.genVos(pageDate.getItems());

		ret.put("collectionSize", pageDate.getTotal());
        ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "getUsers", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getUsers(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		String projectId = json.getString("projectId");

		List <Map> vos = userService.getProjectUsers(orgId, Long.valueOf(projectId));

		ret.put("data", vos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "get", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> get(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		Long userId = json.getLong("id");

		List<RelationOrgGroupUserVo> relations = orgGroupUserService.listRelationsByUser(orgId, userId);

		if (userId == null) {
			ret.put("user", new UserVo());
	        ret.put("relations", relations);
			ret.put("code", Constant.RespCode.SUCCESS.getCode());
			return ret;
		}

		TestUser po = (TestUser) userService.get(TestUser.class, Long.valueOf(userId));
		UserVo vo = userService.genVo(po);

        ret.put("user", vo);
        ret.put("relations", relations);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "save", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> save(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();

		UserVo user = JSON.parseObject(JSON.toJSONString(json.get("user")), UserVo.class);
		TestUser po = userService.save(user, orgId);

		if (po == null) {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "邮箱已存在");
			return ret;
		}

		List<RelationOrgGroupUserVo> relations = (List<RelationOrgGroupUserVo>) json.get("relations");
		orgGroupUserService.saveRelations(relations);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());

		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "invite", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> invite(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
        Long prjId = userVo.getDefaultPrjId();

		UserVo user = JSON.parseObject(JSON.toJSONString(json.get("user")), UserVo.class);
		List<RelationOrgGroupUserVo> relations = (List<RelationOrgGroupUserVo>) json.get("relations");
		TestUser po = userService.invitePers(userVo, user, relations);

		if (po == null) {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "邮箱已加入当期组织");
			return ret;
		}

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "disable", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> disable(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long userId = json.getLong("id");
		Long orgId = json.getLong("orgId");

		boolean success = userService.disable(json.getLong("id"), orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "delete", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> delete(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		Long userId = json.getLong("id");
		Long orgId = json.getLong("orgId");

		boolean success = userService.remove(userId, orgId);

		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "search", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> search(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = json.getLong("orgId");
		String keywords = json.getString("keywords");

		List userPos = userService.search(orgId, keywords, null);
		List<UserVo> userVos = userService.genVos(userPos);

		ret.put("data", userVos);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "saveInfo", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveInfo(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		json.put("id", userVo.getId());

		TestUser user = accountService.saveInfo(json);
		userVo = userService.genVo(user);
		request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userVo);

		pushSettingsService.pushUserSettings(userVo);

		ret.put("data", userVo);
		ret.put("code", RespCode.SUCCESS.getCode());
		return ret;
	}

	@AuthPassport(validate = true)
	@RequestMapping(value = "setLeftSize", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setLeftSize(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		Integer left = json.getInteger("left");

		TestUser user = accountService.setLeftSizePers(userVo.getId(), left);
		userVo = userService.genVo(user);
		request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userVo);

		ret.put("data", userVo);
		ret.put("code", Constant.RespCode.SUCCESS.getCode());
		return ret;
	}

	@RequestMapping(value = "getProfile", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getProfile(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		Long orgId = userVo.getDefaultOrgId();
		Long prjId = userVo.getDefaultPrjId();

		Long orgIdNew = json.getLong("orgId");
		Long prjIdNew = json.getLong("prjId");

		if (orgIdNew != null && orgIdNew.longValue() != orgId.longValue()) { // org不能为空
			orgService.setDefaultPers(orgId, userVo);
		}
		if (prjIdNew != null && (prjId == null || prjIdNew.longValue() != prjId.longValue())) { // prj可能为空
			projectService.viewPers(prjIdNew, userVo);
		}

		Long userId = userVo.getId();

		ret.put("profile", userVo);

		Map<String, Boolean> sysPrivileges = sysPrivilegeService.listByUser(userId);
		ret.put("sysPrivileges", sysPrivileges);

		List<OrgVo> orgs = orgService.listVo(null, "false", userId);
		ret.put("myOrgs", orgs);

		Map<String, Boolean> orgPrivileges = orgRolePrivilegeService.listByUser(userVo.getId(), orgId);
		ret.put("orgPrivileges", orgPrivileges);

		Map<String,Map<String,String>> casePropertyMap = casePropertyService.getMap(orgId);
		ret.put("casePropertyMap", casePropertyMap);

		List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(orgId, userId);
		ret.put("recentProjects", recentProjects);
		if (recentProjects.size() > 0) {
			userVo.setDefaultPrjId(recentProjects.get(0).getProjectId());
		}

		if (userVo.getDefaultPrjId() != null) {
			Map<String, Boolean> prjPrivileges = projectPrivilegeService.listByUserPers(userId, prjId, orgId);
			ret.put("prjPrivileges", prjPrivileges);
		}

		ret.put("code", RespCode.SUCCESS.getCode());

		return ret;
	}

}
