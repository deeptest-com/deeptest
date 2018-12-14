package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;

public interface IssueHistoryService extends BaseService {
    void saveHistory(TstUser user, Constant.EntityAct act, IsuIssue issue, String field);
}
