package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;

import java.util.Map;

public interface AccountDao {

    Integer register(TstUser record);

    void initUser(Integer userId);

    void genVerifyCode(Map<String, Object> map);
}
