package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.IssueOptDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.IsuComments;
import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.service.intf.IssueOptService;
import com.ngtesting.platform.service.intf.IssueStatusService;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IssueOptServiceImpl extends BaseServiceImpl implements IssueOptService {
    Log logger = LogFactory.getLog(IssueOptServiceImpl.class);

    @Autowired
    IssueOptDao issueOptDao;

    @Autowired
    UserDao userDao;

    @Autowired
    IssueHistoryService issueHistoryService;

    @Autowired
    IssueCommentsService issueCommentsService;

    @Autowired
    IssueStatusService issueStatusService;

    @Override
    public void statusTran(Integer id, Integer dictStatusId, String dictStatusName, TstUser user) {
        IsuStatus status = issueStatusService.get(dictStatusId, user.getDefaultOrgId());
        if (status.getFinalVal()) {

        }

        issueOptDao.statusTran(id, dictStatusId, status.getFinalVal(), user.getDefaultPrjId());

        issueHistoryService.saveHistory(user, Constant.EntityAct.changeStatus, id, dictStatusName);
    }

    @Override
    public void assign(Integer id, Integer userId, String content, TstUser user) {
        IsuComments po = new IsuComments(id, "修改经办人", content);
        issueCommentsService.save(po, user);

        issueOptDao.assign(id, userId, user.getDefaultPrjId());

        TstUser u = userDao.get(userId);

        issueHistoryService.saveHistory(user, Constant.EntityAct.assign, id, u.getNickname());
    }

}

