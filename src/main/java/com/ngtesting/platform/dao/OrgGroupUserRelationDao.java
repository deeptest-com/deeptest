package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgGroupUserRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgGroupUserRelationDao {
    List<TstOrgGroupUserRelation> query(@Param("orgId") Integer orgId,
                                        @Param("groupId") Integer groupId,
                                        @Param("userId") Integer userId);
}
