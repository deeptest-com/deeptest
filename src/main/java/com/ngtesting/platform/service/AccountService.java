package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.entity.TestVerifyCode;
import com.ngtesting.platform.vo.UserVo;

public interface AccountService extends BaseService {


	TestUser loginPers(String mobile, String password, Boolean rememberMe);
	TestUser logoutPers(String email);

	TestUser registerPers(String name, String email, String phone, String password);
	
	TestUser saveProfile(UserVo vo);
	boolean changePasswordPers(Long userId, String oldPassword, String password);
	TestVerifyCode forgotPasswordPers(Long userId);
	TestUser resetPasswordPers(String verifyCode, String password);

	TestUser getByToken(String token);
	TestUser getByPhone(String token);
	TestUser getByEmail(String email);

}
