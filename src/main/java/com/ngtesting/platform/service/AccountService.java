package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.entity.TestVerifyCode;
import com.ngtesting.platform.vo.UserVo;

public interface AccountService extends BaseService {


	TestUser loginPers(String mobile, String password, Boolean rememberMe);
	TestUser logoutPers(String email);

	TestUser registerPers(String name, String email, String phone, String password);
	
	TestUser saveProfile(UserVo vo);
	TestUser saveInfo(JSONObject json);
	boolean changePasswordPers(Long userId, String oldPassword, String password);
	TestVerifyCode genVerifyCodePers(Long userId);

	boolean checkResetPassword(String verifyCode);
	TestUser resetPasswordPers(String verifyCode, String password);

	TestUser getByToken(String token);

    TestUser getByPhone(String token);
	TestUser getByEmail(String email);
    TestUser setLeftSizePers(Long userId, Integer left);
}
