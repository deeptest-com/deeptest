package com.ngtesting.platform.service;

import com.ngtesting.platform.model.IsuField;

import java.util.List;
import java.util.Map;

public interface IssueFieldService extends BaseService {
    Map<String,List> getProps(Integer orgId);

	Map<String,Map<String,String>> getMap(Integer orgId);

	Map<String,String> getPriorityMap(Integer orgId);

	Map<String,String> getTypeMap(Integer orgId);

	Map<String,String> getStatusMap(Integer orgId);

	Map<String,String> getResolutionsMap(Integer orgId);

    List<IsuField> listOrgField(Integer orgId, Integer tabId);
}
