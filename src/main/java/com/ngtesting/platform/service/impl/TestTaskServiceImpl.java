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
    public TstTask getById(Integer id, Integer projectId) {
        TstTask po = taskDao.getDetail(id, projectId);
        if (po == null) {
            return null;
        }

        TstTask vo = genVo(po);
        return vo;
    }

    @Override
    @Transactional
    public TstTask save(JSONObject json, TstUser user) {
        TstTask task = JSON.parseObject(JSON.toJSONString(json), TstTask.class);
        task.setUserId(user.getId());
        task.setProjectId(user.getDefaultPrjId());

        if (task.getCaseProjectId() == null) {
            task.setCaseProjectId(task.getProjectId());
        }

        Constant.MsgType action = null;
        if (task.getId() == null) {
            action = Constant.MsgType.create;

            taskDao.save(task);
        } else {
            action = Constant.MsgType.update;
            Integer count = taskDao.update(task);
            if (count == null) {
                return null;
            }

            taskDao.removeAssignees(task.getId());
        }

        List assignees = json.getJSONArray("assignees");
        taskDao.saveAssignees(task.getId(), assignees);

        List<TstSuite> suites = JSON.parseObject(JSON.toJSONString(json.get("suites")), List.class);

        importSuiteCasesPers(task, suites, user);

        alertService.create(task);
        msgService.create(task, action, user);
        historyService.create(task.getProjectId(), user, action.msg, TstHistory.TargetType.task,
                task.getId(), task.getName());

        TstTask ret = taskDao.getDetail(task.getId(), null);
        return ret;
    }

    @Override
    @Transactional
    public void importSuiteCasesPers(TstTask task, List<TstSuite> suites, TstUser optUser) {
        if (suites == null || suites.size() == 0) {
            return;
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

        if (suiteIds.size() == 0) {
            return;
        }

        String suiteIdsStr = StringUtil.join(suiteIds.toArray(), ",");
        taskDao.addCasesBySuites(task.getId(), suiteIdsStr);

        if (caseProjectId != null &&
                (task.getCaseProjectId() == null ||  caseProjectId.intValue() != task.getCaseProjectId().intValue())) {
            taskDao.updateCaseProject(task.getId(), caseProjectId);
        }

        msgService.create(task, Constant.MsgType.create, optUser);
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
    public Boolean delete(Integer id, Integer projectId) {
        Integer count = taskDao.delete(id, projectId);
        return count > 0;
    }

    @Override
    public Boolean close(Integer id, Integer projectId) {
        Integer count = taskDao.close(id, projectId);
        return count > 0;
    }

    @Override
    public void closePlanIfAllTaskClosed(Integer planId) {
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

}

