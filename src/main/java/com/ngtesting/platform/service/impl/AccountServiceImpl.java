package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.AccountDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.model.TstVerifyCode;
import com.ngtesting.platform.service.intf.AccountService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service(value = "accountService")
public class AccountServiceImpl implements AccountService {
    @Autowired
    private AccountDao accountDao;
    @Autowired
    private UserDao userDao;

    @Override
    public TstUser register(TstUser user) {
        Integer count = accountDao.register(user);
        TstUser po = userDao.get(32);
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
    public boolean changePassword(Long userId, String oldPassword, String password) {
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
    public TstVerifyCode genVerifyCode(Long userId) {
        return null;
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
