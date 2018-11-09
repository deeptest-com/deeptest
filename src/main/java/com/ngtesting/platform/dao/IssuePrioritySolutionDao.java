package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPrioritySolution;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePrioritySolutionDao {
    List<IsuPrioritySolution> list(@Param("orgId") Integer orgId);

    IsuPrioritySolution get(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);

    Integer save(IsuPrioritySolution vo);

    Integer update(IsuPrioritySolution vo);

    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);
}
