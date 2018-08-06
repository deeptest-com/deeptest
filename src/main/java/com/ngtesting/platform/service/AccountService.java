package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstUser;

public interface AccountService {
    TstUser register(TstUser user);
    TstUser login(String email, String password, Boolean rememberMe);
    Boolean logout(String email);
    Boolean changePassword(Integer userId, String oldPassword, String password);

    Boolean checkResetPassword(Integer userId, String verifyCode);

    TstUser resetPassword(Integer userId, String verifyCode, String password);

    String forgotPassword(TstUser user);
}
