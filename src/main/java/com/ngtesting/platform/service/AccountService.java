package com.ngtesting.platform.service;

import java.util.List;
import java.util.Map;

import com.ngtesting.platform.entity.SysUser;
import com.ngtesting.platform.entity.SysVerifyCode;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.UserVo;

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
