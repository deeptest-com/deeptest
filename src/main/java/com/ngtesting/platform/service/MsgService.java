package com.ngtesting.platform.service;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstTask;
import com.ngtesting.platform.model.TstUser;

import java.util.List;

public interface MsgService extends BaseService {
    List<TstMsg> list(Integer userId, Boolean isRead, String keywords);

	TstMsg getById(Integer id);
	void delete(Integer msgId, Integer userId);

	TstMsg create(TstTask run, Constant.MsgType action, TstUser optUser);

    TstMsg markReadPers(Integer id);
    void markAllReadPers(Integer id);

}
