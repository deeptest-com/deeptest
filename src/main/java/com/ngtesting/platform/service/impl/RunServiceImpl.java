package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
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
import com.ngtesting.platform.vo.*;
import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.HashSet;
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
    public TestRun save(JSONObject json, UserVo user) {
        Long prjId = json.getLong("prjId");
        Long planId = json.getLong("planId");
        Long envId = json.getLong("envId");
        Long runId = json.getLong("id");

        List assignees = json.getJSONArray("assignees");
        String runName = json.getString("name");

        Constant.MsgType action = null;
        TestRun run;
        if (runId != null) {
            run = (TestRun) get(TestRun.class, runId);
            action = Constant.MsgType.update;
        } else {
            run = new TestRun();
            run.setProjectId(prjId);
            run.setCaseProjectId(prjId);
            run.setPlanId(planId);
            action = Constant.MsgType.create;
        }
        run.setName(runName);
        run.setUserId(user.getId());
        run.setEnvId(envId);

        run.setAssignees(new HashSet());
        for (Object obj : assignees) {
            JSONObject jsonObject = JSON.parseObject(obj.toString());
            TestUser u = (TestUser)get(TestUser.class, jsonObject.getLong("id"));
            run.getAssignees().add(u);
        }
        run.setUserId(user.getId());

        saveOrUpdate(run);

        importSuiteCasesPers(run, JSON.parseObject(JSON.toJSONString(json.get("suites")), List.class));

        alertService.saveAlert(run);
        msgService.create(run, action, user);
        historyService.create(run.getProjectId(), user, action.msg, TestHistory.TargetType.run,
                run.getId(), run.getName());
        return run;
    }

    @Override
    public boolean importSuiteCasesPers(TestRun run, List<TestSuiteVo> suites) {
        if (suites == null || suites.size() == 0) {
            return false;
        }

        Long caseProjectId = null;
        List<Long> suiteIds = new LinkedList<>();
        for (Object obj: suites) {
            TestSuiteVo vo = JSON.parseObject(JSON.toJSONString(obj), TestSuiteVo.class);
            if (vo.getSelecting() != null && vo.getSelecting()) {
                suiteIds.add(vo.getId());

                if (caseProjectId == null && run.getCaseProjectId().longValue() != vo.getCaseProjectId().longValue()) {
                    caseProjectId = vo.getCaseProjectId().longValue();
                }
            }
        }
        addCasesBySuitesPers(run.getId(), suiteIds);
        if (caseProjectId != null) {
            run.setCaseProjectId(caseProjectId);
            saveOrUpdate(run);
        }

        return true;
    }

    @Override
    public TestRun saveCases(JSONObject json, UserVo optUser) {
        Long projectId = json.getLong("projectId");
        Long caseProjectId = json.getLong("caseProjectId");
        Long planId = json.getLong("planId");
        Long runId = json.getLong("runId");
        JSONArray data = json.getJSONArray("cases");

        return saveCases(projectId, caseProjectId, planId, runId, data.toArray(), optUser);
    }

    @Override
    public TestRun saveCases(Long projectId, Long caseProjectId, Long planId, Long runId, Object[] ids, UserVo optUser) {
        TestRun run;
        if (runId != null) {
            run = (TestRun) get(TestRun.class, runId);
        } else {
            run = new TestRun();
            run.setPlanId(planId);
        }
        run.setProjectId(projectId);
        run.setCaseProjectId(caseProjectId);

        run.setTestcases(new LinkedList<TestCaseInRun>());
        saveOrUpdate(run);

        List<Long> caseIds = new LinkedList<>();
        for (Object obj : ids) {
            Long id = Long.valueOf(obj.toString());
            caseIds.add(id);
        }
        addCasesPers(run.getId(), caseIds);

        Constant.MsgType action = Constant.MsgType.update_case;
        msgService.create(run, action, optUser);
        historyService.create(run.getProjectId(), optUser, action.msg, TestHistory.TargetType.run,
                run.getId(), run.getName());

        return run;
    }

    @Override
    public void addCasesBySuitesPers(Long id, List<Long> suiteIds) {
        String ids = StringUtils.join(suiteIds.toArray(), ",");
        getDao().querySql("{call add_cases_to_run_by_suites(?,?)}", id, ids);
    }
    @Override
    public void addCasesPers(Long id, List<Long> caseIds) {
        String ids = StringUtils.join(caseIds.toArray(), ",");
        getDao().querySql("{call add_cases_to_run(?,?,?)}", id, ids, false);
    }

    @Override
    public TestRun delete(Long id, Long clientId) {
        TestRun run = (TestRun) get(TestRun.class, id);
        run.setDeleted(true);
        saveOrUpdate(run);
        return run;
    }

    @Override
    public TestRun closePers(Long id, Long userId) {
        TestRun run = (TestRun) get(TestRun.class, id);
        run.setStatus(TestRun.RunStatus.end);
        saveOrUpdate(run);

        return run;
    }
    @Override
    public void closePlanIfAllRunClosedPers(Long planId) {
        getDao().querySql("{call close_plan_if_all_run_closed(?)}", planId);
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
		TestUser user = (TestUser)get(TestUser.class, po.getUserId());
        TestProject project = (TestProject)get(TestProject.class, po.getProjectId());
        TestProject caseProject = (TestProject)get(TestProject.class, po.getCaseProjectId());

        TestRunVo vo = new TestRunVo(po.getId(), po.getName(), po.getEstimate(), po.getStatus().toString(),
                po.getDescr(), po.getOrdr(),
                po.getProjectId(), project.getName(),
                po.getCaseProjectId(), caseProject.getName(),
                po.getPlanId(), po.getUserId(), user.getName());

        if (po.getEnvId() != null) {
            vo.setEnvId(po.getEnvId());
            TestEnv env = (TestEnv)get(TestEnv.class, po.getEnvId());
            vo.setEnvName(env.getName());
        }

        for (TestUser u : po.getAssignees()) {
            UserVo userVo = new UserVo(u.getId(), u.getName());
            vo.getAssignees().add(userVo);
        }

        String sql = "select tcin.`status` status, count(tcin.id) count from tst_case_in_run tcin "
                +               "where tcin.run_id  = " + po.getId()
                +                   " AND tcin.deleted != true AND tcin.disabled != true AND tcin.is_leaf = true "
                +     " group by tcin.`status`";

//		String sql = "select cs1.`status` status, count(cs1.tcin_id) count from "
//                +          "(select tcin.id tcin_id,  tcin.case_id tcin_case_id, tcin.`status` from tst_case_in_run tcin "
//                +               "where tcin.run_id  = " + po.getId()
//                +                   " AND tcin.deleted != true AND tcin.disabled != true) cs1 "
//                +     "where cs1.tcin_case_id not in " // 排除父节点
//                +          "(select distinct tcin.p_id from tst_case_in_run tcin "
//                +               "where tcin.run_id  = " + po.getId() + " and tcin.p_id is not NULL "
//                +                   " AND tcin.deleted != true AND tcin.disabled != true ) "
//                +     "group by cs1.`status`";

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

        Integer barWidth = 200;
        for (String status : vo.getCountMap().keySet()) {
		    if ("total".equals(status)) {
		        continue;
            }

            int numb = vo.getCountMap().get(status);
            if (total != 0) {
                int width = vo.getCountMap().get(status) * barWidth / total;
                if (width > 0) {
                    if (width < 10 && numb < 10) {
                        width = 10;
                    } else if (width < 18 && numb >= 10 && numb < 100) {
                        width = 18;
                    } else if (width < 27 && numb >= 100) {
                        width = 27;
                    }
                }

                vo.getWidthMap().put(status, width);

                sum += width;
                if (maxWidth < width) {
                    maxWidth = width;
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

