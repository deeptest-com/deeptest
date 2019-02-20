package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuHistory;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.utils.MsgUtil;

import java.util.List;

public interface IssueHistoryService extends BaseService {
    List<IsuHistory> query(Integer issueId);
    void saveHistory(TstUser user, MsgUtil.MsgAction act, Integer issueId, String field);
}
