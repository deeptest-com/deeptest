package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuWorkflowTransition;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueWorkflowTransitionDao {
    List<IsuWorkflowTransition> listTransition(@Param("workflowId") Integer id,
                                               @Param("projectRoleIds") List<Integer> projectRoleIds);

    List<IsuWorkflowTransition> listByStatus(@Param("projectId") Integer projectId,
                                             @Param("statusId") Integer statusId);

    IsuWorkflowTransition get(@Param("id") Integer id,
                              @Param("orgId") Integer orgId);

    IsuWorkflowTransition emptyObject(@Param("srcStatusId") Integer srcStatusId,
                                      @Param("dictStatusId") Integer dictStatusId,
                                      @Param("orgId") Integer orgId);

    Integer save(IsuWorkflowTransition vo);
    Integer update(IsuWorkflowTransition vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    List<Integer> listProjectRoleId(@Param("id") Integer id,
                                    @Param("orgId") Integer orgId);

    void removeAllRoles(@Param("id") Integer id,
                        @Param("orgId") Integer orgId);

    void addRoles(@Param("tran") IsuWorkflowTransition tran,
                  @Param("projectRoleIds") List<Integer> projectRoleIds);
}
