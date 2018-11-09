package com.ngtesting.platform.dao;

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
}
