package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONArray;

public interface IssuePageElementService extends BaseService {

    void saveAll(Integer orgId, Integer pageId, Integer tabId, JSONArray jsonArr);

    void updateProp(String id, String prop, String val, Integer orgId);

//    void add(IsuPageElement element);
//    boolean remove(Integer id, Integer orgId);
}
