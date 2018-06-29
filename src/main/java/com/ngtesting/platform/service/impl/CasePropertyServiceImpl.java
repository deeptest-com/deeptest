package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.service.inf.CasePropertyService;
import org.springframework.stereotype.Service;

import java.util.LinkedHashMap;
import java.util.Map;

@Service
public class CasePropertyServiceImpl extends BaseServiceImpl implements CasePropertyService {
	@Override
	public Map<String,Map<String,String>> getMap(Integer orgId) {
		Map<String,String> priorityMap = getPriorityMap(orgId);
		Map<String,String> typeMap = getTypeMap(orgId);
		Map<String,String> exeStatusMap = getExeStatusMap(orgId);

		Map map = new LinkedHashMap();
		map.put("priority", priorityMap);
		map.put("type", typeMap);
		map.put("status", exeStatusMap);

		return map;
	}

	@Override
	public Map<String,String> getPriorityMap(Integer orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestCasePriority.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("displayOrder"));
//		List<TestCasePriority> ls = findAllByCriteria(dc);
//
//		Map<String,String> map = new LinkedHashMap();
//		for (TestCasePriority item : ls) {
//			map.put(item.getCode(), item.getName());
//		}
//
//		return map;

		return null;
	}

	@Override
	public Map<String,String> getTypeMap(Integer orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestCaseType.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("displayOrder"));
//		List<TestCaseType> ls = findAllByCriteria(dc);
//
//		Map<String,String> map = new LinkedHashMap();
//		for (TestCaseType item : ls) {
//			map.put(item.getCode(), item.getName());
//		}
//
//		return map;

		return null;
	}

	@Override
	public Map<String,String> getExeStatusMap(Integer orgId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TestCaseExeStatus.class);
//
//		dc.add(Restrictions.eq("orgId", orgId));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("displayOrder"));
//		List<TestCaseExeStatus> ls = findAllByCriteria(dc);
//
//		Map<String,String> map = new LinkedHashMap();
//		for (TestCaseExeStatus item : ls) {
//			map.put(item.getCode(), item.getName());
//		}
//
//		return map;

		return null;
	}
}
