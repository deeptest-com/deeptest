package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.model.TstCaseInSuite;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.service.SuiteService;
import com.ngtesting.platform.vo.Page;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class SuiteServiceImpl extends BaseServiceImpl implements SuiteService {

    @Autowired
    HistoryService historyService;
    @Autowired
    ProjectService projectService;

    @Autowired
    MsgService msgService;

    @Override
    public Page page(Integer projectId, String keywords, Integer currentPage, Integer itemsPerPage) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstSuite.class);
//
//        if (projectId != null) {
//            List<Integer> ids = projectService.listBrotherIds(projectId);
//            dc.add(Restrictions.in("projectId", ids));
//        }
//        if (StringUtils.isNotEmpty(keywords)) {
//            dc.add(Restrictions.like("name", "%" + keywords + "%"));
//        }
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("caseProjectId"));
//        dc.addOrder(Order.asc("id"));
//        Page page = findPage(dc, currentPage * itemsPerPage, itemsPerPage);
//
//        return page;

        return null;
    }

    @Override
    public List<TstSuite> query(Integer projectId, String keywords) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstSuite.class);
//
//        if (projectId != null) {
//            List<Integer> ids = projectService.listBrotherIds(projectId);
//            dc.add(Restrictions.in("projectId", ids));
//        }
//        if (StringUtils.isNotEmpty(keywords)) {
//            dc.add(Restrictions.like("name", "%" + keywords + "%"));
//        }
//
//        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
//        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
//        dc.addOrder(Order.asc("caseProjectId"));
//        dc.addOrder(Order.asc("id"));
//        List<TstSuite> ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
    }

    @Override
    public TstSuite getById(Integer caseId) {
//        TstSuite po = (TstSuite) get(TstSuite.class, caseId);
//        TstSuite vo = genVo(po);
//
//        return vo;

        return null;
    }
    @Override
    public TstSuite getById(Integer caseId, Boolean withCases) {
//        TstSuite po = (TstSuite) get(TstSuite.class, caseId);
//        TstSuite vo = genVo(po, withCases);
//
//        return vo;

        return null;
    }

    @Override
    public TstSuite save(JSONObject json, TstUser optUser) {
//        Integer id = json.getInteger("id");
//
//        TstSuite po;
//        TstSuite vo = JSON.parseObject(JSON.toJSONString(json), TstSuite.class);
//
//        Constant.MsgType action;
//        if (id != null) {
//            po = (TstSuite)get(TstSuite.class, id);
//            action = Constant.MsgType.update;
//        } else {
//            po = new TstSuite();
//            action = Constant.MsgType.create;
//        }
//        po.setName(vo.getName());
//        po.setEstimate(vo.getEstimate());
//        po.setDescr(vo.getDescr());
//        po.setProjectId(vo.getProjectId());
//        po.setCaseProjectId(vo.getCaseProjectId());
//        po.setUserId(optUser.getId());
//
//        saveOrUpdate(po);
//
//        historyService.create(po.getProjectId(), optUser, action.msg, TestHistory.TargetType.suite,
//                po.getId(), po.getName());
//
//        return po;

        return null;
    }

    @Override
    public TstSuite delete(Integer id, Integer clientId) {
//        TstSuite po = (TstSuite)get(TstSuite.class, id);
//        po.setDeleted(true);
//        saveOrUpdate(po);
//        return po;

        return null;
    }

    @Override
    public List<TstSuite> list(Integer projectId, String projectType) {
//        DetachedCriteria dc = DetachedCriteria.forClass(TstSuite.class);
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
//        List<TstSuite> ls = findAllByCriteria(dc);
//
//        return ls;

        return null;
    }

    @Override
    public TstSuite saveCases(JSONObject json, TstUser optUser) {
        Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
        Integer suiteId = json.getInteger("suiteId");
        JSONArray data = json.getJSONArray("cases");

        return saveCases(projectId, caseProjectId, suiteId, data.toArray(), optUser);
    }
    @Override
    public TstSuite saveCases(Integer projectId, Integer caseProjectId, Integer suiteId, Object[] ids, TstUser optUser) {
        TstSuite suite = null;
//        if (suiteId != null) {
//            suite = (TstSuite) get(TstSuite.class, suiteId);
//        } else {
//            suite = new TstSuite();
//        }
//        suite.setProjectId(projectId);
//        suite.setCaseProjectId(caseProjectId);
//
//        suite.setTestcases(new LinkedList<TstCaseInSuite>());
//        saveOrUpdate(suite);
//
//        List<Integer> caseIds = new LinkedList<>();
//        for (Object obj : ids) {
//            Integer id = Integer.valueOf(obj.toString());
//            caseIds.add(id);
//        }
//        addCasesPers(suite.getId(), caseIds);
//
//        Constant.MsgType action = Constant.MsgType.update_case;
//        historyService.create(suite.getProjectId(), optUser, action.msg, TestHistory.TargetType.run,
//                suite.getId(), suite.getName());

        return suite;
    }

    @Override
    public void addCasesPers(Integer suiteId, List<Integer> caseIds) {
//        String ids = StringUtils.join(caseIds.toArray(), ",");
//        getDao().querySql("{call add_cases_to_suite(?,?)}", suiteId, ids);
    }

    @Override
    public Integer countCase(Integer suiteId) {
//        String hql = "select count(id) from TstCaseInSuite where isLeaf=true and suiteId=" + suiteId;
//        Integer count = (Integer) getByHQL(hql);
//
//        return count;

        return null;
    }

    @Override
    public List<TstSuite> genVos(List<TstSuite> pos) {
//        List<TstSuite> vos = new LinkedList<TstSuite>();
//
//        for (TstSuite po : pos) {
//            TstSuite vo = genVo(po);
//            vos.add(vo);
//        }
//        return vos;

        return null;
    }

    @Override
    public TstSuite genVo(TstSuite po) {
        return genVo(po, false);
    }
    @Override
    public TstSuite genVo(TstSuite po, Boolean withCases) {
        TstSuite vo = new TstSuite();

//        vo.setId(po.getId());
//        vo.setName(po.getName());
//        vo.setEstimate(po.getEstimate());
//        vo.setDescr(po.getDescr());
//
//        vo.setProjectId(po.getProjectId());
//        TestProject prj1 = (TestProject)get(TestProject.class, po.getProjectId());
//        vo.setProjectName(prj1.getName());
//
//        vo.setCaseProjectId(po.getCaseProjectId());
//        TestProject prj2 = (TestProject)get(TestProject.class, po.getCaseProjectId());
//        vo.setCaseProjectName(prj2.getName());
//
//        vo.setUserId(po.getUserId());
//
//        TestUser user = (TestUser) get(TestUser.class, po.getUserId());
//        vo.setUserName(user.getName());
//        vo.setCreateTime(po.getCreateTime());
//        vo.setUpdateTime(po.getUpdateTime());
//
//        int count = 0;
//        if (withCases) {
//            for (TstCaseInSuite p : po.getTestcases()) {
//                TstCaseInSuite v = genCaseVo(p);
//                vo.getTestcases().add(v);
//                if (p.getLeaf()) {
//                    count++;
//                }
//            }
//        } else {
//            vo.setCount(countCase(vo.getId()).intValue());
//        }

        return vo;
    }

    @Override
    public TstCaseInSuite genCaseVo(TstCaseInSuite po) {
        TstCaseInSuite vo = new TstCaseInSuite();

//        TestCase testcase = po.getTestCase();
//        BeanUtilEx.copyProperties(vo, testcase);

//        vo.setSteps(new LinkedList<TstCaseStep>());
//
//        List<TestCaseStep> steps = testcase.getSteps();
//        for (TestCaseStep step : steps) {
//            TstCaseStep stepVo = new TstCaseStep(
//                    step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());
//
//            vo.getSteps().add(stepVo);
//        }
        return vo;
    }

    @Override
    public TstSuite updatePo(TstSuite vo) {
        TstSuite po = new TstSuite();
//        po.setName(vo.getName());
//        po.setEstimate(vo.getEstimate());
//        po.setDescr(vo.getDescr());
//        po.setProjectId(vo.getProjectId());
//        po.setUserId(vo.getUserId());

        return po;
    }

}

