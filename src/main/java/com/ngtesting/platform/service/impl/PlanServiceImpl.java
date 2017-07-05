package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestPlan;
import com.ngtesting.platform.service.PlanService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestPlanVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class PlanServiceImpl extends BaseServiceImpl implements PlanService {

	@Override
	public List<TestPlan> query(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestPlan.class);
        
        if (projectId != null) {
        	dc.add(Restrictions.eq("projectId", projectId));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.desc("createTime"));
        dc.addOrder(Order.asc("id"));
        List<TestPlan> ls = findAllByCriteria(dc);
        
        return ls;
	}

	@Override
	public TestPlanVo getById(Long caseId) {
		TestPlan po = (TestPlan) get(TestPlan.class, caseId);
		TestPlanVo vo = genVo(po);

		return vo;
	}

	@Override
	public List<TestPlanVo> genVos(List<TestPlan> pos) {
        List<TestPlanVo> vos = new LinkedList<TestPlanVo>();

        for (TestPlan po: pos) {
        	TestPlanVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TestPlanVo genVo(TestPlan po) {
		TestPlanVo vo = new TestPlanVo();
		BeanUtilEx.copyProperties(vo, po);

		return vo;
	}

	@Override
	public TestPlan save(JSONObject json) {
		TestPlan testCase = (TestPlan) get(TestPlan.class, json.getLong("id"));
		return testCase;
	}

	@Override
	public TestPlan delete(Long vo, Long clientId) {
		// TODO Auto-generated method stub
		return null;
	}
	
	private Integer getChildMaxOrderNumb(TestPlan parent) {
		String hql = "select max(ordr) from TestPlan where parentId = " + parent.getId();
		Integer maxOrder = (Integer) getByHQL(hql);
		
		if (maxOrder == null) {
			maxOrder = 0;
		}
        
		return maxOrder;
	}

}

