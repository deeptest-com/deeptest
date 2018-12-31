package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.*;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.IssueFieldService;
import com.ngtesting.platform.service.intf.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssueFieldServiceImpl extends BaseServiceImpl implements IssueFieldService {

	@Autowired
    IssueFieldDao fieldDao;

//    @Autowired
//    CustomFieldDao customFieldDao;
//
//	@Autowired
//	IssuePriorityDao issuePriorityDao;
//
//	@Autowired
//	IssueTypeDao issueTypeDao;
//
//	@Autowired
//	IssueStatusDao issueStatusDao;
//
//	@Autowired
//	IssueResolutionDao issueResolutionDao;
//
//    IssueDynamicFormDao dynamicFormDao;

    @Override
    public IsuField getField(String key, Integer orgId) {
        String[] arr = key.split("-");
        String src = arr[0];
        Integer id = Integer.valueOf(arr[1]);

        IsuField field = null;
        if ("sys".equals(src)) {
            field = fieldDao.getSysField(id, orgId);
        } else if ("cust".equals(src)) {
            field = fieldDao.getCustField(id, orgId);
        }

        return field;
    }

//    @Override
//	public List<IsuField> listOrgField(Integer orgId, Integer tabId) {
//    	return dynamicFormDao.listOrgField(orgId, tabId);
//	}

//    @Override
//    public Map<String,List> getProps(Integer orgId) {
//
//        List<IsuType> types = issueTypeDao.listByIssueId(orgId);
//        List<IsuStatus> statuses = issueStatusDao.listByIssueId(orgId);
//        List<IsuPriority> priorities = issuePriorityDao.listByIssueId(orgId);
//        List<IsuResolution> resolutions = issueResolutionDao.listByIssueId(orgId);
//
//        Map map = new LinkedHashMap();
//        map.put("types", types);
//        map.put("statuses", statuses);
//        map.put("priorities", priorities);
//        map.put("resolutions", resolutions);
//
//        return map;
//    }
//
//	@Override
//	public Map<String,Map<String,String>> getMap(Integer orgId) {
//		Map<String,String> typeMap = getTypeMap(orgId);
//		Map<String,String> priorityMap = getPriorityMap(orgId);
//		Map<String,String> statusMap = getStatusMap(orgId);
//        Map<String,String> resolutionMap = getResolutionsMap(orgId);
//
//		Map map = new LinkedHashMap();
//		map.put("type", typeMap);
//		map.put("priority", priorityMap);
//		map.put("status", statusMap);
//        map.put("resolution", resolutionMap);
//
//		return map;
//	}
//
//	@Override
//	public Map<String,String> getPriorityMap(Integer orgId) {
//		List<IsuPriority> ls = issuePriorityDao.listByIssueId(orgId);
//
//		Map<String,String> map = new LinkedHashMap();
//		for (IsuPriority item : ls) {
//			map.put(item.getValue(), item.getLabel());
//		}
//
//		return map;
//	}
//
//	@Override
//	public Map<String,String> getTypeMap(Integer orgId) {
//		List<IsuType> ls = issueTypeDao.listByIssueId(orgId);
//		Map<String,String> map = new LinkedHashMap();
//		for (IsuType item : ls) {
//			map.put(item.getValue(), item.getLabel());
//		}
//
//		return map;
//	}
//
//	@Override
//	public Map<String,String> getStatusMap(Integer orgId) {
//		List<IsuStatus> ls = issueStatusDao.listByIssueId(orgId);
//		Map<String,String> map = new LinkedHashMap();
//		for (IsuStatus item : ls) {
//			map.put(item.getValue(), item.getLabel());
//		}
//
//		return map;
//	}
//
//	@Override
//	public Map<String, String> getResolutionsMap(Integer orgId) {
//		List<IsuResolution> ls = issueResolutionDao.listByIssueId(orgId);
//		Map<String,String> map = new LinkedHashMap();
//		for (IsuResolution item : ls) {
//			map.put(item.getValue(), item.getLabel());
//		}
//
//		return map;
//	}
//

}
