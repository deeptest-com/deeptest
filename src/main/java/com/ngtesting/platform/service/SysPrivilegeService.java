package com.ngtesting.platform.service;

import java.util.Map;

public interface SysPrivilegeService extends BaseService {

	Map<String, Boolean> listByUser(Long userId);
}
