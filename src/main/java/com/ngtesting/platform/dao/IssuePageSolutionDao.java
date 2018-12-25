package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPageSolution;
import com.ngtesting.platform.model.IsuPageSolutionItem;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePageSolutionDao {

    List<IsuPageSolution> list(Integer orgId);

    IsuPageSolution get(@Param("id") Integer id,
                        @Param("orgId") Integer orgId);

    void save(IsuPageSolution vo);

    Integer update(IsuPageSolution vo);

    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    List<IsuPageSolutionItem> getItems(@Param("solutionId") Integer solutionId,
                                       @Param("orgId") Integer orgId);

    Integer changeItem(@Param("typeId") Integer typeId,
                        @Param("opt") String opt,
                        @Param("pageId") Integer pageId,
                        @Param("solutionId") Integer solutionId,
                        @Param("orgId") Integer orgId);

    Integer removeDefault(@Param("orgId") Integer orgId);
    Integer setDefault(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);

    // For Project
    IsuPageSolution getByProject(@Param("projectId") Integer projectId,
                                     @Param("orgId") Integer orgId);

    void setByProject(@Param("solutionId") Integer solutionId,
                      @Param("projectId") Integer projectId,
                      @Param("orgId") Integer orgId);
}
