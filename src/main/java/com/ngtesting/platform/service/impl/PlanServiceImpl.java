package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstPlan;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.PlanService;
import com.ngtesting.platform.vo.Page;
import org.springframework.stereotype.Service;

import java.io.File;
import java.util.LinkedList;
import java.util.List;

@Service
public class PlanServiceImpl extends BaseServiceImpl implements PlanService {
//    @Autowired
//    ProjectService projectService;
//    @Autowired
//    RunService runService;
//
//    @Autowired
//    HistoryService historyService;

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
    public List<TstPlan> genVos(List<TstPlan> pos) {
        List<TstPlan> vos = new LinkedList<TstPlan>();

//        for (TstPlan po : pos) {
//            TstPlan vo = genVo(po);
//            vos.add(vo);
//        }
        return vos;
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
//        DetachedCriteria dc = DetachedCriteria.forClass(TstPlan.class);
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        dc.createAlias("project", "project");
//        dc.add(Restrictions.eq("project.orgId", orgId));
//
//        dc.addOrder(Order.asc("createTime"));
//
//        Page page = findPage(dc, 0, 10);
//
//        return page.getItems();

        return null;
    }

    @Override
    public List<TstPlan> listByProject(Integer projectId, String projectType) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstPlan.class);
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//
//        if (projectType.equals(TestProject.ProjectType.project.toString())) {
//            dc.add(Restrictions.eq("projectId", projectId));
//        } else {
//            dc.createAlias("project", "project");
//            dc.add(Restrictions.eq("project.parentId", projectId));
//        }
//
//        dc.addOrder(Order.asc("createTime"));
//
//        List<TstPlan> ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
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
    public TstPlan genVo(TstPlan po) {
        TstPlan vo = new TstPlan();

//        vo.setId(po.getId());
//        vo.setName(po.getName());
//        vo.setEstimate(po.getEstimate());
//        vo.setStartTime(po.getStartTime());
//        vo.setEndTime(po.getEndTime());
//        vo.setVerId(po.getVerId());
//
//        TestVer ver = po.getVerId()==null? null: (TestVer) get(TestVer.class, po.getVerId());
//        vo.setVerName(ver!=null?ver.getName():"");
//
//        vo.setDescr(po.getDescr());
//        vo.setProjectId(po.getProjectId());
//
//        TestProject project = (TestProject) get(TestProject.class, po.getProjectId());
//        vo.setProjectName(project!=null?project.getName():"");
//
//        vo.setStatus(po.getStatus().toString());
//
//        for (TestRun run : po.getRuns()) {
//            TstRun runVo = runService.genVo(run);
//            vo.getRunVos().add(runVo);
//        }

        return vo;
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

