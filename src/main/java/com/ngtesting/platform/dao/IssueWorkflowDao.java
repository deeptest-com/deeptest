package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuStatus;
import com.ngtesting.platform.model.IsuWorkflow;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueWorkflowDao {
    List<IsuWorkflow> list(@Param("orgId") Integer orgId);

    IsuWorkflow get(@Param("id") Integer id,
                @Param("orgId") Integer orgId);

    Integer save(IsuWorkflow vo);
    Integer update(IsuWorkflow vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    Integer removeDefault(@Param("orgId") Integer orgId);
    Integer setDefault(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);

    List<IsuStatus> listStatus(@Param("id") Integer id);

    void removeTransitions(@Param("workflowId") Integer workflowId,
                           @Param("statusIds") String statusIds,
                           @Param("orgId") Integer orgId);

    void removeStatuses(@Param("workflowId") Integer workflowId,
                        @Param("statusIds") String statusIds,
                        @Param("orgId") Integer orgId);

    void saveStatuses(@Param("workflowId") Integer workflowId,
                      @Param("statusIds") List<Integer> statusIds,
                      @Param("orgId") Integer orgId);
}
