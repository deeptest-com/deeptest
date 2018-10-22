package com.ngtesting.platform.service;

import java.util.Map;

public interface IssuePropertyService extends BaseService {

	Map<String,Map<String,String>> getMap(Integer orgId);

	Map<String,String> getPriorityMap(Integer orgId);

	Map<String,String> getTypeMap(Integer orgId);

	Map<String,String> getStatusMap(Integer orgId);

	Map<String,String> getResolutionsMap(Integer orgId);
}
