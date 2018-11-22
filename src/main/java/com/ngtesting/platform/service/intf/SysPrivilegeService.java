package com.ngtesting.platform.service.intf;

import java.util.Map;

public interface SysPrivilegeService extends BaseService {

	Map<String, Boolean> listByUser(Integer userId);
}
