package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

public interface IssueOptService extends BaseService {
    void statusTran(Integer id, Integer dictStatusId, String dictStatusName, TstUser user);
    void assign(Integer id, Integer userId, String content, TstUser user);
}
