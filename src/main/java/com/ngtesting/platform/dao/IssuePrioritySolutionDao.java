package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPrioritySolution;
import com.ngtesting.platform.model.IsuPriority;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePrioritySolutionDao {
    List<IsuPrioritySolution> list(@Param("orgId") Integer orgId);

    IsuPrioritySolution get(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);
    IsuPrioritySolution getDetail(@Param("id") Integer id,
                                  @Param("orgId") Integer orgId);

    Integer save(IsuPrioritySolution vo);

    Integer update(IsuPrioritySolution vo);

    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    Integer addPriority(@Param("priorityId") Integer priorityId,
                    @Param("solutionId") Integer solutionId,
                    @Param("orgId") Integer orgId);

    Integer removePriority(@Param("priorityId") Integer priorityId,
                       @Param("solutionId") Integer solutionId,
                       @Param("orgId") Integer orgId);

    Integer addAll(@Param("priorities") List<IsuPriority> prioritys,
                   @Param("solutionId") Integer solutionId,
                   @Param("orgId") Integer orgId);

    Integer removeAll(@Param("solutionId") Integer solutionId,
                      @Param("orgId") Integer orgId);
}
