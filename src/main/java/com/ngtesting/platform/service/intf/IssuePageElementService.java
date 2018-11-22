package com.ngtesting.platform.service.intf;

import java.util.List;
import java.util.Map;

public interface IssuePageElementService extends BaseService {

    void saveAll(Integer orgId, Integer pageId, Integer tabId, List<Map> jsonArr);

    void updateProp(String id, String prop, String val, Integer orgId);

//    void add(IsuPageElement element);
//    boolean remove(Integer id, Integer orgId);
}
