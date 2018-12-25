package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface IssueWatchService extends BaseService {
    List<TstUser> list(Integer issueId);
    List<TstUser> search(Integer issueId, Integer orgId, String keywords, List<Integer> exceptIds, TstUser user);

    Boolean remove(Integer id, Integer issueId, TstUser user);

    Boolean batchSave(Integer issueId, List<Integer> userIds, TstUser user);
    Boolean watch(Integer id, TstUser user, Boolean status);
}
