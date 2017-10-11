package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestCase;
import com.ngtesting.platform.entity.TestCaseInRun;
import com.ngtesting.platform.entity.TestCaseStep;
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
    public TestRun save(JSONObject json) {
        Long planId = json.getLong("planId");
        Long runId = json.getLong("id");
        String runName = json.getString("name");

        TestRun run;
        if (runId != null) {
            run = (TestRun) get(TestRun.class, runId);
        } else {
            run = new TestRun();
            run.setPlanId(planId);
        }
        run.setName(runName);
        saveOrUpdate(run);

        return run;
    }

    @Override
    public TestRun saveCases(JSONObject json) {
        Long planId = json.getLong("planId");
        Long runId = json.getLong("runId");
        JSONArray data = json.getJSONArray("cases");

        TestRun run;
        if (runId != null) {
            run = (TestRun) get(TestRun.class, runId);
        } else {
            run = new TestRun();
            run.setPlanId(planId);
        }

        for (TestCaseInRun item : run.getTestcases()) {
            getDao().delete(item);
        }

        run.setTestcases(new LinkedList<TestCaseInRun>());
        saveOrUpdate(run);
        for (Object obj : data) {
            Long id = Long.valueOf(obj.toString());

            TestCaseInRun caseInRun = new TestCaseInRun(runId, id);
            run.getTestcases().add(caseInRun);
        }
        saveOrUpdate(run);

        return run;
    }

    @Override
    public TestRun delete(Long vo, Long clientId) {

        return null;
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

//        for (TestCaseInRun testcase : po.getTestcases()) {
//            vo.getCaseIds().add(testcase.getCaseId());
//        }

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

        TestCase testcase = po.getTestCase();
		BeanUtilEx.copyProperties(vo, testcase);

		vo.setSteps(new LinkedList<TestCaseStepVo>());

		List<TestCaseStep> steps = testcase.getSteps();
		for (TestCaseStep step : steps) {
			TestCaseStepVo stepVo = new TestCaseStepVo(
					step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());

			vo.getSteps().add(stepVo);
		}
		return vo;
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

