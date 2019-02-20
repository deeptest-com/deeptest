package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueAttachmentDao;
import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.model.IsuAttachment;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueAttachmentService;
import com.ngtesting.platform.service.intf.IssueHistoryService;
import com.ngtesting.platform.service.intf.IssueService;
import com.ngtesting.platform.service.intf.MsgService;
import com.ngtesting.platform.utils.MsgUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class IssueAttachmentServiceImpl extends BaseServiceImpl implements IssueAttachmentService {
    @Autowired
    IssueAttachmentDao issueAttachmentDao;
    @Autowired
    IssueHistoryService issueHistoryService;
    @Autowired
    IssueDao issueDao;

    @Autowired
    IssueService issueService;
    @Autowired
    MsgService msgService;

    @Override
    public List<IsuAttachment> query(Integer issueId) {
        return issueAttachmentDao.query(issueId);
    }

    @Override
    public IsuAttachment get(Integer id) {
        return issueAttachmentDao.get(id);
    }

    @Override
    @Transactional
    public Boolean save(Integer issueId, String name, String path, TstUser user) {
        IsuIssue issue = issueDao.get(issueId, user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }

        IsuAttachment attach = new IsuAttachment(name, path, issueId, user.getId());
        issueAttachmentDao.save(attach);

        msgService.createForIssue(user, issue, MsgUtil.HistoryMsgTemplate.create_attament_for_issue,
                user.getNickname(), issue.getTitle(), name);

        issueHistoryService.saveHistory(user, MsgUtil.MsgAction.attachment_upload, issueId, name);
        return true;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        IsuAttachment attach = issueAttachmentDao.get(id);
        IsuIssue issue = issueDao.get(attach.getIssueId(), user.getId(), user.getDefaultPrjId());
        if (issue == null) {
            return false;
        }

        issueAttachmentDao.delete(id);

        IsuAttachment attchement = issueAttachmentDao.get(id);
        msgService.createForIssue(user, issue, MsgUtil.HistoryMsgTemplate.remove_attament_for_issue,
                user.getNickname(), issue.getTitle(), attchement.getName());

        issueHistoryService.saveHistory(user, MsgUtil.MsgAction.attachment_delete, attach.getIssueId(), attach.getName());

        return true;
    }

}

