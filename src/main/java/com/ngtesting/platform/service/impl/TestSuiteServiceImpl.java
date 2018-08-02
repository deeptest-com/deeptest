package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.dao.TestSuiteDao;
import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.HistoryService;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.service.TestSuiteService;
import com.ngtesting.platform.utils.StringUtil;
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
    ProjectDao projectDao;

    @Autowired
    MsgService msgService;

    @Override
    public List listByPage(Integer projectId, String keywords, Boolean disabled) {
        List<TstSuite> suites = testSuiteDao.query(projectId, keywords, disabled);

        return suites;
    }

    @Override
    public List<TstSuite> listForImport(Integer projectId) {
        List<Integer> projectIds = projectDao.listBrotherIds(projectId);

        List<TstSuite> suites = testSuiteDao.listForImport(projectIds);

        return suites;
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
    public void delete(Integer id, Integer userId) {
        testSuiteDao.get(id);
    }

    @Override
    public TstSuite saveCases(JSONObject json, TstUser optUser) {
        Integer projectId = json.getInteger("projectId");
        Integer caseProjectId = json.getInteger("caseProjectId");
        Integer suiteId = json.getInteger("suiteId");
        List<Integer> ids = JSON.parseArray(json.getString("cases"), Integer.class) ;

        return saveCases(projectId, caseProjectId, suiteId, ids, optUser);
    }

    @Override
    public TstSuite saveCases(Integer projectId, Integer caseProjectId, Integer suiteId, List<Integer> caseIds, TstUser optUser) {
        testSuiteDao.updateSuiteProject(suiteId, projectId, caseProjectId, optUser.getId());

        String caseIdsStr = StringUtil.join(caseIds.toArray(), ",");
        testSuiteDao.addCases(suiteId, caseIdsStr);

        TstSuite suite = testSuiteDao.get(suiteId);
        Constant.MsgType action = Constant.MsgType.update_case;
        historyService.create(suite.getProjectId(), optUser, action.msg, TstHistory.TargetType.task,
                suite.getId(), suite.getName());

        return suite;
    }

//    @Override
//    public List<TstSuite> genVos(List<TstSuite> pos) {
//        for (TstSuite po : pos) {
//            genVo(po);
//        }
//        return pos;
//    }
//
//    @Override
//    public TstSuite genVo(TstSuite po) {
//        return genVo(po, false);
//    }
//    @Override
//    public TstSuite genVo(TstSuite po, Boolean withCases) {
//        TstSuite vo = new TstSuite();
//
//        vo.setId(po.getId());
//        vo.setName(po.getName());
//        vo.setEstimate(po.getEstimate());
//        vo.setDescr(po.getDescr());
//
//        vo.setProjectId(po.getProjectId());
//        TestProject prj1 = (TestProject)getDetail(TestProject.class, po.getProjectId());
//        vo.setProjectName(prj1.getName());
//
//        vo.setCaseProjectId(po.getCaseProjectId());
//        TestProject prj2 = (TestProject)getDetail(TestProject.class, po.getCaseProjectId());
//        vo.setCaseProjectName(prj2.getName());
//
//        vo.setUserId(po.getUserId());
//
//        TestUser user = (TestUser) getDetail(TestUser.class, po.getUserId());
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
//
//        return vo;
//    }

//    @Override
//    public TstCaseInSuite genCaseVo(TstCaseInSuite po) {
//        TstCaseInSuite vo = new TstCaseInSuite();
//
////        TestCase testcase = po.getTestCase();
////        BeanUtilEx.copyProperties(vo, testcase);
//
////        vo.setSteps(new LinkedList<TstCaseStep>());
////
////        List<TestCaseStep> steps = testcase.getSteps();
////        for (TestCaseStep step : steps) {
////            TstCaseStep stepVo = new TstCaseStep(
////                    step.getId(), step.getOpt(), step.getExpect(), step.getOrdr(), step.getTestCaseId());
////
////            vo.getSteps().add(stepVo);
////        }
//        return vo;
//    }

}

