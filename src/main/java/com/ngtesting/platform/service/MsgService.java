package com.ngtesting.platform.service;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstRun;
import com.ngtesting.platform.model.TstUser;
import com.ngtesting.platform.vo.Page;

import java.util.List;

public interface MsgService extends BaseService {
    List<TstMsg> list(Integer userId, Boolean isRead);

    Page listByPage(Integer userId, String isRead, String keywords, Integer currentPage, Integer itemsPerPage);

	TstMsg getById(Integer id);
	void delete(Integer msgId, Integer userId);

	TstMsg create(TstRun run, Constant.MsgType action, TstUser optUser);
	TstMsg markReadPers(Integer id, Integer id1);
    void markAllReadPers(Integer id);

}
