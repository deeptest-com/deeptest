package cn.linkr.events.action.admin;

import java.util.HashMap;
import java.util.Map;

import javax.servlet.http.HttpServletRequest;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.ResponseBody;

import cn.linkr.events.action.client.BaseAction;
import cn.linkr.events.constants.Constant;
import cn.linkr.events.constants.Constant.RespCode;
import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysVerifyCode;
import cn.linkr.events.service.RegisterService;
import cn.linkr.events.service.SessionService;
import cn.linkr.events.service.UserService;
import cn.linkr.events.util.AuthPassport;
import cn.linkr.events.util.BeanUtilEx;
import cn.linkr.events.vo.UserVo;

import com.alibaba.fastjson.JSONObject;


@Controller
@RequestMapping(Constant.API_PATH_ADMIN + "user/")
public class UserAdmin extends BaseAction {
	@Autowired
	UserService userService;
	
	@Autowired
	RegisterService registerService;
	
	@Autowired
	SessionService sessionService;
	
	@AuthPassport(validate=false)
	@RequestMapping(value = "login", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> login(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String email = json.getString("email");
		String password = json.getString("password");
		boolean rememberMe = json.getBoolean("rememberMe") != null? json.getBoolean("rememberMe"): false;
		
		SysUser user = userService.loginPers(email, password, rememberMe);
		request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userService.genVo(user));

		if (user != null) {
			ret.put("token", user.getToken());

			UserVo vo = userService.genVo(user); 
			ret.put("data", vo);
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "登录失败");
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
		SysUser user = userService.logoutPers(vo.getEmail());
		
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

	@AuthPassport(validate=false)
	@RequestMapping(value = "register", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> register(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String name = json.getString("name");
		String phone = json.getString("phone");
		String email = json.getString("email");
		String password = json.getString("password");

		SysUser user = userService.registerPers(name, email, phone, password);
		request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, userService.genVo(user));

		if (user != null) {
			ret.put("token", user.getToken());

			UserVo vo = userService.genVo(user); 
			ret.put("data", vo);
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "邮箱已存在");
		}

		return ret;
	}

	@AuthPassport(validate=false)
	@RequestMapping(value = "forgotPassword", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> forgotPassword(HttpServletRequest request, @RequestBody JSONObject json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		String phone = json.getString("phone");

		SysVerifyCode verifyCode = userService.forgetPaswordPers(phone);

		if (verifyCode != null) {
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

		String verifyCode = json.getString("verifyCode");
		String phone = json.getString("phone");
		String password = json.getString("password");
		String platform = json.getString("platform");
		String isWebView = json.getString("isWebView");
		String deviceToken = json.getString("deviceToken");

		SysUser user = userService.resetPasswordPers(verifyCode, phone, password, platform, isWebView, deviceToken);

		if (user != null) {
			ret.put("token", user.getToken());

			UserVo vo = new UserVo();
			BeanUtilEx.copyProperties(vo, user);
			ret.put("data", vo);
			ret.put("code", RespCode.SUCCESS.getCode());
		} else {
			ret.put("code", RespCode.BIZ_FAIL.getCode());
			ret.put("msg", "重置密码失败");
		}

		return ret;
	}

	@RequestMapping(value = "getProfile", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> getProfile(HttpServletRequest request, @RequestBody Map<String, String> json) {
		Map<String, Object> ret = new HashMap<String, Object>();

		UserVo vo = (UserVo) request.getSession().getAttribute(Constant.HTTP_SESSION_USER_KEY);

		ret.put("data", vo);
		ret.put("code", RespCode.SUCCESS.getCode());

		return ret;
	}

	@RequestMapping(value = "saveProfile", method = RequestMethod.POST)
	@ResponseBody
	public Map<String, Object> saveProfile(HttpServletRequest request, @RequestBody UserVo vo) {
		Map<String, Object> ret = new HashMap<String, Object>();
		
		SysUser user = (SysUser) userService.saveProfile(vo);
		vo = userService.genVo(user);
		request.getSession().setAttribute(Constant.HTTP_SESSION_USER_KEY, vo);
		
		ret.put("data", vo);
		ret.put("code", RespCode.SUCCESS.getCode());
		return ret;
	}

}
