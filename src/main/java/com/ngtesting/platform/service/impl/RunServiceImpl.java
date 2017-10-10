package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCaseInRun;
import com.ngtesting.platform.entity.TestCaseStepInRun;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.service.RunService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseInRunVo;
import com.ngtesting.platform.vo.TestCaseStepVo;
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
	public List<TestCaseInRun> lodaCase(Long runId) {
		DetachedCriteria dc = DetachedCriteria.forClass(TestCaseInRun.class);

		if (runId != null) {
			dc.add(Restrictions.eq("runId", runId));
		}

		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
		dc.add(Restrictions.eq("disabled", Boolean.FALSE));

		dc.addOrder(Order.asc("pId"));
		dc.addOrder(Order.asc("ordr"));

		List<TestCaseInRun> ls = findAllByCriteria(dc);

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

		String hql = "select cs.status, count(cs.id) from TestCaseInRun cs where cs.runId = ? group by cs.status";

		List counts = getListByHQL(hql, po.getId());
		for (Object obj : counts) {
			Object[] arr = (Object[])obj;
			String status = arr[0].toString();
			Integer count = Integer.valueOf(arr[1].toString());

			vo.getCountMap().put(status, count);
			vo.getCountMap().put("total", vo.getCountMap().get("total") + count);
		}

		return vo;
	}

	@Override
	public List<TestCaseInRunVo> genCaseVos(List<TestCaseInRun> pos) {
		List<TestCaseInRunVo> vos = new LinkedList();

		for (TestCaseInRun po: pos) {
			TestCaseInRunVo vo = genCaseVo(po);
			vos.add(vo);
		}
		return vos;
	}

	@Override
	public TestCaseInRunVo genCaseVo(TestCaseInRun po) {
		TestCaseInRunVo vo = new TestCaseInRunVo();

		BeanUtilEx.copyProperties(vo, po);

		vo.setSteps(new LinkedList<TestCaseStepVo>());

		List<TestCaseStepInRun> steps = po.getSteps();
		for (TestCaseStepInRun step : steps) {
			TestCaseStepVo stepVo = new TestCaseStepVo(
					step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseInRunId());

			vo.getSteps().add(stepVo);
		}
		return vo;
	}

	@Override
	public TestRun save(JSONObject json) {
		TestRun run = (TestRun) get(TestRun.class, json.getLong("runId"));

		return run;
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

