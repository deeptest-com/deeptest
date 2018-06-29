package com.ngtesting.platform.service;

import com.github.pagehelper.PageInfo;
import com.ngtesting.platform.model.TstUser;

public interface UserService {
    PageInfo<TstUser> query(int pageNum, int pageSize);

    TstUser get(Integer id);

    TstUser getByToken(String token);

    void update(TstUser record);

}
