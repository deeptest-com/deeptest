package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.dao.IssueDao;
import com.ngtesting.platform.dao.MsgDao;
import com.ngtesting.platform.dao.TestTaskDao;
import com.ngtesting.platform.model.IsuIssue;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.intf.MsgService;
import com.ngtesting.platform.utils.MsgUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.text.MessageFormat;
import java.util.List;

@Service
public class MsgServiceImpl extends BaseServiceImpl implements MsgService {
    @Autowired
    private MsgDao msgDao;
    @Autowired
    private TestTaskDao taskDao;
    @Autowired
    private IssueDao issueDao;

	@Override
	public List<TstMsg> list(Integer userId, Boolean isRead, String keywords) {
        List<TstMsg> ls = msgDao.query(userId, isRead, keywords);

        return ls;
	}

    @Override
    @Transactional // 针对每个经办人
    public void createForTask(TstUser optUser, TstTask task, MsgUtil.HistoryMsgTemplate template, Object... params) {
        String content = MessageFormat.format(template.toString(), params);

        List<Integer> ids = taskDao.listAssigneeIds(task.getId());
        for (Integer id: ids) {
            TstMsg msg = new TstMsg();

            msg.setTitle(content);
            msg.setUserId(optUser.getId());
            msg.setAssigneeId(id);
            msgDao.create(msg);
        }
    }

    @Override
    @Transactional // 针对经办人和监听者
    public void createForIssue(TstUser optUser, IsuIssue issue, MsgUtil.HistoryMsgTemplate template, Object... params) {
        String content = MessageFormat.format(template.toString(), params);

        List<Integer> ids = issueDao.listAssigneeAndWatcherIds(issue.getId());
        for (Integer id: ids) {
            TstMsg msg = new TstMsg();

            msg.setTitle(content);
            msg.setUserId(optUser.getId());
            msg.setAssigneeId(id);
            msgDao.create(msg);
        }
    }

    @Override
    @Transactional
    public Boolean delete(Integer msgId, Integer userId) {
        Integer count = msgDao.delete(msgId, userId);
        return count > 0;
    }

    @Override
    @Transactional
    public Boolean markRead(Integer id, Integer userId) {
	    Integer count = msgDao.markRead(id, userId);
        return count > 0;
    }

    @Override
    @Transactional
    public void markAllRead(Integer userId) {
        msgDao.markAllRead(userId);
    }

}

