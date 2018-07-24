package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgRolePrivilegeRelation;

import java.util.List;

public interface OrgRolePrivilegeRelationDao {
    List<TstOrgRolePrivilegeRelation> query(Integer orgId, Integer orgRoleId, Integer privilegeId);

    void removeAllPrivilegesForRole(Integer orgId, Integer roleId);
    void removeAllRolesForPrivilege(Integer orgId, Integer privilegeId);
    void saveRelations(List<TstOrgRolePrivilegeRelation> selectedList);

}
