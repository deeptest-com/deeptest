package com.ngtesting.platform.service;

import com.ngtesting.platform.config.Constant;
import com.ngtesting.platform.entity.TestMsg;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.Page;
import com.ngtesting.platform.vo.TestMsgVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface MsgService extends BaseService {
    List<TestMsgVo> list(Long userId, Boolean isRead);

    Page listByPage(Long userId, String isRead, String keywords, Integer currentPage, Integer itemsPerPage);

	TestMsgVo getById(Long id);
	void delete(Long msgId, Long userId);

	TestMsg create(TestRun run, Constant.MsgType action, UserVo optUser);
    TestMsg markRead(Long id, Long id1);
    void markAllReadPers(Long id);

	List<TestMsgVo> genVos(List<TestMsg> pos);
	TestMsgVo genVo(TestMsg po);

}
