package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

import java.util.List;
import java.util.Map;

public interface IssueWatchService extends BaseService {
    List<Map> list(Integer issueId);
    List<TstUser> search(Integer issueId, Integer orgId, String keywords, List<Integer> exceptIds, TstUser user);

    Boolean remove(Integer id, Integer issueId, TstUser user);

    Boolean batchWatch(Integer issueId, List<Integer> userIds, TstUser user);
    Boolean watch(Integer id, TstUser user, Boolean status);
}
