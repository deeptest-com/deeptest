package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestCaseStep;
import com.ngtesting.platform.service.CaseService;
import com.ngtesting.platform.service.CustomFieldService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseStepVo;
import com.ngtesting.platform.vo.TestCaseVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class CaseServiceImpl extends BaseServiceImpl implements CaseService {

	@Autowired
	CustomFieldService customFieldService;

	@Override
	public List<TestCase> query(Long suiteId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestCase.class);
        
        if (suiteId != null) {
        	dc.add(Restrictions.eq("suiteId", suiteId));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("ordr"));
        dc.addOrder(Order.asc("id"));
        List<TestCase> ls = findAllByCriteria(dc);
        
        return ls;
	}

	@Override
	public TestCaseVo getById(Long caseId) {
		TestCase po = (TestCase) get(TestCase.class, caseId);
		TestCaseVo vo = genVo(po);

		return vo;
	}

	@Override
	public List<TestCaseVo> genVos(List<TestCase> pos) {
        List<TestCaseVo> vos = new LinkedList<TestCaseVo>();

        for (TestCase po: pos) {
        	TestCaseVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TestCaseVo genVo(TestCase po) {
		TestCaseVo vo = new TestCaseVo();

		BeanUtilEx.copyProperties(vo, po);

        vo.setSteps(new LinkedList<TestCaseStepVo>());

		List<TestCaseStep> steps = po.getSteps();
		for (TestCaseStep step : steps) {

			TestCaseStepVo stepVo = new TestCaseStepVo(
					step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());

			vo.getSteps().add(stepVo);
		}

//		List<TestCaseProp> props = po.getProps();
//		for (TestCaseProp propPo : props) {
//
//			TestCasePropVo propVo = new TestCasePropVo(propPo.getId(), propPo.getCode(),
//                    propPo.getLabel(), propPo.getValue(), propPo.getFieldId());
//
////			CustomFieldVo fieldVo = new CustomFieldVo();
////			BeanUtilEx.copyProperties(fieldVo, propPo.getField());
////			propVo.setField(fieldVo);
//
//			vo.getProps().add(propVo);
//		}

		return vo;
	}

	@Override
	public TestCase create(Long id, String title, String type, Long pid, Long userId) {
		TestCase parent = (TestCase) get(TestCase.class, pid);
		
		TestCase testCase = new TestCase();
		testCase.setTitle(title);
		testCase.setSuiteId(pid);
		testCase.setProjectId(parent.getProjectId());
		testCase.setUserId(userId);
		
		testCase.setOrdr(getChildMaxOrderNumb(parent) + 1);
		
		saveOrUpdate(testCase);
		
		return testCase;
	}

	@Override
	public TestCase rename(Long id, String value, Long id2) {
		// TODO Auto-generated method stub
		return null;
	}
	
	@Override
	public TestCase move(Long id, Long pid, Long prePid, Long id2) {
		
		
		return null;
	}

	@Override
	public TestCase save(JSONObject json) {
		TestCase testCase = (TestCase) get(TestCase.class, json.getLong("id"));

//		long i = 0;
//		while (i < 10000) {
//			create(i, "性能测试", "leaf", Long.valueOf(2), Long.valueOf(-1));
//			i++;
//		}

		return testCase;
	}

	@Override
	public TestCase delete(Long vo, Long clientId) {
		// TODO Auto-generated method stub
		return null;
	}
	
	private Integer getChildMaxOrderNumb(TestCase parent) {
		String hql = "select max(orderInParent) from TestCase where parentId = " + parent.getId();
		Integer maxOrder = (Integer) getByHQL(hql);
		
		if (maxOrder == null) {
			maxOrder = 0;
		}
        
		return maxOrder;
	}

}

