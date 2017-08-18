package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.entity.TestVerifyCode;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.util.Constant;
import com.ngtesting.platform.util.Constant.RespCode;
import com.ngtesting.platform.util.PropertyConfig;
import com.ngtesting.platform.vo.OrgVo;
import com.ngtesting.platform.vo.TestProjectAccessHistoryVo;
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
@RequestMapping(Constant.API_PATH_CLIENT + "account/")
public class AccountAction extends BaseAction {
	@Autowired
	AccountService accountService;
	@Autowired
	UserService userService;
	@Autowired
    ProjectService projectService;
	
	@Autowired
	RegisterService registerService;
	
	@Autowired
	MailService mailService;

    @Autowired
    OrgService orgService;
	
	@AuthPassport(validate=false)
	@RequestMapping(value = "login", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> login(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String email = json.getString("email");
		String password = json.getString("password");
		boolean rememberMe = json.getBoolean("rememberMe") != null? json.getBoolean("rememberMe"): false;
		
		TestUser user = accountService.loginPers(email, password, rememberMe);
			
		if (user != null) {
			UserVo userVo = userService.genVo(user);	
			request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userVo);
			
//			List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(user.getDefaultOrgId(), user.getDefaultProjectId());		
//			ret.put("profile", userVo);
//			ret.put("recentProjects", recentProjects);
			ret.put("token", user.getToken());
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "登录失败");
		}

		return ret;
	}
	
	@AuthPassport(validate=false)
	@RequestMapping(value = "register", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> register(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String name = json.getString("name");
		String phone = json.getString("phone");
		String email = json.getString("email");
		String password = json.getString("password");

		TestUser user = accountService.registerPers(name, email, phone, password);

		if (user != null) {
			UserVo userVo = userService.genVo(user);	
			request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userVo);
			
//			List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(user.getDefaultOrgId(), user.getDefaultProjectId());		
//			ret.put("profile", userVo);
//			ret.put("recentProjects", recentProjects);
			ret.put("token", user.getToken());
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "邮箱已存在");
		}

		return ret;
	}
	

	@AuthPassport(validate=false)
	@RequestMapping(value = "resetPassword", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> resetPassword(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String verifyCode = json.getString("vcode");
		String password = json.getString("password");

		TestUser user = accountService.resetPasswordPers(verifyCode, password);
		
		if (user != null) {
			UserVo userVo = userService.genVo(user);	
			request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userVo);
			
//			List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(user.getDefaultOrgId(), userVo.getDefaultProjectId());
//			ret.put("profile", userVo);
//			ret.put("recentProjects", recentProjects);
			ret.put("token", user.getToken());
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("data", "重置密码失败");
		}

		return ret;
	}
	
	@RequestMapping(value = "getProfile", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getProfile(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo userVo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		
		List<TestProjectAccessHistoryVo> recentProjects = projectService.listRecentProjectVo(userVo.getDefaultOrgId(), userVo.getId());
        List<OrgVo> orgs = orgService.listVo(null, "false", userVo.getId());

		ret.put("profile", userVo);
		ret.put("recentProjects", recentProjects);
        ret.put("myOrgs", orgs);
		
		ret.put("code", RespCode.SUCCESS.getCode());

		return ret;
	}

	@AuthPassport(validate=false)
	@RequestMapping(value = "forgotPassword", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> forgotPassword(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String email = json.getString("email");
		TestUser user = (TestUser) accountService.getByEmail(email);
		if (user == null) {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "用户不存在");
		}
		
		TestVerifyCode verifyCode = accountService.forgotPasswordPers(user.getId());
		if (verifyCode != null) {
			Map<String, String> map = new HashMap<String, String>();
			map.put("name", user.getName());
			map.put("vcode", verifyCode.getCode());
			// map.put("url", Constant.WEB_ROOT + "admin-path");
			map.put("url", PropertyConfig.getConfig("admin.url.forgot.password"));
			mailService.sendTemplateMail("[聆客]忘记密码", "forgot-password.ftl", user.getEmail(), map);
			
			ret.put("data", verifyCode);
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "用户不存在");
		}

		return ret;
	}
	
	
	@AuthPassport(validate=true)
	@RequestMapping(value = "logout", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> logout(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo vo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		if (vo == null) {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "您不在登录状态");
			return ret;
		}
		TestUser user = accountService.logoutPers(vo.getEmail());
		
		if (user != null) {
			request.getSession().removeAttribute(Constant.HTTP_SESSION_USER_KEY);
			
			ret.put("token", user.getToken());
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "登出失败");
		}

		return ret;
	}

	@RequestMapping(value = "saveProfile", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveProfile(HttpServletRequest request, @RequestBody UserVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		TestUser user = (TestUser) accountService.saveProfile(vo);
		vo = userService.genVo(user);
		request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, vo);
		
		ret.put("data", vo);
		ret.put("code", RespCode.SUCCESS.getCode());
		return ret;
	}
	
	@RequestMapping(value = "changePassword", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> changePassword(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		UserVo vo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);
		String oldPassword = json.getString("oldPassword");
		String password = json.getString("password");
		
		boolean success = accountService.changePasswordPers(vo.getId(), oldPassword, password);
		int code = success?RespCode.SUCCESS.getCode(): RespCode.BIZ_FAIL.getCode();
		
		ret.put("code", code);
		return ret;
	}
		
}
