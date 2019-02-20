package com.ngtesting.platform.service.impl;

import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.dao.IssueCommentsDao;
import com.ngtesting.platform.model.IsuComments;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.IssueCommentsService;
import com.ngtesting.platform.service.intf.IssueService;
import com.ngtesting.platform.service.intf.MsgService;
import com.ngtesting.platform.utils.MsgUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class IssueCommentsServiceImpl extends BaseServiceImpl implements IssueCommentsService {
    @Autowired
    IssueCommentsDao issueCommentsDao;

    @Autowired
    IssueService issueService;
    @Autowired
    MsgService msgService;

    @Override
    public IsuComments get(Integer id) {
        return issueCommentsDao.get(id);
    }

    @Override
    @Transactional
    public IsuComments save(JSONObject json, TstUser user) {
        IsuComments vo = JSON.parseObject(JSON.toJSONString(json), IsuComments.class);

        vo.setIssueId(json.getInteger("modelId"));

        save(vo, user);
        return vo;
    }

    @Override
    @Transactional
    public IsuComments save(IsuComments vo, TstUser user) {
        vo.setUserId(user.getId());
        vo.setUserName(user.getNickname());
        vo.setUserAvatar(user.getAvatar());

        if (vo.getId() == null) {
            issueCommentsDao.save(vo);
        } else {
            issueCommentsDao.update(vo);
        }

        Integer issueId = vo.getIssueId();

        IsuIssue issue = issueService.get(issueId);
        msgService.createForIssue(user, issue, MsgUtil.HistoryMsgTemplate.create_comments_for_issue,
                user.getNickname(), issue.getTitle(), vo.getSummary());

        return vo;
    }

    @Override
    @Transactional
    public Boolean delete(Integer id, TstUser user) {
        Boolean result = issueCommentsDao.delete(id, user.getId());

        IsuComments comments = issueCommentsDao.get(id);
        IsuIssue issue = issueService.get(comments.getIssueId());
        msgService.createForIssue(user, issue, MsgUtil.HistoryMsgTemplate.remove_comments_for_issue,
                user.getNickname(), issue.getTitle(), comments.getSummary());

//        caseHistoryService.saveHistory(user, MsgUtil.MsgAction.comments_delete, testCase, comments.getContent());
        return result;
    }

}
