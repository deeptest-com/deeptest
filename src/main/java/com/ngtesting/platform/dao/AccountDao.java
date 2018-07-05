package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;
import org.apache.ibatis.annotations.Param;

import java.util.Date;
import java.util.Map;

public interface AccountDao {

    Integer register(TstUser record);
    void login(@Param("id") Integer id, @Param("token") String token, @Param("lastLoginTime") Date lastLoginTime);

    void initUser(Integer userId);

    void genVerifyCode(Map<String, Object> map);

}
