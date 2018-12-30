package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.CaseAttachmentDao;
import com.ngtesting.platform.dao.CaseInTaskDao;
import com.ngtesting.platform.dao.CaseInTaskIssueDao;
import com.ngtesting.platform.model.*;
import com.ngtesting.platform.service.intf.CaseHistoryService;
import com.ngtesting.platform.service.intf.CaseInTaskIssueService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class CaseInTaskIssueServiceImpl extends BaseServiceImpl implements CaseInTaskIssueService {
    @Autowired
    CaseInTaskIssueDao caseInTaskIssueDao;
    @Autowired
    CaseHistoryService caseHistoryService;
    @Autowired
    CaseInTaskDao caseInTaskDao;

    @Override
    @Transactional
    public Boolean save(Integer caseInTaskId, Integer issueId, TstUser user) {
        TstCaseInTaskIssue link = new TstCaseInTaskIssue(issueId, caseInTaskId, user.getId());
        caseInTaskIssueDao.save(link);
//        caseHistoryService.saveHistory(user, Constant.EntityAct.attachment_upload, testCase, name);
        return true;
    }

    @Override
    @Transactional
    public Boolean remove(Integer caseInTaskId, Integer id, TstUser user) {
        caseInTaskIssueDao.delete(id, user.getId());
//        caseHistoryService.saveHistory(user, Constant.EntityAct.attachment_delete, testCase, attach.getName());

        return true;
    }

}

