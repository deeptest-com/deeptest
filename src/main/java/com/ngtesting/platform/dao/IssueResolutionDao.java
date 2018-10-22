package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuResolution;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueResolutionDao {
    List<IsuResolution> list(@Param("orgId") Integer orgId);

    IsuResolution get(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);


    Integer save(IsuResolution vo);
    Integer update(IsuResolution vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    Integer removeDefault(@Param("orgId") Integer orgId);
    Integer setDefault(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);
    Integer setOrder(@Param("id") Integer id,
                     @Param("ordr") Integer ordr,
                     @Param("orgId") Integer orgId);

    IsuResolution getPrev(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);

    IsuResolution getNext(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);
}
