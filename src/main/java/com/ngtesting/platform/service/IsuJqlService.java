package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuIssue;

import java.util.List;

public interface IsuJqlService extends BaseService {
    List<IsuIssue> query(String jql, Integer orgId, Integer projectId);

    String buildDefaultJql(Integer orgId, Integer projectId);

}
