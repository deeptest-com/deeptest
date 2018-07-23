package com.ngtesting.platform.dao;

import org.apache.ibatis.annotations.Param;

public interface OrgRoleUserDao {
    boolean userInOrg(@Param("orgId") Integer orgId, @Param("userId") Integer userId);
}
