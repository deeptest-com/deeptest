package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.entity.TestMsg;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.entity.TestUser;
import com.ngtesting.platform.vo.TestMsgVo;

import java.util.List;

public interface MsgService extends BaseService {

	List<TestMsg> list(Long userId);
	TestMsgVo getById(Long id);

	TestMsg create(TestRun run, TestAlert.AlertAction action, TestUser optUser);

	List<TestMsgVo> genVos(List<TestMsg> pos);
	TestMsgVo genVo(TestMsg po);
}
