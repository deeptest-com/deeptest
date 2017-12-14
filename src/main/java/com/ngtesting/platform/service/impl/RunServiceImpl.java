package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.bean.websocket.OptFacade;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.*;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.RunService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseInRunVo;
import com.ngtesting.platform.vo.TestCaseStepVo;
import com.ngtesting.platform.vo.TestRunVo;
import com.ngtesting.platform.vo.UserVo;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

@Service
public class RunServiceImpl extends BaseServiceImpl implements RunService {
    @Autowired
    private OptFacade optFacade;

    @Autowired
    MsgService msgService;

    @Autowired
    AlertService alertService;

    @Autowired
    HistoryService historyService;

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
    public TestRun save(JSONObject json, UserVo optUser) {
        Long planId = json.getLong("planId");
        Long runId = json.getLong("id");
        Long userId = json.getLong("userId");
        String runName = json.getString("name");

        Constant.MsgType action = null;
        TestRun run;
        if (runId != null) {
            run = (TestRun) get(TestRun.class, runId);
            action = Constant.MsgType.update;
        } else {
            run = new TestRun();
            run.setPlanId(planId);
            action = Constant.MsgType.create;
        }
        run.setName(runName);
        run.setUserId(userId);
        saveOrUpdate(run);

        msgService.create(run, action, optUser);
        historyService.create(run.getProjectId(), optUser, action.msg, TestHistory.TargetType.run,
                run.getId(), run.getName());

        return run;
    }

    @Override
    public TestRun saveCases(JSONObject json, UserVo optUser) {
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
            TestCase testcase = (TestCase) get(TestCase.class, id);

            TestCaseInRun caseInRun = new TestCaseInRun(runId, id, testcase.getOrdr(), testcase.getpId());
            run.getTestcases().add(caseInRun);
        }
        saveOrUpdate(run);

        Constant.MsgType action = Constant.MsgType.update_case;
        msgService.create(run, action, optUser);
        historyService.create(run.getProjectId(), optUser, action.msg, TestHistory.TargetType.run,
                run.getId(), run.getName());

        return run;
    }

    @Override
    public TestRun delete(Long id, Long clientId) {
        TestRun run = (TestRun) get(TestRun.class, id);
        run.setDeleted(true);
        saveOrUpdate(run);
        return run;
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
		TestUser user = (TestUser)get(TestUser.class, po.getUserId());
        vo.setUserName(user.getName());

		String sql = "select cs1.`status` status, count(cs1.tcinid) count from "
                +          "(select tcin.id tcinid, tcin.`status`, tc.id tcid from tst_case_in_run tcin "
                +               "left join tst_case tc on tcin.case_id = tc.id "
                +               "where tcin.run_id  = " + po.getId() + " order by tc.ordr) cs1 "
                +     "where cs1.tcid not in "
                +          "(select distinct tc.p_id from tst_case_in_run tcin left join tst_case tc on tcin.case_id = tc.id "
                +               "where tcin.run_id  = " + po.getId() + " and tc.p_id is not NULL) "
                +     "group by cs1.`status`";

		List<Map> counts = findListBySql(sql);
		for (Map obj : counts) {
			String status = obj.get("status").toString();
			Integer count = Integer.valueOf(obj.get("count").toString());

			vo.getCountMap().put(status, count);
			vo.getCountMap().put("total", vo.getCountMap().get("total") + count);
		}

        String maxStatus = "";
        int maxWidth = 0;
		int sum = 0;
		Integer total = vo.getCountMap().get("total");

        Integer barWidth = 240;
        for (String status : vo.getCountMap().keySet()) {
		    if ("total".equals(status)) {
		        continue;
            }

            if (total != 0) {
                int width = vo.getCountMap().get(status) * barWidth / total;
                vo.getWidthMap().put(status, width);

                sum += width;
                if (maxWidth < width) {
                    maxStatus = status;
                }
            }
        }
        if (total != 0) {
            vo.getWidthMap().put(maxStatus, vo.getWidthMap().get(maxStatus) + (barWidth - sum));
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

