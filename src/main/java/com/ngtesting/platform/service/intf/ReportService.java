package com.ngtesting.platform.service.intf;

import java.util.List;
import java.util.Map;

public interface ReportService extends BaseService {

    Map<String, List<Object>> countByStatus(List<Map> ls);

    Map<String, Object> countByUser(List<Map> ls);

    List<Map<Object, Object>> orderByExeStatus(Map map);
}
