package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuWorkflowTransition;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueWorkflowTransitionDao {

    Integer save(IsuWorkflowTransition vo);
    Integer update(IsuWorkflowTransition vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    List<Integer> listProjectRoleId(@Param("id") Integer id,
                                    @Param("orgId") Integer orgId);
}
