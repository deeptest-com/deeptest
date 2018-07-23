package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrgGroup;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgGroupDao {
    List<TstOrgGroup> query(@Param("orgId") Integer orgId,
                        @Param("keywords") String keywords,
                        @Param("disabled") String disabled);

    List<TstOrgGroup> search(@Param("orgId") Integer orgId,
                             @Param("keywords") String keywords,
                             @Param("exceptIds") String exceptIds);
    List<TstOrgGroup> list(@Param("orgId") Integer orgId);

    void save(TstOrgGroup vo);

    void update(TstOrgGroup vo);

    TstOrgGroup get(Integer id);
}
