package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.inf.PropService;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.stereotype.Service;

@Service
public class PropServiceImpl implements PropService {

    @Value("${sys.name}")
    public String sysName;

    @Value("${url.login}")
    public String urlLogin;

    @Override
    public String getSysName() {
        return sysName;
    }

    @Override
    public String getUrlLogin() {
        return urlLogin;
    }
}
