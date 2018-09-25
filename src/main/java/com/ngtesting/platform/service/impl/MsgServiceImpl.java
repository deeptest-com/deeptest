package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.MsgDao;
import com.ngtesting.platform.dao.TestTaskDao;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@Service
public class MsgServiceImpl extends BaseServiceImpl implements MsgService {
    @Autowired
    private MsgDao msgDao;
    @Autowired
    private TestTaskDao taskDao;

	@Override
	public List<TstMsg> list(Integer userId, Boolean isRead, String keywords) {
        List<TstMsg> ls = msgDao.query(userId, isRead, keywords);

        return ls;
	}

    @Override
    @Transactional
    public void create(TstTask task, Constant.MsgType action, TstUser optUser) {
        List<Integer> ids = taskDao.listAssigneeIds(task.getId());
        for (Integer id: ids) {
            TstMsg msg = new TstMsg();

            msg.setTitle("用户" + StringUtil.highlightDict(optUser.getNickname()) + action.msg
                    + "任务" + StringUtil.highlightDict(task.getName()));

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

