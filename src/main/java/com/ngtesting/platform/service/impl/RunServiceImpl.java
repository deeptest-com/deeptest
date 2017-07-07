package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.service.RunService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestRunVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class RunServiceImpl extends BaseServiceImpl implements RunService {

	@Override
	public List<TestRun> query(Long projectId) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestRun.class);

        if (projectId != null) {
        	dc.add(Restrictions.eq("projectId", projectId));
        }
        
        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.asc("ordr"));
        dc.addOrder(Order.asc("id"));
        List<TestRun> ls = findAllByCriteria(dc);
        
        return ls;
	}

	@Override
	public TestRunVo getById(Long id) {
		TestRun po = (TestRun) get(TestRun.class, id);
		TestRunVo vo = genVo(po);

		return vo;
	}

	@Override
	public List<TestRunVo> genVos(List<TestRun> pos) {
        List<TestRunVo> vos = new LinkedList<>();

        for (TestRun po: pos) {
			TestRunVo vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TestRunVo genVo(TestRun po) {
		TestRunVo vo = new TestRunVo();
		BeanUtilEx.copyProperties(vo, po);

		String hql = "select cs.status, count(cs.id) from TestRunToCase cs where cs.runId = ? group by cs.status";

		List counts = getListByHQL(hql, po.getId());
		for (Object obj : counts) {
			Object[] arr = (Object[])obj;
			String status = arr[0].toString();
			Integer count = Integer.valueOf(arr[1].toString());

			vo.getCountMap().put(status, Integer.valueOf(count));
			vo.getCountMap().put("total", vo.getCountMap().get("total") + count);
		}

		return vo;
	}

	@Override
	public TestRun save(JSONObject json) {
		TestRun testCase = (TestRun) get(TestRun.class, json.getLong("id"));

		return testCase;
	}

	@Override
	public TestRun delete(Long vo, Long clientId) {

		return null;
	}
	
	private Integer getChildMaxOrderNumb(TestRun parent) {
		String hql = "select max(ordr) from TestRun where parentId = " + parent.getId();
		Integer maxOrder = (Integer) getByHQL(hql);
		
		if (maxOrder == null) {
			maxOrder = 0;
		}
        
		return maxOrder;
	}

}

