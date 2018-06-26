package com.ngtesting.platform.service.intf;

import com.github.pagehelper.PageInfo;
import com.ngtesting.platform.model.TstUser;

public interface UserService {

    int addUser(TstUser user);

    PageInfo<TstUser> findAllUser(int pageNum, int pageSize);

}
