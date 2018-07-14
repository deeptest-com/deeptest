package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgPrivilegeDefine;

import java.util.List;

public interface OrgRoleDao {
    List<TstOrgPrivilegeDefine> queryByUser(Integer userId, Integer orgId);
}
