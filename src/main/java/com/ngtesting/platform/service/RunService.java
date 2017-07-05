package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestRun;
import com.ngtesting.platform.vo.TestRunVo;

import java.util.List;

public interface RunService extends BaseService {

	List<TestRun> query(Long sutieId);
	TestRunVo getById(Long caseId);
	TestRun save(JSONObject json);
	TestRun delete(Long vo, Long userId);

	List<TestRunVo> genVos(List<TestRun> pos);
	TestRunVo genVo(TestRun po);

}
