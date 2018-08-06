package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstVerifyCode;
import org.apache.ibatis.annotations.Param;

import java.util.Map;

public interface AccountVerifyCodeDao {

    void genVerifyCode(Map<String, Object> map);
    TstVerifyCode getByCode(@Param("code") String code);

    Integer disableCode(@Param("id") Integer id);
}
