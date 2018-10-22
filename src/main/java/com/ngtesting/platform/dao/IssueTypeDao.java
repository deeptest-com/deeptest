package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuType;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueTypeDao {
    List<IsuType> list(@Param("orgId") Integer orgId);

    IsuType get(@Param("id") Integer id,
                    @Param("orgId") Integer orgId);


    Integer save(IsuType vo);
    Integer update(IsuType vo);
    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    Integer removeDefault(@Param("orgId") Integer orgId);
    Integer setDefault(@Param("id") Integer id,
                       @Param("orgId") Integer orgId);
    Integer setOrder(@Param("id") Integer id,
                     @Param("ordr") Integer ordr,
                     @Param("orgId") Integer orgId);

    IsuType getPrev(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);

    IsuType getNext(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);
}
