package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.AccountDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AccountService;
import com.ngtesting.platform.service.MailService;
import com.ngtesting.platform.service.OrgService;
import com.ngtesting.platform.service.PropService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Isolation;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

@Service(value = "accountService")
public class AccountServiceImpl implements AccountService {
    @Autowired
    private AccountDao accountDao;
    @Autowired
    private UserDao userDao;
    @Autowired
    private OrgService orgService;
    @Autowired
    private PropService propService;
    @Autowired
    private MailService mailService;

    @Transactional(propagation = Propagation.REQUIRED,isolation = Isolation.DEFAULT,timeout=36000,rollbackFor=Exception.class)
    @Override
    public TstUser register(TstUser user) {
        TstUser existUser = userDao.getByEmail(user.getEmail());
        if (existUser != null) {
            return null;
        }

        accountDao.register(user);
        TstUser po = userDao.get(user.getId());

        if (po != null) {
            accountDao.initUser(user.getId());

            String verifyCode = genVerifyCode(po);
            String sys = propService.getSysName();

            Map<String, String> map = new HashMap<String, String>();
            map.put("name", user.getNickname());
            map.put("vcode", verifyCode);

            String url = propService.getUrlLogin();
            if (!url.startsWith("http")) {
                url = Constant.WEB_ROOT + url;
            }
            map.put("url", url);
            mailService.sendTemplateMail("[\"" + sys + "\"]注册成功", "register-success.ftl",
                    user.getEmail(), map);
        }

        return po;
    }

    @Override
    public TstUser login(String mobile, String password, Boolean rememberMe) {
        return null;
    }

    @Override
    public TstUser logout(String email) {
        return null;
    }

    @Override
    public boolean changePassword(Integer userId, String oldPassword, String password) {
        return false;
    }

    @Override
    public boolean checkResetPassword(String verifyCode) {
        return false;
    }

    @Override
    public TstUser resetPasswordPers(String verifyCode, String password) {
        return null;
    }

    @Override
    public String genVerifyCode(TstUser user) {
        String code = UUID.randomUUID().toString().replaceAll("-", "");
        Map<String, Object> map = new HashMap();
        map.put("userId", user.getId().toString());
        map.put("code", code);

        Date now = new Date();
        map.put("createTime", now);
        map.put("expireTime", new Date(now.getTime() + 10 * 60 * 1000));

        accountDao.genVerifyCode(map);

        return code;
    }

    @Override
    public TstUser getByToken(String token) {
        return null;
    }

    @Override
    public TstUser getByPhone(String token) {
        return null;
    }

    @Override
    public TstUser getByEmail(String email) {
        return null;
    }

}
