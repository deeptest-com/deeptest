package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.AccountVerifyCodeDao;
import com.ngtesting.platform.service.intf.AccountVerifyCodeService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

@Service
public class AccountVerifyCodeServiceImpl implements AccountVerifyCodeService {
    @Autowired
    private AccountVerifyCodeDao verifyCodeDao;

    @Override
    public String genVerifyCode(Integer userId) {
        String code = UUID.randomUUID().toString().replaceAll("-", "");
        Map<String, Object> map = new HashMap();
        map.put("userId", userId);
        map.put("code", code);

        Date now = new Date();
        map.put("createTime", now);
        map.put("expireTime", new Date(now.getTime() + 10 * 60 * 1000));

        verifyCodeDao.genVerifyCode(map);

        return code;
    }

}
