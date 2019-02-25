package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.IssueOptDao;
import com.ngtesting.platform.dao.UserDao;
import com.ngtesting.platform.model.IsuComments;
import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.*;
import com.ngtesting.platform.utils.MsgUtil;
import org.apache.commons.logging.Log;
import org.apache.commons.logging.LogFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class IssueOptServiceImpl extends BaseServiceImpl implements IssueOptService {
    Log logger = LogFactory.getLog(IssueOptServiceImpl.class);

    @Autowired
    IssueOptService issueOptService;
    @Autowired
    IssueService issueService;

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

        issueHistoryService.saveHistory(user, MsgUtil.MsgAction.changeStatus, id, dictStatusName);
    }

    @Override
    public void assign(Integer id, Integer userId, String content, TstUser user) {
        IsuComments po = new IsuComments(id, "修改经办人", content);
        issueCommentsService.save(po, user);

        issueOptDao.assign(id, userId, user.getDefaultPrjId());

        TstUser u = userDao.get(userId);

        issueHistoryService.saveHistory(user, MsgUtil.MsgAction.assign, id, u.getNickname());
    }

    @Override
    public void updateThenStatusTran(JSONObject json, TstUser user) {
        JSONObject issue = json.getJSONObject("issue");
        Integer id = issue.getInteger("id");

        JSONObject tran = json.getJSONObject("tran");
        Integer pageId = tran.getInteger("actionPageId");
        Integer dictStatusId = tran.getInteger("dictStatusId");
        String dictStatusName = tran.getString("dictStatusName");

        issueService.update(issue, pageId, user);
        issueOptService.statusTran(id, dictStatusId, dictStatusName, user);
    }

}

