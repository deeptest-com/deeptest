package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

public interface IssueOptService extends BaseService {
    void statusTran(Integer id, Integer dictStatusId, Integer projectId);
    void assign(Integer id, TstUser user, String comments);
}
