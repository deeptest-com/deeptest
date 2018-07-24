package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgRoleUserRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgRoleUserRelationDao {
    List<TstOrgRoleUserRelation> query(@Param("orgId") Integer orgId,
                                       @Param("roleId") Integer roleId,
                                       @Param("userId") Integer userId);

    void removeAllRolesForUser(@Param("orgId") Integer orgId, @Param("userId") Integer userId);

    void removeAllUsersForRole(@Param("orgId") Integer orgId, @Param("roleId") Integer roleId);

    void saveRelations(@Param("list") List<TstOrgRoleUserRelation> selectedList);
}
