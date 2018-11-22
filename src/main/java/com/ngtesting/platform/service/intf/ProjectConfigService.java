package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstProject;

public interface ProjectConfigService extends BaseService {
	TstProject get(Integer id);
}
