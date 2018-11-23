package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.IsuField;

public interface IssueFieldService extends BaseService {
//  List<IsuField> listOrgField(Integer orgId, Integer tabId);

  IsuField getField(String key, Integer orgId);

//    Map<String,List> getProps(Integer orgId);
//
//	Map<String,Map<String,String>> getMap(Integer orgId);
//
//	Map<String,String> getPriorityMap(Integer orgId);
//
//	Map<String,String> getTypeMap(Integer orgId);
//
//	Map<String,String> getStatusMap(Integer orgId);
//
//	Map<String,String> getResolutionsMap(Integer orgId);
//
//    List<IsuField> listOrgField(Integer orgId, Integer tabId);
}
