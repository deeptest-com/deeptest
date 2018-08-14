package com.ngtesting.platform.service;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface MsgService extends BaseService {
    List<TstMsg> list(Integer userId, Boolean isRead, String keywords);

	void delete(Integer msgId, Integer userId);

    void create(TstTask run, Constant.MsgType action, TstUser optUser);

    void markRead(Integer id, Integer userId);
    void markAllRead(Integer id);

}
