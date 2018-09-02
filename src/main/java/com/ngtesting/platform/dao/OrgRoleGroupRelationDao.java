package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgRoleGroupRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgRoleGroupRelationDao {
    List<TstOrgRoleGroupRelation> query(@Param("orgId") Integer orgId,
                                        @Param("roleId") Integer roleId,
                                        @Param("groupId") Integer groupId);

    void removeAllRolesForGroup(@Param("orgId") Integer orgId, @Param("groupId") Integer groupId);

    void removeAllGroupsForRole(@Param("orgId") Integer orgId, @Param("roleId") Integer roleId);

    void saveRelations(@Param("list") List<TstOrgRoleGroupRelation> selectedList);
}

