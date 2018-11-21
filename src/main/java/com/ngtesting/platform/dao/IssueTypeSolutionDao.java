package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuType;
import com.ngtesting.platform.model.IsuTypeSolution;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueTypeSolutionDao {
    List<IsuTypeSolution> list(@Param("orgId") Integer orgId);

    IsuTypeSolution get(@Param("id") Integer id,
                        @Param("orgId") Integer orgId);


    Integer save(IsuTypeSolution vo);
    Integer update(IsuTypeSolution vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    IsuTypeSolution getDetail(@Param("id") Integer id,
                              @Param("orgId") Integer orgId);

    Integer addType(@Param("typeId") Integer typeId,
                    @Param("solutionId") Integer solutionId,
                    @Param("orgId") Integer orgId);

    Integer removeType(@Param("typeId") Integer typeId,
                       @Param("solutionId") Integer solutionId,
                       @Param("orgId") Integer orgId);

    Integer addAll(@Param("types") List<IsuType> types,
                   @Param("solutionId") Integer solutionId,
                   @Param("orgId") Integer orgId);

    Integer removeAll(@Param("solutionId") Integer solutionId,
                      @Param("orgId") Integer orgId);

    Integer removeDefault(@Param("orgId") Integer orgId);
    Integer setDefault(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);

    // For Project
    IsuTypeSolution getByProject(@Param("projectId") Integer projectId,
                                 @Param("orgId") Integer orgId);

    void setByProject(@Param("solutionId") Integer solutionId,
                      @Param("projectId") Integer projectId,
                      @Param("orgId") Integer orgId);
}
