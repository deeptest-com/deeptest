package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.ProjectDao;
import com.ngtesting.platform.dao.TestSuiteDao;
import com.ngtesting.platform.model.TstHistory;
import com.ngtesting.platform.model.TstSuite;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.MsgService;
import com.ngtesting.platform.service.intf.ProjectHistoryService;
import com.ngtesting.platform.service.intf.TestSuiteService;
import com.ngtesting.platform.utils.MsgUtil;
import com.ngtesting.platform.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class TestSuiteServiceImpl extends BaseServiceImpl implements TestSuiteService {
    @Autowired
    TestSuiteDao testSuiteDao;
    @Autowired
    ProjectHistoryService historyService;
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
    @Transactional
    public TstSuite save(JSONObject json, TstUser user) {
        TstSuite vo = JSON.parseObject(JSON.toJSONString(json), TstSuite.class);
        vo.setUserId(user.getId());
        vo.setProjectId(user.getDefaultPrjId());

        MsgUtil.MsgAction action;
        if (vo.getId() == null) {
            action = MsgUtil.MsgAction.create;

            testSuiteDao.save(vo);
        } else {
            action = MsgUtil.MsgAction.update;

            Integer count = testSuiteDao.update(vo);
            if (count == 0) {
                return null;
            }
        }

        historyService.create(vo.getProjectId(), user, action.msg, TstHistory.TargetType.suite,
                vo.getId(), vo.getName());

        return vo;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, Integer projectId) {
        Integer count = testSuiteDao.delete(id, projectId);
        return count > 0;
    }

    @Override
    @Transactional
    public TstSuite saveCases(Integer projectId, Integer caseProjectId, Integer suiteId,
                              List<Integer> caseIds, TstUser user) {
        testSuiteDao.updateSuiteProject(suiteId, projectId, caseProjectId, user.getId());

        String caseIdsStr = StringUtil.join(caseIds.toArray(), ",");
        testSuiteDao.addCases(caseIdsStr, suiteId);

        TstSuite suite = testSuiteDao.get(suiteId, projectId);

        historyService.create(suite.getProjectId(), user,
                MsgUtil.MsgAction.update.msg, TstHistory.TargetType.suite,
                suite.getId(), suite.getName());

        return suite;
    }

}

