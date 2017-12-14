package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.entity.TestVerifyCode;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.MailService;
import com.ngtesting.platform.service.RelationOrgGroupUserService;
import com.ngtesting.platform.service.UserService;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.util.PropertyConfig;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.RelationOrgGroupUserVo;
import com.ngtesting.platform.vo.RelationProjectRoleEntityVo;
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
@RequestMapping(Constant.API_PATH_CLIENT + "user/")
public class UserAction extends BaseAction {
	@Autowired
	UserService userService;
    @Autowired
    AccountService accountService;
	@Autowired
	RelationOrgGroupUserService orgGroupUserService;
	@Autowired
	MailService mailService;

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

		List <RelationProjectRoleEntityVo> vos = userService.getProjectUsers(orgId, Long.valueOf(projectId));

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
		boolean success = orgGroupUserService.saveRelations(relations);
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

		UserVo user = JSON.parseObject(JSON.toJSONString(json.get("user")), UserVo.class);
		TestUser po = userService.invitePers(user, orgId);

		if (po == null) {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "邮箱已存在");
			return ret;
		}

		List<RelationOrgGroupUserVo> relations = (List<RelationOrgGroupUserVo>) json.get("relations");
		orgGroupUserService.saveRelations(relations);

        TestVerifyCode verifyCode = accountService.forgotPasswordPers(po.getId());
		String sys = PropertyConfig.getConfig("sys.name");
		Map<String, String> map = new HashMap<String, String>();
		map.put("user", userVo.getName() + "(" + userVo.getEmail() + ")");
		map.put("name", user.getName());
        map.put("vcode", verifyCode.getCode());
		map.put("sys", sys);
        map.put("url", PropertyConfig.getConfig("url.reset.password"));
		mailService.sendTemplateMail("来自[" + sys + "]的邀请", "invite-user.ftl", user.getEmail(), map);

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
	@RequestMapping(value = "setSize", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> setSize(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		Integer left = json.getInteger("left");
		Integer right = json.getInteger("right");

		boolean success = userService.setSizePers(userVo.getId(), left, right);

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

}
