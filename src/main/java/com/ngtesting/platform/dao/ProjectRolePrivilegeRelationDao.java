package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProjectPrivilegeDefine;
import com.ngtesting.platform.model.TstProjectRolePriviledgeRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface ProjectRolePrivilegeRelationDao {
    List<TstProjectRolePriviledgeRelation> listProjectRolePrivileges(
            @Param("projectRoleId") Integer projectRoleId);

    void removeAllPrivilegsForRole(@Param("projectRoleId") Integer projectRoleId);

    void saveRelations(@Param("orgId") Integer orgId,
                       @Param("projectRoleId") Integer projectRoleId,
                       @Param("list") List<TstProjectPrivilegeDefine> selectedList);
}
