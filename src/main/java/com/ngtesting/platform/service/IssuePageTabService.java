package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPageTab;

public interface IssuePageTabService extends BaseService {

    void add(IsuPageTab tab);

    IsuPageTab get(Integer tabId, Integer orgId);

    boolean remove(Integer id, Integer orgId);
}
