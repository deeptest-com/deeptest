package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.action.OptFacade;
import com.ngtesting.platform.model.TstCaseInRun;
import com.ngtesting.platform.model.TstRun;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.RunService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

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
	public List<TstCaseInRun> lodaCase(Integer runId) {
//		DetachedCriteria dc = DetachedCriteria.forClass(TstCaseInRun.class);
//
//		if (runId != null) {
//			dc.add(Restrictions.eq("runId", runId));
//		}
//
//		dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//		dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//		dc.addOrder(Order.asc("pId"));
//		dc.addOrder(Order.asc("ordr"));
//
//		List<TstCaseInRun> ls = findAllByCriteria(dc);
//
//		return ls;

        return null;
	}

    @Override
    public TstRun getById(Integer id) {
//        TstRun po = (TstRun) get(TstRun.class, id);
//        TstRun vo = genVo(po);
//
//        return vo;

        return null;
    }

    @Override
    public TstRun save(JSONObject json, TstUser user) {
//        Integer prjId = json.getInteger("prjId");
//        Integer planId = json.getInteger("planId");
//        Integer envId = json.getInteger("envId");
//        Integer runId = json.getInteger("id");
//
//        List assignees = json.getJSONArray("assignees");
//        String runName = json.getString("name");
//
//        Constant.MsgType action = null;
//        TstRun run;
//        if (runId != null) {
//            run = (TstRun) get(TstRun.class, runId);
//            action = Constant.MsgType.update;
//        } else {
//            run = new TstRun();
//            run.setProjectId(prjId);
//            run.setCaseProjectId(prjId);
//            run.setPlanId(planId);
//            action = Constant.MsgType.create;
//        }
//        run.setName(runName);
//        run.setUserId(user.getId());
//        run.setEnvId(envId);
//
//        run.setAssignees(new HashSet());
//        for (Object obj : assignees) {
//            JSONObject jsonObject = JSON.parseObject(obj.toString());
//            TestUser u = (TestUser)get(TestUser.class, jsonObject.getInteger("id"));
//            run.getAssignees().add(u);
//        }
//        run.setUserId(user.getId());
//
//        saveOrUpdate(run);
//
//        importSuiteCasesPers(run, JSON.parseObject(JSON.toJSONString(json.get("suites")), List.class));
//
//        alertService.saveAlert(run);
//        msgService.create(run, action, user);
//        historyService.create(run.getProjectId(), user, action.msg, TestHistory.TargetType.run,
//                run.getId(), run.getName());
//        return run;

        return null;
    }

    @Override
    public boolean importSuiteCasesPers(TstRun run, List<TstSuite> suites) {
//        if (suites == null || suites.size() == 0) {
//            return false;
//        }
//
//        Integer caseProjectId = null;
//        List<Integer> suiteIds = new LinkedList<>();
//        for (Object obj: suites) {
//            TstSuite vo = JSON.parseObject(JSON.toJSONString(obj), TstSuite.class);
//            if (vo.getSelecting() != null && vo.getSelecting()) {
//                suiteIds.add(vo.getId());
//
//                if (caseProjectId == null && run.getCaseProjectId().longValue() != vo.getCaseProjectId().longValue()) {
//                    caseProjectId = vo.getCaseProjectId().longValue();
//                }
//            }
//        }
//        addCasesBySuitesPers(run.getId(), suiteIds);
//        if (caseProjectId != null) {
//            run.setCaseProjectId(caseProjectId);
//            saveOrUpdate(run);
//        }

        return true;
    }

    @Override
    public TstRun saveCases(JSONObject json, TstUser optUser) {
//        Integer projectId = json.getInteger("projectId");
//        Integer caseProjectId = json.getInteger("caseProjectId");
//        Integer planId = json.getInteger("planId");
//        Integer runId = json.getInteger("runId");
//        JSONArray data = json.getJSONArray("cases");
//
//        return saveCases(projectId, caseProjectId, planId, runId, data.toArray(), optUser);

        return null;
    }

    @Override
    public TstRun saveCases(Integer projectId, Integer caseProjectId, Integer planId, Integer runId, Object[] ids, TstUser optUser) {
        TstRun run = null;
//        if (runId != null) {
//            run = (TstRun) get(TstRun.class, runId);
//        } else {
//            run = new TstRun();
//            run.setPlanId(planId);
//        }
//        run.setProjectId(projectId);
//        run.setCaseProjectId(caseProjectId);
//
//        run.setTestcases(new LinkedList<TstCaseInRun>());
//        saveOrUpdate(run);
//
//        List<Integer> caseIds = new LinkedList<>();
//        for (Object obj : ids) {
//            Integer id = Integer.valueOf(obj.toString());
//            caseIds.add(id);
//        }
//        addCasesPers(run.getId(), caseIds);
//
//        Constant.MsgType action = Constant.MsgType.update_case;
//        msgService.create(run, action, optUser);
//        historyService.create(run.getProjectId(), optUser, action.msg, TestHistory.TargetType.run,
//                run.getId(), run.getName());

        return run;
    }

    @Override
    public void addCasesBySuitesPers(Integer id, List<Integer> suiteIds) {
//        String ids = StringUtils.join(suiteIds.toArray(), ",");
//        getDao().querySql("{call add_cases_to_run_by_suites(?,?)}", id, ids);
    }
    @Override
    public void addCasesPers(Integer id, List<Integer> caseIds) {
//        String ids = StringUtils.join(caseIds.toArray(), ",");
//        getDao().querySql("{call add_cases_to_run(?,?,?)}", id, ids, false);
    }

    @Override
    public TstRun delete(Integer id, Integer clientId) {
//        TstRun run = (TstRun) get(TstRun.class, id);
//        run.setDeleted(true);
//        saveOrUpdate(run);
//        return run;

        return null;
    }

    @Override
    public TstRun closePers(Integer id, Integer userId) {
//        TstRun run = (TstRun) get(TstRun.class, id);
//        run.setStatus(TstRun.RunStatus.end);
//        saveOrUpdate(run);
//
//        return run;

        return null;
    }
    @Override
    public void closePlanIfAllRunClosedPers(Integer planId) {
//        getDao().querySql("{call close_plan_if_all_run_closed(?)}", planId);
    }

    @Override
	public List<TstRun> genVos(List<TstRun> pos) {
        List<TstRun> vos = new LinkedList<>();

        for (TstRun po: pos) {
			TstRun vo = genVo(po);
        	vos.add(vo);
        }
		return vos;
	}

	@Override
	public TstRun genVo(TstRun po) {
//		TestUser user = (TestUser)get(TestUser.class, po.getUserId());
//        TstProject project = (TstProject)get(TstProject.class, po.getProjectId());
//        TstProject caseProject = (TstProject)get(TstProject.class, po.getCaseProjectId());
//
//        TstRun vo = new TstRun(po.getId(), po.getName(), po.getEstimate(), po.getStatus().toString(),
//                po.getDescr(), po.getOrdr(),
//                po.getProjectId(), project.getName(),
//                po.getCaseProjectId(), caseProject.getName(),
//                po.getPlanId(), po.getUserId(), user.getName());
//
//        if (po.getEnvId() != null) {
//            vo.setEnvId(po.getEnvId());
//            TstEnv env = (TstEnv)get(TstEnv.class, po.getEnvId());
//            vo.setEnvName(env.getName());
//        }
//
//        for (TestUser u : po.getAssignees()) {
//            TstUser TstUser = new TstUser(u.getId(), u.getName());
//            vo.getAssignees().add(TstUser);
//        }
//
//        String sql = "select tcin.`status` status, count(tcin.id) count from tst_case_in_run tcin "
//                +               "where tcin.run_id  = " + po.getId()
//                +                   " AND tcin.deleted != true AND tcin.disabled != true AND tcin.is_leaf = true "
//                +     " group by tcin.`status`";
//
////		String sql = "select cs1.`status` status, count(cs1.tcin_id) count from "
////                +          "(select tcin.id tcin_id,  tcin.case_id tcin_case_id, tcin.`status` from tst_case_in_run tcin "
////                +               "where tcin.run_id  = " + po.getId()
////                +                   " AND tcin.deleted != true AND tcin.disabled != true) cs1 "
////                +     "where cs1.tcin_case_id not in " // 排除父节点
////                +          "(select distinct tcin.p_id from tst_case_in_run tcin "
////                +               "where tcin.run_id  = " + po.getId() + " and tcin.p_id is not NULL "
////                +                   " AND tcin.deleted != true AND tcin.disabled != true ) "
////                +     "group by cs1.`status`";
//
//		List<Map> counts = findListBySql(sql);
//		for (Map obj : counts) {
//			String status = obj.get("status").toString();
//			Integer count = Integer.valueOf(obj.get("count").toString());
//
//			vo.getCountMap().put(status, count);
//			vo.getCountMap().put("total", vo.getCountMap().get("total") + count);
//		}
//
//        String maxStatus = "";
//        int maxWidth = 0;
//		int sum = 0;
//		Integer total = vo.getCountMap().get("total");
//
//        Integer barWidth = 200;
//        for (String status : vo.getCountMap().keySet()) {
//		    if ("total".equals(status)) {
//		        continue;
//            }
//
//            int numb = vo.getCountMap().get(status);
//            if (total != 0) {
//                int width = vo.getCountMap().get(status) * barWidth / total;
//                if (width > 0) {
//                    if (width < 10 && numb < 10) {
//                        width = 10;
//                    } else if (width < 18 && numb >= 10 && numb < 100) {
//                        width = 18;
//                    } else if (width < 27 && numb >= 100) {
//                        width = 27;
//                    }
//                }
//
//                vo.getWidthMap().put(status, width);
//
//                sum += width;
//                if (maxWidth < width) {
//                    maxWidth = width;
//                    maxStatus = status;
//                }
//            }
//        }
//        if (total != 0) {
//            vo.getWidthMap().put(maxStatus, vo.getWidthMap().get(maxStatus) + (barWidth - sum));
//        }
//
//		return vo;

        return null;
	}

	@Override
	public List<TstCaseInRun> genCaseVos(List<TstCaseInRun> pos) {
		List<TstCaseInRun> vos = new LinkedList();

		for (TstCaseInRun po: pos) {
			TstCaseInRun vo = genCaseVo(po);
			vos.add(vo);
		}
		return vos;
	}

	@Override
	public TstCaseInRun genCaseVo(TstCaseInRun po) {
		TstCaseInRun vo = new TstCaseInRun();

//        TestCase testcase = po.getTestCase();
//		BeanUtilEx.copyProperties(vo, testcase);
//
//		vo.setSteps(new LinkedList<TstCaseStep>());
//
//		List<TestCaseStep> steps = testcase.getSteps();
//		for (TestCaseStep step : steps) {
//			TstCaseStep stepVo = new TstCaseStep(
//					step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());
//
//			vo.getSteps().add(stepVo);
//		}
		return vo;
	}

	private Integer getChildMaxOrderNumb(TstRun parent) {
//		String hql = "select max(ordr) from TstRun where parentId = " + parent.getId();
//		Integer maxOrder = (Integer) getByHQL(hql);
//
//		if (maxOrder == null) {
//			maxOrder = 0;
//		}
//
//		return maxOrder;

        return 1;
	}

}

