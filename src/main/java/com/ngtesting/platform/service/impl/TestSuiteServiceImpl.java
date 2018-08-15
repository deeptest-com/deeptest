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
    public TstSuite get(Integer id, Integer projectId) {
        TstSuite po = testSuiteDao.get(id, projectId);

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
    public void delete(Integer id, Integer projectId) {
        testSuiteDao.delete(id, projectId);
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

        TstSuite suite = testSuiteDao.get(suiteId, projectId);
        Constant.MsgType action = Constant.MsgType.update_case;
        historyService.create(suite.getProjectId(), optUser, action.msg, TstHistory.TargetType.task,
                suite.getId(), suite.getName());

        return suite;
    }

}

