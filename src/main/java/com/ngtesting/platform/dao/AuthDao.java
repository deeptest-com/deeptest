package com.ngtesting.platform.dao;

import org.apache.ibatis.annotations.Param;

public interface AuthDao {
    Boolean userNotInProject(@Param("userId") Integer userId, @Param("projectId") Integer projectId);
}
