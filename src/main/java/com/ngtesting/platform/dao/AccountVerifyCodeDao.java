package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUserVerifyCode;
import org.apache.ibatis.annotations.Param;

import java.util.Map;

public interface AccountVerifyCodeDao {

    void genVerifyCode(Map<String, Object> map);
    TstUserVerifyCode getByCode(@Param("code") String code);

    Integer disableCode(@Param("id") Integer id);
}
