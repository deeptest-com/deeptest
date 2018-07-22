package com.ngtesting.platform.service;

import com.github.pagehelper.PageInfo;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface UserService {
    PageInfo<TstUser> query(int pageNum, int pageSize);

    TstUser get(Integer id);

    TstUser getByToken(String token);

    void update(TstUser record);

    void setDefaultOrg(TstUser user, Integer orgId);

    void setDefaultPrj(TstUser user, Integer prjId);

    List<TstUser> search(Integer orgId, String keywords, String exceptIds);
}
