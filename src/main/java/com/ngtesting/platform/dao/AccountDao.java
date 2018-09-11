package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;
import org.apache.ibatis.annotations.Param;

import java.util.Date;

public interface AccountDao {

    Integer register(TstUser record);
    void loginWithVerifyCode(TstUser user);
    void login(@Param("id") Integer id, @Param("token") String token, @Param("lastLoginTime") Date lastLoginTime);
    Integer logout(@Param("email") String email);

    void initUser(@Param("userId") Integer userId, @Param("orgName") String orgName);

    Integer changePassword(@Param("userId") Integer userId,
                        @Param("oldPassword") String oldPassword,
                        @Param("password") String password);

    void resetPassword(TstUser user);
}
