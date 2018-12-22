package com.ngtesting.platform.service.intf;

import java.util.Map;

public interface CasePropertyService extends BaseService {


	Map<String,Map<String,String>> getMap(Integer orgId);

	Map<String,String> getPriorityMap(Integer orgId);

	Map<String,String> getTypeMap(Integer orgId);

//	Map<String,String> getExeStatusMap(Integer orgId);
}
