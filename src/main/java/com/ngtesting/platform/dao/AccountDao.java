package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVerifyCode;
import org.apache.ibatis.annotations.Param;

import java.util.Date;

public interface AccountDao {

    Integer register(TstUser record);
    void login(@Param("id") Integer id, @Param("token") String token, @Param("lastLoginTime") Date lastLoginTime);
    Integer logout(@Param("email") String email);

    void initUser(Integer userId);

    Integer changePassword(@Param("userId") Integer userId,
                        @Param("oldPassword") String oldPassword,
                        @Param("password") String password);

    TstVerifyCode checkResetPassword(@Param("userId") Integer userId,
                                     @Param("verifyCode") String verifyCode);

    void resetPassword(@Param("userId") Integer userId,
                       @Param("password") String password,
                       @Param("newToken") String newToken);
}
