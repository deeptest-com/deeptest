package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestMsg;
import com.ngtesting.platform.vo.TestMsgVo;

import java.util.List;

public interface MsgService extends BaseService {

	List<TestMsg> list(Long userId);
	TestMsgVo getById(Long id);

	TestMsg save(JSONObject json);
	List<TestMsgVo> genVos(List<TestMsg> pos);
	TestMsgVo genVo(TestMsg po);
}
