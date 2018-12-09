package com.ngtesting.platform.service.intf;

import java.util.Map;

public interface CasePropertyService extends BaseService {


	Map<String,Map<String,String>> getMap(Integer orgId);

	Map<Integer,String> getPriorityMap(Integer orgId);

	Map<Integer,String> getTypeMap(Integer orgId);

//	Map<Integer,String> getExeStatusMap(Integer orgId);
}
