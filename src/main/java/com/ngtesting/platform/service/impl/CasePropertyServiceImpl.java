package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CaseExeStatusDao;
import com.ngtesting.platform.dao.CasePriorityDao;
import com.ngtesting.platform.dao.CaseTypeDao;
import com.ngtesting.platform.model.TstCaseExeStatus;
import com.ngtesting.platform.model.TstCasePriority;
import com.ngtesting.platform.model.TstCaseType;
import com.ngtesting.platform.service.CasePropertyService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

@Service
public class CasePropertyServiceImpl extends BaseServiceImpl implements CasePropertyService {

	@Autowired
    CasePriorityDao casePriorityDao;

	@Autowired
    CaseTypeDao caseTypeDao;

	@Autowired
    CaseExeStatusDao caseExeStatusDao;

	@Override
	public Map<String,Map<String,String>> getMap(Integer orgId) {
		Map<String,String> typeMap = getTypeMap(orgId);
		Map<String,String> priorityMap = getPriorityMap(orgId);
		Map<String,String> exeStatusMap = getExeStatusMap(orgId);

		Map map = new LinkedHashMap();
		map.put("type", typeMap);
		map.put("priority", priorityMap);
		map.put("status", exeStatusMap);

		return map;
	}

	@Override
	public Map<String,String> getPriorityMap(Integer orgId) {
        List<TstCasePriority> ls = casePriorityDao.list(orgId);

        Map<String,String> map = new LinkedHashMap();
		for (TstCasePriority item : ls) {
			map.put(item.getValue(), item.getLabel());
		}

		return map;
	}

	@Override
	public Map<String,String> getTypeMap(Integer orgId) {
        List<TstCaseType> ls = caseTypeDao.list(orgId);
        Map<String,String> map = new LinkedHashMap();
		for (TstCaseType item : ls) {
			map.put(item.getValue(), item.getLabel());
		}

		return map;
	}

	@Override
	public Map<String,String> getExeStatusMap(Integer orgId) {
        List<TstCaseExeStatus> ls = caseExeStatusDao.listExeStatus(orgId);
        Map<String,String> map = new LinkedHashMap();
		for (TstCaseExeStatus item : ls) {
			map.put(item.getValue(), item.getLabel());
		}

		return map;
	}
}
