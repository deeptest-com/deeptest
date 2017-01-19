package cn.linkr.events.service;

import java.util.Map;

import cn.linkr.events.entity.SysUser;
import cn.linkr.events.entity.SysVerifyCode;
import cn.linkr.events.vo.UserVo;

public interface UserService extends BaseService {

	SysUser getByToken(String token);

	SysUser getByPhone(String token);

	SysUser loginPers(String mobile, String password, Boolean rememberMe);
	SysUser logoutPers(SysUser user);

	SysUser registerPers(String name, String email, String phone, String password);

	SysVerifyCode forgetPaswordPers(String mobile);

	SysUser resetPasswordPers(String verifyCode, String mobile,
			String password, String platform, String isWebView,
			String deviceToken);

	UserVo genVo(SysUser user);

}
