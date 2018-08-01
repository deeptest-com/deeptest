package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.TestTaskDao;
import com.ngtesting.platform.model.TstCaseInTask;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.TestTaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class TestTaskServiceImpl extends BaseServiceImpl implements TestTaskService {

    @Autowired
    MsgService msgService;

    @Autowired
    AlertService alertService;

    @Autowired
    HistoryService historyService;

    @Autowired
    TestTaskDao taskDao;

    @Override
    public TstTask getById(Integer id) {
        TstTask po = taskDao.get(id);
        TstTask vo = genVo(po);

        return vo;
    }

    @Override
    public TstTask save(JSONObject json, TstUser user) {
//        Integer prjId = json.getInteger("prjId");
//        Integer planId = json.getInteger("planId");
//        Integer envId = json.getInteger("envId");
//        Integer taskId = json.getInteger("id");
//
//        List assignees = json.getJSONArray("assignees");
//        String taskName = json.getString("name");
//
//        Constant.MsgType action = null;
//        TstTask task;
//        if (taskId != null) {
//            task = (TstTask) get(TstTask.class, taskId);
//            action = Constant.MsgType.update;
//        } else {
//            task = new TstTask();
//            task.setProjectId(prjId);
//            task.setCaseProjectId(prjId);
//            task.setPlanId(planId);
//            action = Constant.MsgType.create;
//        }
//        task.setName(taskName);
//        task.setUserId(user.getId());
//        task.setEnvId(envId);
//
//        task.setAssignees(new HashSet());
//        for (Object obj : assignees) {
//            JSONObject jsonObject = JSON.parseObject(obj.toString());
//            TestUser u = (TestUser)get(TestUser.class, jsonObject.getInteger("id"));
//            task.getAssignees().add(u);
//        }
//        task.setUserId(user.getId());
//
//        saveOrUpdate(task);
//
//        importSuiteCasesPers(task, JSON.parseObject(JSON.toJSONString(json.get("suites")), List.class));
//
//        alertService.saveAlert(task);
//        msgService.create(task, action, user);
//        historyService.create(task.getProjectId(), user, action.msg, TestHistory.TargetType.task,
//                task.getId(), task.getName());
//        return task;

        return null;
    }

    @Override
    public boolean importSuiteCasesPers(TstTask task, List<TstSuite> suites) {
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
//                if (caseProjectId == null && task.getCaseProjectId().longValue() != vo.getCaseProjectId().longValue()) {
//                    caseProjectId = vo.getCaseProjectId().longValue();
//                }
//            }
//        }
//        addCasesBySuitesPers(task.getId(), suiteIds);
//        if (caseProjectId != null) {
//            task.setCaseProjectId(caseProjectId);
//            saveOrUpdate(task);
//        }

        return true;
    }

    @Override
    public TstTask saveCases(JSONObject json, TstUser optUser) {
        Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
        Integer planId = json.getInteger("planId");
        Integer taskId = json.getInteger("taskId");
        JSONArray data = json.getJSONArray("cases");

        return saveCases(projectId, caseProjectId, planId, taskId, data.toArray(), optUser);
    }

    @Override
    public TstTask saveCases(Integer projectId, Integer caseProjectId, Integer planId, Integer taskId, Object[] ids, TstUser optUser) {
        TstTask task = null;
//        if (taskId != null) {
//            task = (TstTask) get(TstTask.class, taskId);
//        } else {
//            task = new TstTask();
//            task.setPlanId(planId);
//        }
//        task.setProjectId(projectId);
//        task.setCaseProjectId(caseProjectId);
//
//        task.setTestCases(new LinkedList<TstCaseInTask>());
//        saveOrUpdate(task);
//
//        List<Integer> caseIds = new LinkedList<>();
//        for (Object obj : ids) {
//            Integer id = Integer.valueOf(obj.toString());
//            caseIds.add(id);
//        }
//        addCasesPers(task.getId(), caseIds);
//
//        Constant.MsgType action = Constant.MsgType.update_case;
//        msgService.create(task, action, optUser);
//        historyService.create(task.getProjectId(), optUser, action.msg, TestHistory.TargetType.task,
//                task.getId(), task.getName());

        return task;
    }

    @Override
    public void addCasesBySuitesPers(Integer id, List<Integer> suiteIds) {
//        String ids = StringUtils.join(suiteIds.toArray(), ",");
//        getDao().querySql("{call add_cases_to_task_by_suites(?,?)}", id, ids);
    }
    @Override
    public void addCasesPers(Integer id, List<Integer> caseIds) {
//        String ids = StringUtils.join(caseIds.toArray(), ",");
//        getDao().querySql("{call add_cases_to_task(?,?,?)}", id, ids, false);
    }

    @Override
    public void delete(Integer id, Integer userId) {
        taskDao.delete(id, userId);
    }

    @Override
    public void closePers(Integer id, Integer userId) {
        taskDao.close(id, userId);
    }
    @Override
    public void closePlanIfAllTaskClosedPers(Integer planId) {
        taskDao.closePlanIfAllTaskClosed(planId);
//        getDao().querySql("{call close_plan_if_all_task_closed(?)}", planId);
    }

    @Override
    public List<TstTask> listByPlan(Integer planId) {
        List<TstTask> tasks = taskDao.listByPlan(planId);
        return genVos(tasks);
    }

    @Override
	public List<TstTask> genVos(List<TstTask> pos) {
        for (TstTask po: pos) {
			genVo(po);
        }
		return pos;
	}

	@Override
	public TstTask genVo(TstTask po) {

//        for (TestUser u : po.getAssignees()) {
//            TstUser TstUser = new TstUser(u.getId(), u.getName());
//            vo.getAssignees().add(TstUser);
//        }
//
//        String sql = "select tcin.`status` status, count(tcin.id) count from tst_case_in_task tcin "
//                +               "where tcin.task_id  = " + po.getId()
//                +                   " AND tcin.deleted != true AND tcin.disabled != true AND tcin.is_leaf = true "
//                +     " group by tcin.`status`";
//
////		String sql = "select cs1.`status` status, count(cs1.tcin_id) count from "
////                +          "(select tcin.id tcin_id,  tcin.case_id tcin_case_id, tcin.`status` from tst_case_in_task tcin "
////                +               "where tcin.task_id  = " + po.getId()
////                +                   " AND tcin.deleted != true AND tcin.disabled != true) cs1 "
////                +     "where cs1.tcin_case_id not in " // 排除父节点
////                +          "(select distinct tcin.p_id from tst_case_in_task tcin "
////                +               "where tcin.task_id  = " + po.getId() + " and tcin.p_id is not NULL "
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
	public List<TstCaseInTask> genCaseVos(List<TstCaseInTask> pos) {
		List<TstCaseInTask> vos = new LinkedList();

		for (TstCaseInTask po: pos) {
			TstCaseInTask vo = genCaseVo(po);
			vos.add(vo);
		}
		return vos;
	}

	@Override
	public TstCaseInTask genCaseVo(TstCaseInTask po) {
		TstCaseInTask vo = new TstCaseInTask();

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

	private Integer getChildMaxOrderNumb(TstTask parent) {
//		String hql = "select max(ordr) from TstTask where parentId = " + parent.getId();
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

