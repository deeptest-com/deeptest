package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuWorkflowSolution;
import com.ngtesting.platform.model.IsuWorkflowSolutionItem;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueWorkflowSolutionDao {

    List<IsuWorkflowSolution> list(Integer orgId);

    IsuWorkflowSolution get(@Param("id") Integer id, @Param("orgId") Integer orgId);

    void save(IsuWorkflowSolution vo);

    Integer update(IsuWorkflowSolution vo);

    Integer delete(@Param("id") Integer id, @Param("orgId") Integer orgId);

    List<IsuWorkflowSolutionItem> getItems(@Param("solutionId") Integer solutionId,
                                           @Param("orgId") Integer orgId);

    Integer changeItem(@Param("typeId") Integer typeId,
                       @Param("workflowId") Integer workflowId,
                       @Param("solutionId") Integer solutionId,
                       @Param("orgId") Integer orgId);
}
