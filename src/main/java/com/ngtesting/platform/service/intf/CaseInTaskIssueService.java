package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstUser;

public interface CaseInTaskIssueService extends BaseService {

    Boolean save(Integer caseInTaskId, Integer issueId, TstUser user);
    Boolean remove(Integer caseInTaskId, Integer id, TstUser user);

}
