package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

public interface AccountService {
    TstUser register(TstUser user);
    Boolean logout(String email);
    Boolean changePassword(Integer userId, String oldPassword, String password);

    String forgotPassword(TstUser user);
    Boolean beforResetPassword(String verifyCode);
    TstUser resetPassword(String verifyCode, String password);

    TstUser loginWithVerifyCode(String vcode);
}
