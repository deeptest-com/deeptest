package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuPageElement;

public interface IssuePageElementService extends BaseService {
    void add(IsuPageElement element);

    boolean remove(Integer id, Integer orgId);
}
