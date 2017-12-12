package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.vo.TestAlertVo;

import java.util.List;

public interface AlertService extends BaseService {

	List<TestAlert> list(Long userId);
	TestAlertVo getById(Long id);

	TestAlert save(JSONObject json);
	List<TestAlertVo> genVos(List<TestAlert> pos);
	TestAlertVo genVo(TestAlert po);

}
