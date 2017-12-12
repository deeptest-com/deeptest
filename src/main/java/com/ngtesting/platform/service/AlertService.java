package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestAlertVo;

import java.util.List;

public interface AlertService extends BaseService {

	List<TestAlert> list(Long userId);
	TestAlertVo getById(Long id);

    TestAlert create(TestRun run, TestAlert.AlertType type, Long optUserId);

    List<TestAlertVo> genVos(List<TestAlert> pos);
	TestAlertVo genVo(TestAlert po);

}
