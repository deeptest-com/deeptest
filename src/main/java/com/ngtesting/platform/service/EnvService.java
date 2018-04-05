package com.ngtesting.platform.service;

import com.alibaba.fastjson.JSONObject;
import com.ngtesting.platform.entity.TestEnv;
import com.ngtesting.platform.vo.TestEnvVo;
import com.ngtesting.platform.vo.UserVo;

import java.util.List;

public interface EnvService extends BaseService {

	List<TestEnv> list(Long projectId, String keywords, String disabled);

	TestEnvVo getById(Long id);
	TestEnv save(JSONObject json, UserVo optUser);
	TestEnv delete(Long vo, Long userId);

    boolean changeOrderPers(Long id, String act, Long orgId);

	List<TestEnvVo> listVos(Long projectId);

	List<TestEnvVo> genVos(List<TestEnv> pos);
	TestEnvVo genVo(TestEnv po);
}
