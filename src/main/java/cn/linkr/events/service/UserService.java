package cn.linkr.events.service;

import java.util.Map;

import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysVerifyCode;

public interface UserService extends BaseService {

	SysUser getUserByToken(String token);

	SysUser getUserByPhone(String token);

	SysUser loginPers(String mobile, String password, String platform,
			String agent, String deviceToken);

	SysUser registerPers(String mobile, String password, String platform,
			String isWebView, String deviceToken);

	SysVerifyCode forgetPaswordPers(String mobile);

	SysUser resetPasswordPers(String verifyCode, String mobile,
			String password, String platform, String isWebView,
			String deviceToken);

}
