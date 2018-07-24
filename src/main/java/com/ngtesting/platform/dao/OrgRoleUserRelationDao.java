package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgRoleUserRelation;

import java.util.List;

public interface OrgRoleUserRelationDao {
    List<TstOrgRoleUserRelation> query(Integer orgId, Integer orgRoleId, Integer userId);

    void removeAllRolesForUser(Integer orgId, Integer userId);

    void removeAllUsersForRole(Integer orgId, Integer roleId);

    void saveRelations(List<TstOrgRoleUserRelation> selectedList);
}
