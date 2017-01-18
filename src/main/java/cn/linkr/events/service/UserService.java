package cn.linkr.events.service;

import java.util.Map;

import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysVerifyCode;

public interface UserService extends BaseService {

	SysUser getByToken(String token);

	SysUser getByPhone(String token);

	SysUser loginPers(String mobile, String password, Boolean rememberMe, String platform,
			String agent, String deviceToken);
	SysUser logoutPers(SysUser user);

	SysUser registerPers(String name, String email, String phone, String password, String platform,
			String isWebView, String deviceToken);

	SysVerifyCode forgetPaswordPers(String mobile);

	SysUser resetPasswordPers(String verifyCode, String mobile,
			String password, String platform, String isWebView,
			String deviceToken);

}
