package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface IssueWatchService extends BaseService {
    List<TstUser> list(Integer issueId);
    List<TstUser> search(Integer issueId, Integer orgId, String keywords, List<Integer> exceptIds);

    void remove(Integer id);

    void batchSave(Integer issueId, List<Integer> userIds);
    void watch(Integer id, TstUser user, Boolean status);
}
