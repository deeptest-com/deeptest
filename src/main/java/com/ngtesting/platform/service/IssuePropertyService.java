package com.ngtesting.platform.service;

import java.util.Map;

public interface IssuePropertyService extends BaseService {


	Map<String,Map<String,String>> getMap(Long orgId);

	Map<String,String> getPriorityMap(Long orgId);

	Map<String,String> getTypeMap(Long orgId);

	Map<String,String> getExeStatusMap(Long orgId);
}
