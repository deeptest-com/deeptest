package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.*;
import com.ngtesting.platform.model.IsuPriority;
import com.ngtesting.platform.model.IsuResolution;
import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.service.IssueFieldService;
import com.ngtesting.platform.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class IssueFieldServiceImpl extends BaseServiceImpl implements IssueFieldService {
    @Autowired
    UserService userService;

	@Autowired
    IssueFieldDao fieldDao;

	@Autowired
	IssuePriorityDao issuePriorityDao;

	@Autowired
	IssueTypeDao issueTypeDao;

	@Autowired
	IssueStatusDao issueStatusDao;

	@Autowired
	IssueResolutionDao issueResolutionDao;

    @Override
    public Map<String,List> getProps(Integer orgId) {

        List<IsuType> types = issueTypeDao.list(orgId);
        List<IsuStatus> statuses = issueStatusDao.list(orgId);
        List<IsuPriority> priorities = issuePriorityDao.list(orgId);
        List<IsuResolution> resolutions = issueResolutionDao.list(orgId);

        Map map = new LinkedHashMap();
        map.put("types", types);
        map.put("statuses", statuses);
        map.put("priorities", priorities);
        map.put("resolutions", resolutions);

        return map;
    }

	@Override
	public Map<String,Map<String,String>> getMap(Integer orgId) {
		Map<String,String> typeMap = getTypeMap(orgId);
		Map<String,String> priorityMap = getPriorityMap(orgId);
		Map<String,String> statusMap = getStatusMap(orgId);
        Map<String,String> resolutionMap = getResolutionsMap(orgId);

		Map map = new LinkedHashMap();
		map.put("type", typeMap);
		map.put("priority", priorityMap);
		map.put("status", statusMap);
        map.put("resolution", resolutionMap);

		return map;
	}

	@Override
	public Map<String,String> getPriorityMap(Integer orgId) {
		List<IsuPriority> ls = issuePriorityDao.list(orgId);

		Map<String,String> map = new LinkedHashMap();
		for (IsuPriority item : ls) {
			map.put(item.getValue(), item.getLabel());
		}

		return map;
	}

	@Override
	public Map<String,String> getTypeMap(Integer orgId) {
		List<IsuType> ls = issueTypeDao.list(orgId);
		Map<String,String> map = new LinkedHashMap();
		for (IsuType item : ls) {
			map.put(item.getValue(), item.getLabel());
		}

		return map;
	}

	@Override
	public Map<String,String> getStatusMap(Integer orgId) {
		List<IsuStatus> ls = issueStatusDao.list(orgId);
		Map<String,String> map = new LinkedHashMap();
		for (IsuStatus item : ls) {
			map.put(item.getValue(), item.getLabel());
		}

		return map;
	}

	@Override
	public Map<String, String> getResolutionsMap(Integer orgId) {
		List<IsuResolution> ls = issueResolutionDao.list(orgId);
		Map<String,String> map = new LinkedHashMap();
		for (IsuResolution item : ls) {
			map.put(item.getValue(), item.getLabel());
		}

		return map;
	}

}
