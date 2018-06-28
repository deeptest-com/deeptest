package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface UserDao {

    Integer insert(TstUser record);

    List<TstUser> selectUsers();
    TstUser get(Integer userId);

}
