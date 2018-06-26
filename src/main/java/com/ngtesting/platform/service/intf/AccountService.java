package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVerifyCode;

public interface AccountService {
    TstUser register(TstUser user);
    TstUser login(String mobile, String password, Boolean rememberMe);
    TstUser logout(String email);
    boolean changePassword(Long userId, String oldPassword, String password);

    boolean checkResetPassword(String verifyCode);
    TstUser resetPasswordPers(String verifyCode, String password);

    TstVerifyCode genVerifyCode(Long userId);
    TstUser getByToken(String token);
    TstUser getByPhone(String token);
    TstUser getByEmail(String email);

}
