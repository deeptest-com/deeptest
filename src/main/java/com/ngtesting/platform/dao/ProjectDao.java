package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProject;
import org.apache.ibatis.annotations.Param;

import java.util.List;
import java.util.Map;

public interface ProjectDao {
    List<TstProject> query(@Param("orgId") Integer orgId, @Param("keywords") String keywords,
                           @Param("disabled") Boolean disabled);

    List<Map<String, String>> getProjectPrivilegeByOrgForUser(@Param("userId") Integer userId, @Param("orgId") Integer orgId);
}
