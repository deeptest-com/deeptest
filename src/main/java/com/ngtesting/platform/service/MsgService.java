package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.entity.TestMsg;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestMsgVo;

import java.util.List;

public interface MsgService extends BaseService {

	List<TestMsg> list(Long userId);
	TestMsgVo getById(Long id);

	TestMsg create(TestRun run, TestAlert.AlertType type, Long optUserId);

	List<TestMsgVo> genVos(List<TestMsg> pos);
	TestMsgVo genVo(TestMsg po);
}
