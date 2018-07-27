package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.TestSuiteDao;
import com.ngtesting.platform.model.TstCaseInSuite;
import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.ProjectService;
import com.ngtesting.platform.service.TestSuiteService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class TestSuiteServiceImpl extends BaseServiceImpl implements TestSuiteService {
    @Autowired
    TestSuiteDao testSuiteDao;
    @Autowired
    HistoryService historyService;
    @Autowired
    ProjectService projectService;

    @Autowired
    MsgService msgService;

    @Override
    public List listByPage(Integer projectId, String keywords, String disabled) {
        List<TstSuite> groups = testSuiteDao.query(projectId, keywords, disabled);

        return groups;
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
    public TstSuite get(Integer id) {
        TstSuite po = testSuiteDao.get(id);

        return po;
    }
    @Override
    public TstSuite getWithCases(Integer id) {
        TstSuite po = testSuiteDao.getWithCases(id);

        return po;
    }

    @Override
    public TstSuite save(JSONObject json, TstUser optUser) {
        TstSuite vo = JSON.parseObject(JSON.toJSONString(json), TstSuite.class);
        vo.setUserId(optUser.getId());

        Constant.MsgType action;
        if (vo.getId() != null) {
            action = Constant.MsgType.update;

            testSuiteDao.update(vo);
        } else {
            action = Constant.MsgType.create;

            testSuiteDao.save(vo);
        }

        historyService.create(vo.getProjectId(), optUser, action.msg, TstHistory.TargetType.suite,
                vo.getId(), vo.getName());

        return vo;
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
//        suite.setTestCases(new LinkedList<TstCaseInSuite>());
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
//            for (TstCaseInSuite p : po.getTestCases()) {
//                TstCaseInSuite v = genCaseVo(p);
//                vo.getTestCases().add(v);
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

