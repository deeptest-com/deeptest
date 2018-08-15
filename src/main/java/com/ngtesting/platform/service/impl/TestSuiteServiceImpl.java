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
    public TstSuite save(JSONObject json, TstUser user) {
        TstSuite vo = JSON.parseObject(JSON.toJSONString(json), TstSuite.class);
        vo.setUserId(user.getId());
        vo.setProjectId(user.getDefaultPrjId());

        Constant.MsgType action;
        if (vo.getId() == null) {
            action = Constant.MsgType.create;

            testSuiteDao.save(vo);
        } else {
            action = Constant.MsgType.update;

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
    public Boolean delete(Integer id, Integer projectId) {
        Integer count = testSuiteDao.delete(id, projectId);
        return count > 0;
    }

    @Override
    public TstSuite saveCases(Integer projectId, Integer caseProjectId, Integer suiteId,
                              List<Integer> caseIds, TstUser user) {
        testSuiteDao.updateSuiteProject(suiteId, projectId, caseProjectId, user.getId());

        String caseIdsStr = StringUtil.join(caseIds.toArray(), ",");
        testSuiteDao.addCases(suiteId, caseIdsStr);

        TstSuite suite = testSuiteDao.get(suiteId, projectId);
        Constant.MsgType action = Constant.MsgType.update_case;
        historyService.create(suite.getProjectId(), user, action.msg, TstHistory.TargetType.task,
                suite.getId(), suite.getName());

        return suite;
    }

}

