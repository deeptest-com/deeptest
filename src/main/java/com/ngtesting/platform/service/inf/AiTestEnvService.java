package com.ngtesting.platform.service.inf;

import com.ngtesting.platform.model.AiTestEnv;

import java.util.List;

public interface AiTestEnvService extends BaseService {

	List<AiTestEnv> listTestEnvVo(Long projectId);
}
