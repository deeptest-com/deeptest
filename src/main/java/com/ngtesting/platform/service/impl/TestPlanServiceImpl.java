package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestPlanDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.TestPlanService;
import com.ngtesting.platform.service.TestTaskService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TestPlanServiceImpl extends BaseServiceImpl implements TestPlanService {
//    @Autowired
//    ProjectService projectService;

    @Autowired
    HistoryService historyService;

    @Autowired
    TestPlanDao testPlanDao;
    @Autowired
    TestTaskService taskService;

    @Override
    public List<TstPlan> listByPage(Integer projectId, String keywords, String status) {
        List<TstPlan> pos = testPlanDao.query(projectId, keywords, status);
        return pos;
    }

    @Override
    public TstPlan getById(Integer id) {
        TstPlan po;
        if (id != null) {
            po = testPlanDao.get(id);
            genVo(po);
        } else {
            po = new TstPlan();
        }
        return po;
    }

    @Override
    public TstPlan save(JSONObject json, TstUser user) {
        TstPlan vo = JSON.parseObject(JSON.toJSONString(json), TstPlan.class);
        vo.setUserId(user.getId());

        Constant.MsgType action;
        if (vo.getId() != null) {
            action = Constant.MsgType.update;

            testPlanDao.update(vo);
        } else {
            action = Constant.MsgType.create;

            testPlanDao.save(vo);
        }

        historyService.create(vo.getProjectId(), user, action.msg, TstHistory.TargetType.plan,
                vo.getId(), vo.getName());

        return vo;
    }

    @Override
    public void delete(Integer id, Integer clientId) {
        testPlanDao.delete(id);
    }

    @Override
    public List<TstPlan> listByOrg(Integer orgId) {
        PageHelper.startPage(0, 10);
        List<TstPlan> ls = testPlanDao.listByOrg(orgId);

        return ls;
    }

    @Override
    public List<TstPlan> listByProject(Integer projectId, TstProject.ProjectType projectType) {
        List<TstPlan> pos;

        if (projectType.equals(TstProject.ProjectType.project)) {
            pos = testPlanDao.listByProject(projectId);
        } else {
            pos = testPlanDao.listByProjectGroup(projectId);
        }

        return pos;
    }

    @Override
    public void genVos(List<TstPlan> pos) {
        for (TstPlan po : pos) {
            genVo(po);
        }
    }

    @Override
    public void genVo(TstPlan po) {
        List<TstTask> tasks = taskService.listByPlan(po.getId());
        po.setTasks(tasks);
    }

}

