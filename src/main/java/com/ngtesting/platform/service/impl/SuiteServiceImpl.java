package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.*;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.SuiteService;
import com.ngtesting.platform.util.BeanUtilEx;
import com.ngtesting.platform.vo.TestCaseInSuiteVo;
import com.ngtesting.platform.vo.TestSuiteVo;
import com.ngtesting.platform.vo.UserVo;
import org.apache.commons.lang.StringUtils;
import org.hibernate.criterion.DetachedCriteria;
import org.hibernate.criterion.Order;
import org.hibernate.criterion.Restrictions;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.LinkedList;
import java.util.List;

@Service
public class SuiteServiceImpl extends BaseServiceImpl implements SuiteService {

    @Autowired
    HistoryService historyService;

    @Autowired
    MsgService msgService;

    @Override
    public List<TestSuite> query(JSONObject json) {
        Long projectId = json.getLong("projectId");
        String status = json.getString("status");
        String keywords = json.getString("keywords");

        DetachedCriteria dc = DetachedCriteria.forClass(TestSuite.class);

        if (projectId != null) {
            dc.add(Restrictions.eq("projectId", projectId));
        }
        if (StringUtils.isNotEmpty(keywords)) {
            dc.add(Restrictions.like("name", "%" + keywords + "%"));
        }

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));
        dc.addOrder(Order.desc("createTime"));
        dc.addOrder(Order.asc("id"));
        List<TestSuite> ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public TestSuiteVo getById(Long caseId) {
        TestSuite po = (TestSuite) get(TestSuite.class, caseId);
        TestSuiteVo vo = genVo(po);

        return vo;
    }
    @Override
    public TestSuiteVo getById(Long caseId, Boolean withCases) {
        TestSuite po = (TestSuite) get(TestSuite.class, caseId);
        TestSuiteVo vo = genVo(po, withCases);

        return vo;
    }

    @Override
    public TestSuite save(JSONObject json, UserVo optUser) {
        Long id = json.getLong("id");

        TestSuite po;
        TestSuiteVo vo = JSON.parseObject(JSON.toJSONString(json), TestSuiteVo.class);

        Constant.MsgType action;
        if (id != null) {
            po = (TestSuite)get(TestSuite.class, id);
            action = Constant.MsgType.update;
        } else {
            po = new TestSuite();
            action = Constant.MsgType.create;
        }
        po.setName(vo.getName());
        po.setEstimate(vo.getEstimate());
        po.setDescr(vo.getDescr());
        po.setProjectId(vo.getProjectId());
        po.setUserId(optUser.getId());

        saveOrUpdate(po);

        historyService.create(po.getProjectId(), optUser, action.msg, TestHistory.TargetType.suite,
                po.getId(), po.getName());

        return po;
    }

    @Override
    public TestSuite delete(Long id, Long clientId) {
        TestSuite po = (TestSuite)get(TestSuite.class, id);
        po.setDeleted(true);
        saveOrUpdate(po);
        return po;
    }

    @Override
    public List<TestSuite> list(Long projectId, String projectType) {
        DetachedCriteria dc = DetachedCriteria.forClass(TestSuite.class);

        dc.add(Restrictions.eq("deleted", Boolean.FALSE));
        dc.add(Restrictions.eq("disabled", Boolean.FALSE));

        if (projectType.equals(TestProject.ProjectType.project.toString())) {
            dc.add(Restrictions.eq("projectId", projectId));
        } else {
            dc.createAlias("project", "project");
            dc.add(Restrictions.eq("project.parentId", projectId));
        }

        dc.addOrder(Order.asc("createTime"));

        List<TestSuite> ls = findAllByCriteria(dc);

        return ls;
    }

    @Override
    public TestSuite saveCases(JSONObject json, UserVo optUser) {
        Long planId = json.getLong("planId");
        Long suiteId = json.getLong("suiteId");
        JSONArray data = json.getJSONArray("cases");

        return saveCases(planId, suiteId, data.toArray(), optUser);
    }
    @Override
    public TestSuite saveCases(Long planId, Long suiteId, Object[] ids, UserVo optUser) {
        TestSuite suite;
        if (suiteId != null) {
            suite = (TestSuite) get(TestSuite.class, suiteId);
        } else {
            suite = new TestSuite();
            suite.setProjectId(planId);
        }

        for (TestCaseInSuite item : suite.getTestcases()) {
            getDao().delete(item);
        }

        suite.setTestcases(new LinkedList<TestCaseInSuite>());
        saveOrUpdate(suite);
        for (Object obj : ids) {
            Long id = Long.valueOf(obj.toString());
            TestCase testcase = (TestCase) get(TestCase.class, id);

            TestCaseInSuite caseInSuite = new TestCaseInSuite(suite.getProjectId(),
                    suite.getId(), id, testcase.getpId(), testcase.getLeaf());
            suite.getTestcases().add(caseInSuite);
        }
        saveOrUpdate(suite);

        Constant.MsgType action = Constant.MsgType.update_case;
        historyService.create(suite.getProjectId(), optUser, action.msg, TestHistory.TargetType.run,
                suite.getId(), suite.getName());

        return suite;
    }

    @Override
    public Long countCase(Long suiteId) {
        String hql = "select count(id) from TestCaseInSuite where isLeaf=true and suiteId=" + suiteId;
        Long count = (Long) getByHQL(hql);

        return count;
    }

    @Override
    public List<TestSuiteVo> genVos(List<TestSuite> pos) {
        List<TestSuiteVo> vos = new LinkedList<TestSuiteVo>();

        for (TestSuite po : pos) {
            TestSuiteVo vo = genVo(po);
            vos.add(vo);
        }
        return vos;
    }

    @Override
    public TestSuiteVo genVo(TestSuite po) {
        return genVo(po, false);
    }
    @Override
    public TestSuiteVo genVo(TestSuite po, Boolean withCases) {
        TestSuiteVo vo = new TestSuiteVo();

        vo.setId(po.getId());
        vo.setName(po.getName());
        vo.setEstimate(po.getEstimate());
        vo.setDescr(po.getDescr());
        vo.setProjectId(po.getProjectId());
        vo.setUserId(po.getUserId());

        TestUser user = (TestUser) get(TestUser.class, po.getUserId());
        vo.setUserName(user.getName());
        vo.setCreateTime(po.getCreateTime());
        vo.setUpdateTime(po.getUpdateTime());

        int count = 0;
        if (withCases) {
            for (TestCaseInSuite p : po.getTestcases()) {
                TestCaseInSuiteVo v = genCaseVo(p);
                vo.getTestcases().add(v);
                if (p.getLeaf()) {
                    count++;
                }
            }
        } else {
            vo.setCount(countCase(vo.getId()).intValue());
        }

        return vo;
    }

    @Override
    public TestCaseInSuiteVo genCaseVo(TestCaseInSuite po) {
        TestCaseInSuiteVo vo = new TestCaseInSuiteVo();

        TestCase testcase = po.getTestCase();
        BeanUtilEx.copyProperties(vo, testcase);

//        vo.setSteps(new LinkedList<TestCaseStepVo>());
//
//        List<TestCaseStep> steps = testcase.getSteps();
//        for (TestCaseStep step : steps) {
//            TestCaseStepVo stepVo = new TestCaseStepVo(
//                    step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());
//
//            vo.getSteps().add(stepVo);
//        }
        return vo;
    }

    @Override
    public TestSuite updatePo(TestSuiteVo vo) {
        TestSuite po = new TestSuite();
        po.setName(vo.getName());
        po.setEstimate(vo.getEstimate());
        po.setDescr(vo.getDescr());
        po.setProjectId(vo.getProjectId());
        po.setUserId(vo.getUserId());

        return po;
    }

}

