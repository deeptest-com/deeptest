package cn.linkr.testspace.service;

import java.util.List;
import java.util.Map;

import cn.linkr.testspace.entity.SysUser;
import cn.linkr.testspace.entity.SysVerifyCode;
import cn.linkr.testspace.vo.Page;
import cn.linkr.testspace.vo.UserVo;

public interface AccountService extends BaseService {


	SysUser loginPers(String mobile, String password, Boolean rememberMe);
	SysUser logoutPers(String email);

	SysUser registerPers(String name, String email, String phone, String password);
	
	SysUser saveProfile(UserVo vo);
	boolean changePasswordPers(Long userId, String oldPassword, String password);
	SysVerifyCode forgotPasswordPers(Long userId);
	SysUser resetPasswordPers(String verifyCode, String password);

	SysUser getByToken(String token);
	SysUser getByPhone(String token);
	SysUser getByEmail(String email);

}
