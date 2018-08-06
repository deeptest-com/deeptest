package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstVerifyCode;

import java.util.Map;

public interface AccountVerifyCodeDao {

    void genVerifyCode(Map<String, Object> map);

    TstVerifyCode findAndDisableCode(Integer userId, String verifyCode);
}
