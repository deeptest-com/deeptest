package com.ngtesting.platform.service;

import com.ngtesting.platform.model.TstProject;

public interface AuthService extends BaseService {

    Boolean noProjectAndProjectGroupPrivilege(Integer userId, TstProject project);

    Boolean hasOrgAdminPrivilege(Integer userId, Integer orgId);
}
