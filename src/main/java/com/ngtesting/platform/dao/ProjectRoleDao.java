package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstProjectRole;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface ProjectRoleDao {
    List<TstProjectRole> query(@Param("orgId") Integer orgId,
                              @Param("keywordsParam") String keywords,
                              @Param("disabledParam") Boolean disabled);

    TstProjectRole getRoleByCode(@Param("orgId") Integer orgId,
                          @Param("roleCode") String roleCode);

    TstProjectRole get(@Param("id") Integer roleId,
                       @Param("orgId") Integer orgId);

    Integer save(TstProjectRole vo);
    Integer update(TstProjectRole vo);
    Integer delete(@Param("id") Integer roleId,
                @Param("orgId") Integer orgId);
}
