package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;

public interface AccountService {
    TstUser register(TstUser user);
    TstUser login(String email, String password, Boolean rememberMe);
    TstUser logout(String email);
    boolean changePassword(Integer userId, String oldPassword, String password);

    boolean checkResetPassword(String verifyCode);
    TstUser resetPasswordPers(String verifyCode, String password);

    String genVerifyCode(Integer userId);

}
