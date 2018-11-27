package com.ngtesting.platform.service.intf;

import java.util.List;
import java.util.Map;

public interface IssuePageElementService extends BaseService {

    void saveAll(Integer orgId, Integer pageId, List<Map> jsonArr);

    void updateProp(String id, String prop, String val, Integer orgId);
}
