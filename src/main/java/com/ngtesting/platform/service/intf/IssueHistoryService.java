package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.IsuHistory;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface IssueHistoryService extends BaseService {
    List<IsuHistory> query(Integer issueId);
    void saveHistory(TstUser user, Constant.EntityAct act, Integer issueId, String field);
}
