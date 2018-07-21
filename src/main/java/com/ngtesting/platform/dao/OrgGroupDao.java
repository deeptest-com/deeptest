package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgGroup;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgGroupDao {
    List<TstOrgGroup> search(@Param("orgId") Integer orgId,
                             @Param("keywords") String keywords,
                             @Param("exceptIds") String exceptIds);
}
