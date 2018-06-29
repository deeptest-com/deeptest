package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface UserDao {
    List<TstUser> query();
    TstUser get(Integer userId);
    TstUser getByEmail(String nickname);
    TstUser getByToken(String token);
    void update(TstUser record);
}
