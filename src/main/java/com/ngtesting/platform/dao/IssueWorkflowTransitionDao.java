package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuWorkflowTransition;
import org.apache.ibatis.annotations.Param;

public interface IssueWorkflowTransitionDao {

    Integer save(IsuWorkflowTransition vo);
    Integer update(IsuWorkflowTransition vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

}
