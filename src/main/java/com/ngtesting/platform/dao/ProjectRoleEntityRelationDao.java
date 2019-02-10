package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstProjectRoleEntityRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface ProjectRoleEntityRelationDao {
    List<TstMsg> query(@Param("userId") Integer userId);

    List<TstProjectRoleEntityRelation> listByProject(@Param("projectId") Integer projectId);

    List<Integer> listIdsByUserAndProject(@Param("userId") Integer userId,
                                          @Param("projectId") Integer projectId);

    void addRole(@Param("orgId") Integer orgId,
                 @Param("projectId") Integer projectId,
                 @Param("projectRoleId") Integer projectRoleId,
                 @Param("entityId") Integer entityId,
                 @Param("type") String type);
    void changeRole(@Param("projectId") Integer projectId,
                    @Param("projectRoleId") Integer projectRoleId,
                    @Param("entityId") Integer entityId);

    void remove(@Param("projectId") Integer projectId,
                @Param("type") String type,
                @Param("entityId") Integer entityId);
}
