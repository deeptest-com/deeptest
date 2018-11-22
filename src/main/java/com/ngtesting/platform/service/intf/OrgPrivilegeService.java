package com.ngtesting.platform.service.intf;

import com.ngtesting.platform.model.TstOrgPrivilegeDefine;

import java.util.List;
import java.util.Map;

public interface OrgPrivilegeService extends BaseService {
    Map<String, Boolean> listByUser(Integer userId, Integer orgId);

    List<TstOrgPrivilegeDefine> listAllOrgPrivileges();
}
