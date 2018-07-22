package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.github.pagehelper.PageHelper;
import com.ngtesting.platform.dao.TestPlanDao;
import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstProject;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.TestPlanService;
import com.ngtesting.platform.service.TestTaskService;
import com.ngtesting.platform.vo.Page;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.List;

@Service
public class TestPlanServiceImpl extends BaseServiceImpl implements TestPlanService {
    @Autowired
    TestPlanService testPlanService;

//    @Autowired
//    ProjectService projectService;

//
//    @Autowired
//    HistoryService historyService;

    @Autowired
    TestPlanDao testPlanDao;
    @Autowired
    TestTaskService taskService;

    @Override
    public Page page(Integer projectId, String status, String keywords, Integer currentPage, Integer itemsPerPage) {

//        DetachedCriteria dc = DetachedCriteria.forClass(TstPlan.class);
//
//        if (projectId != null) {
//            List<Integer> ids = projectService.listBrotherIds(projectId);
//            dc.add(Restrictions.in("projectId", ids));
//        }
//
//        if (StringUtils.isNotEmpty(status)) {
//            dc.add(Restrictions.eq("status", TstPlan.PlanStatus.valueOf(status)));
//        }
//        if (StringUtils.isNotEmpty(keywords)) {
//            dc.add(Restrictions.like("name", "%" + keywords + "%"));
//        }
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.desc("createTime"));
//        dc.addOrder(Order.asc("id"));
//        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
//
//        return page;

        return null;
    }

    @Override
    public TstPlan getById(Integer caseId) {
//        if (caseId == null) {
//            return null;
//        }
//        TstPlan po = (TstPlan) get(TstPlan.class, caseId);
//        TstPlan vo = genVo(po);
//
//        return vo;

        return null;
    }

    @Override
    public TstPlan save(JSONObject json, TstUser optUser) {
//        Integer id = json.getInteger("id");
//
//        TstPlan po;
//        TstPlan vo = JSON.parseObject(JSON.toJSONString(json), TstPlan.class);
//
//        Constant.MsgType action;
//        if (id != null) {
//            po = (TstPlan)get(TstPlan.class, id);
//            action = Constant.MsgType.update;
//        } else {
//            po = new TstPlan();
//            action = Constant.MsgType.create;
//        }
//        po.setName(vo.getName());
//        po.setEstimate(vo.getEstimate());
//        po.setStartTime(vo.getStartTime());
//        po.setEndTime(vo.getEndTime());
//        po.setDescr(vo.getDescr());
//        po.setProjectId(vo.getProjectId());
//        po.setVerId(vo.getVerId());
//
//        saveOrUpdate(po);
//
//        historyService.create(po.getProjectId(), optUser, action.msg, TestHistory.TargetType.plan,
//                po.getId(), po.getName());
//
//        return po;

        return null;
    }

    @Override
    public TstPlan delete(Integer id, Integer clientId) {
//        TstPlan po = (TstPlan)get(TstPlan.class, id);
//        po.setDeleted(true);
//        saveOrUpdate(po);
//        return po;

        return null;
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

    private Integer getChildMaxOrderNumb(TstPlan parent) {
//        String hql = "select max(ordr) from TstPlan where parentId = " + parent.getId();
//        Integer maxOrder = (Integer) getByHQL(hql);
//
//        if (maxOrder == null) {
//            maxOrder = 0;
//        }
//
//        return maxOrder;

        return 1;
    }

    public List traverseFolder(String path, List<String> fileList) {
        File file = new File(path);
        if (file.exists()) {
            File[] files = file.listFiles();
            if (files.length == 0) {
                System.out.println("文件夹是空的!");
            } else {
                for (File file2 : files) {
                    if (file2.isDirectory()) {
                        System.out.println("文件夹:" + file2.getAbsolutePath());
                        traverseFolder(file2.getAbsolutePath(), fileList);
                    } else {
                        System.out.println("文件:" + file2.getAbsolutePath());
                        if (file2.getAbsolutePath().lastIndexOf(".txt") > 0) {
                            fileList.add(file2.getAbsolutePath());
                        }
                    }
                }
            }
        } else {
            System.out.println("文件不存在!");
        }

        return fileList;
    }

    @Override
    public List<TstPlan> genVos(List<TstPlan> pos) {
        for (TstPlan po : pos) {
            genVo(po);
        }
        return pos;
    }

    @Override
    public TstPlan genVo(TstPlan po) {
        List<TstTask> runs = taskService.listByPlan(po.getId());
        po.setRuns(runs);

        return po;
    }

    @Override
    public TstPlan updatePo(TstPlan vo) {
        TstPlan po = new TstPlan();
//        po.setName(vo.getName());
//        po.setName(vo.getName());
//        po.setEstimate(vo.getEstimate());
//        po.setStartTime(vo.getStartTime());
//        po.setEndTime(vo.getEndTime());
//        po.setDescr(vo.getDescr());
//        po.setProjectId(vo.getProjectId());

        return po;
    }

}

