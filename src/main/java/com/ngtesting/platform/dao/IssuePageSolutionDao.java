package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPageSolution;
import com.ngtesting.platform.model.IsuPageSolutionItem;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePageSolutionDao {

    List<IsuPageSolution> list(Integer orgId);

    IsuPageSolution get(@Param("id") Integer id, @Param("orgId") Integer orgId);

    List<IsuPageSolutionItem> getItems(@Param("pageSolutionId") Integer pageSolutionId,
                                       @Param("orgId") Integer orgId);

    void save(IsuPageSolution vo);

    Integer update(IsuPageSolution vo);

    Integer delete(@Param("id") Integer id, @Param("orgId") Integer orgId);

}
