package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProjectRole;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface ProjectRoleDao {
    List<TstProjectRole> list(@Param("orgId") Integer orgId,
                              @Param("keywordsParam") String keywords,
                              @Param("disabledParam") String disabled);
}
