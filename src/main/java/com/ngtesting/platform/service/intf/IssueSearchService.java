package com.ngtesting.platform.service.intf;

import java.util.List;
import java.util.Map;

public interface IssueSearchService extends BaseService {
    List<Map> idAndTitleSearch(String text, List<Integer> exceptIds, Integer orgId);
}
