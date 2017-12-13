package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.TestAlert;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestAlertVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface AlertService extends BaseService {

	List<TestAlert> list();
	TestAlertVo getById(Long id);

	void create(TestRun run, UserVo optUser);

	List<TestAlertVo> genVos(List<TestAlert> pos);
	TestAlertVo genVo(TestAlert po);

}
