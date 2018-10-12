package com.ngtesting.platform.service;

import java.util.Map;

public interface IsuJqlService extends BaseService {
    Map<String, Object> query(String jql, Integer orgId, Integer projectId);

    String buildDefaultJql(Integer orgId, Integer projectId);

}
