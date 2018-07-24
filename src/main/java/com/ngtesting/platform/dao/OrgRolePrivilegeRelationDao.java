package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgRolePrivilegeRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgRolePrivilegeRelationDao {
    List<TstOrgRolePrivilegeRelation> query(@Param("orgId") Integer orgId,
                                            @Param("roleId") Integer roleId,
                                            @Param("privilegeId") Integer privilegeId);

    void removeAllPrivilegesForRole(@Param("orgId") Integer orgId, @Param("roleId") Integer roleId);
    void removeAllRolesForPrivilege(@Param("orgId") Integer orgId, @Param("privilegeId") Integer privilegeId);
    void saveRelations(@Param("list") List<TstOrgRolePrivilegeRelation> selectedList);

}
