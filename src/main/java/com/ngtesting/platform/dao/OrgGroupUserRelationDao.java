package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgGroupUserRelationDao {
    List<TstOrgGroupUserRelation> query(@Param("orgId") Integer orgId,
                                        @Param("groupId") Integer groupId,
                                        @Param("userId") Integer userId);

    void removeAllGroupsForUser(@Param("orgId") Integer orgId, @Param("userId") Integer userId);
    void removeAllUsersForGroup(@Param("orgId") Integer orgId, @Param("groupId") Integer groupId);
    void saveRelations(@Param("list") List<TstOrgGroupUserRelation> orgGroupUserRelation);
}
