package com.ngtesting.platform.action;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.config.Constant.RespCode;
import com.ngtesting.platform.config.PropertyConfig;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.entity.TestVerifyCode;
import com.ngtesting.platform.service.*;
import com.ngtesting.platform.util.AuthPassport;
import com.ngtesting.platform.vo.UserVo;
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
@RequestMapping(Constant.API_PATH_CLIENT + "account/")
public class AccountAction extends BaseAction {

	@Autowired
	AccountService accountService;
	@Autowired
	UserService userService;
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

			ret.put("profile", userVo);
			ret.put("token", user.getToken());
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "登录失败");
		}

		return ret;
	}

    @AuthPassport(validate=false)
    @RequestMapping(value = "loginWithVcode", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> loginWithVcode(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        String vcode = json.getString("vcode");
        TestUser user = accountService.resetPasswordPers(vcode, null);

        if (user != null) {
            UserVo userVo = userService.genVo(user);
            request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userVo);

            ret.put("profile", userVo);
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
            orgService.createDefaultBasicDataPers(user);

            TestVerifyCode verifyCode = accountService.genVerifyCodePers(user.getId());
            String sys = PropertyConfig.getConfig("sys.name");

            Map<String, String> map = new HashMap<String, String>();
            map.put("name", user.getName());
            map.put("vcode", verifyCode.getCode());
            map.put("url", Constant.WEB_ROOT + PropertyConfig.getConfig("url.login"));
            mailService.sendTemplateMail("[\"" + sys + "\"]注册成功", "register-success.ftl", user.getEmail(), map);
            ret.put("msg", "注册成功，请访问您的邮箱进行登录");
            ret.put("code", RespCode.SUCCESS.getCode());
        } else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "邮箱已存在");
		}

		return ret;
	}

    @AuthPassport(validate=false)
    @RequestMapping(value = "checkResetPassword", method = RequestMethod.POST)
    @ResponseBody
    public Map<String, Object> checkResetPassword(HttpServletRequest request, @RequestBody JSONObject json) {
        Map<String, Object> ret = new HashMap<String, Object>();

        String verifyCode = json.getString("vcode");

        boolean success = accountService.checkResetPassword(verifyCode);

        if (success) {
            ret.put("code", RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "重置密码链接已超时失效");
        }

        return ret;
    }

	@AuthPassport(validate=false)
	@RequestMapping(value = "forgotPassword", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> forgotPassword(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String email = json.getString("email");
		TestUser user = accountService.getByEmail(email);
		if (user == null) {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "用户不存在");
			return ret;
		}

		TestVerifyCode verifyCode = accountService.genVerifyCodePers(user.getId());
		if (verifyCode != null) {
            String sys = PropertyConfig.getConfig("sys.name");

			Map<String, String> map = new HashMap<String, String>();
			map.put("name", user.getName());
			map.put("vcode", verifyCode.getCode());
			map.put("url", Constant.WEB_ROOT + PropertyConfig.getConfig("url.reset.password"));
			mailService.sendTemplateMail("[\"" + sys + "\"]忘记密码", "forgot-password.ftl", user.getEmail(), map);

			ret.put("data", verifyCode);
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "用户不存在");
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

            ret.put("token", user.getToken());
            ret.put("code", RespCode.SUCCESS.getCode());
        } else {
            ret.put("code", RespCode.BIZ_FAIL.getCode());
            ret.put("msg", "重置密码失败");
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
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "登出失败");
		}

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
