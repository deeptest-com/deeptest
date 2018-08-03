package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestPlanDao;
import com.ngtesting.platform.dao.TestTaskDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.AlertService;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.TestTaskService;
import com.ngtesting.platform.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.LinkedList;
import java.util.List;
import java.util.Map;

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
    @Autowired
    TestPlanDao planDao;

    @Override
    public TstTask getById(Integer id) {
        TstTask po = taskDao.getDetail(id);
        TstTask vo = genVo(po);

        return vo;
    }

    @Override
    @Transactional
    public TstTask save(JSONObject json, TstUser user) {
        TstTask task = JSON.parseObject(JSON.toJSONString(json), TstTask.class);
        task.setUserId(user.getId());

        Constant.MsgType action = null;
        if (task.getId() != null) {
            action = Constant.MsgType.update;
            taskDao.update(task);

            taskDao.removeAssignees(task.getId());
        } else {
            action = Constant.MsgType.create;

            if (task.getCaseProjectId() == null) {
                task.setCaseProjectId(task.getProjectId());
            }
            taskDao.save(task);
        }

        List assignees = json.getJSONArray("assignees");
        taskDao.saveAssignees(task.getId(), assignees);

        importSuiteCasesPers(task, JSON.parseObject(JSON.toJSONString(json.get("suites")), List.class));

        alertService.create(task);
        msgService.create(task, action, user);
        historyService.create(task.getProjectId(), user, action.msg, TstHistory.TargetType.task,
                task.getId(), task.getName());

        TstTask ret = taskDao.getDetail(task.getId());
        return ret;
    }

    @Override
    @Transactional
    public boolean importSuiteCasesPers(TstTask task, List<TstSuite> suites) {
        if (suites == null || suites.size() == 0) {
            return false;
        }

        Integer caseProjectId = null;
        List<Integer> suiteIds = new LinkedList<>();
        for (Object obj: suites) {
            TstSuite vo = JSON.parseObject(JSON.toJSONString(obj), TstSuite.class);
            if (vo.getSelecting() != null && vo.getSelecting()) {
                suiteIds.add(vo.getId());

                caseProjectId = vo.getCaseProjectId();
            }
        }

        String suiteIdsStr = StringUtil.join(suiteIds.toArray(), ",");
        taskDao.addCasesBySuites(task.getId(), suiteIdsStr);

        if (caseProjectId != null &&
                (task.getCaseProjectId() == null ||  caseProjectId.intValue() != task.getCaseProjectId().intValue())) {
            taskDao.updateCaseProject(task.getId(), caseProjectId);
        } else {
            return false;
        }

        return true;
    }

    @Override
    @Transactional
    public TstTask saveCases(JSONObject json, TstUser optUser) {
        Integer taskId = json.getInteger("taskId");
        Integer caseProjectId = json.getInteger("caseProjectId");
        JSONArray data = json.getJSONArray("cases");

        taskDao.updateCaseProject(taskId, caseProjectId);

        List<Integer> caseIds = new LinkedList<>();
        for (Object obj : data.toArray()) {
            Integer id = Integer.valueOf(obj.toString());
            caseIds.add(id);
        }

        String ids = StringUtil.join(caseIds.toArray(), ",");
        taskDao.addCases(taskId, ids, false);

        TstTask task = taskDao.get(taskId);
        Constant.MsgType action = Constant.MsgType.update_case;
        msgService.create(task, action, optUser);
        historyService.create(task.getProjectId(), optUser, action.msg, TstHistory.TargetType.task,
                task.getId(), task.getName());

        return task;
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
        planDao.closePlanIfAllTaskClosed(planId);
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
		List<Map> counts = taskDao.countStatus(po.getId());
		for (Map obj : counts) {
			String status = obj.get("status").toString();
			Integer count = Integer.valueOf(obj.get("count").toString());

            po.getCountMap().put(status, count);
            po.getCountMap().put("total", po.getCountMap().get("total") + count);
		}

        String maxStatus = "";
        int maxWidth = 0;
		int sum = 0;
		Integer total = po.getCountMap().get("total");

        Integer barWidth = 200;
        for (String status : po.getCountMap().keySet()) {
		    if ("total".equals(status)) {
		        continue;
            }

            int numb = po.getCountMap().get(status);
            if (total != 0) {
                int width = po.getCountMap().get(status) * barWidth / total;
                if (width > 0) {
                    if (width < 10 && numb < 10) {
                        width = 10;
                    } else if (width < 18 && numb >= 10 && numb < 100) {
                        width = 18;
                    } else if (width < 27 && numb >= 100) {
                        width = 27;
                    }
                }

                po.getWidthMap().put(status, width);

                sum += width;
                if (maxWidth < width) {
                    maxWidth = width;
                    maxStatus = status;
                }
            }
        }
        if (total != 0) {
            po.getWidthMap().put(maxStatus, po.getWidthMap().get(maxStatus) + (barWidth - sum));
        }

		return po;
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

