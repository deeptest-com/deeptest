package com.ngtesting.platform.service;

import com.ngtesting.platform.entity.AiTestEnv;
import com.ngtesting.platform.vo.AiTestEnvVo;

import java.util.List;

public interface AiTestEnvService extends BaseService {

	List<AiTestEnvVo> listTestEnvVo(Long projectId);

	List<AiTestEnvVo> genVos(List<AiTestEnv> pos);

	AiTestEnvVo genVo(AiTestEnv po);
}
