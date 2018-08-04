package com.ngtesting.platform.service.impl;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.dao.MsgDao;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.service.MsgService;
import com.ngtesting.platform.utils.StringUtil;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class MsgServiceImpl extends BaseServiceImpl implements MsgService {
    @Autowired
    private MsgDao msgDao;

	@Override
	public List<TstMsg> list(Integer userId, Boolean isRead, String keywords) {
        List<TstMsg> ls = msgDao.query(userId, isRead, keywords);

        return ls;
	}

    @Override
    public TstMsg getById(Integer id) {
//        TstMsg po = (TstMsg) getDetail(TstMsg.class, id);
//        TstMsg vo = genVo(po);
//
//        return vo;
        return null;
    }

    @Override
    public void delete(Integer msgId, Integer userId) {
//        TstMsg po = (TstMsg) getDetail(TstMsg.class, msgId);
//        po.setDeleted(true);
//        saveOrUpdate(po);
    }

    @Override
    public TstMsg create(TstTask task, Constant.MsgType action, TstUser optUser) {
        TstMsg msg = new TstMsg();

        msg.setTitle("用户" + StringUtil.highlightDict(optUser.getNickname()) + action.msg
                + "测试集" + StringUtil.highlightDict(task.getName()));

        msg.setUserId(task.getUserId());
        msgDao.create(msg);

        return msg;
    }

    @Override
    public TstMsg markReadPers(Integer id) {
        msgDao.markRead(id);
        TstMsg po = msgDao.get(id);
        return po;
    }

    @Override
    public void markAllReadPers(Integer userId) {
        msgDao.markAllRead(userId);
    }

}

