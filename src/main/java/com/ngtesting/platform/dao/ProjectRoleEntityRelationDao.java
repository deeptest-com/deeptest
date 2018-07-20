package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstMsg;
import com.ngtesting.platform.model.TstProjectRoleEntityRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface ProjectRoleEntityRelationDao {
    List<TstMsg> query(@Param("userId") Integer userId);

    List<TstProjectRoleEntityRelation> listByProject(@Param("projectId") Integer projectId);
}
