package com.ngtesting.platform.service.inf;

import java.util.Map;

public interface SysPrivilegeService extends BaseService {

	Map<String, Boolean> listByUser(Long userId);
}
