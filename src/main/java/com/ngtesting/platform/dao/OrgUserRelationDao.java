package com.ngtesting.platform.dao;

import org.apache.ibatis.annotations.Param;

public interface OrgUserRelationDao {
    Integer userInOrg(@Param("userId") Integer userId, @Param("orgId") Integer orgId);

    void addUserToOrg(@Param("userId") Integer userId, @Param("orgId") Integer orgId);

    void addUserToDefaultOrgGroup(@Param("userId") Integer userId,
                                  @Param("orgId") Integer orgId);
}
