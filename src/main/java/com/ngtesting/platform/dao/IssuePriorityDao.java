package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPriority;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssuePriorityDao {
    List<IsuPriority> list(@Param("orgId") Integer orgId);

    IsuPriority get(@Param("id") Integer id,
                        @Param("orgId") Integer orgId);

    Integer save(IsuPriority vo);

    Integer update(IsuPriority vo);

    Integer delete(@Param("id") Integer id,
                   @Param("orgId") Integer orgId);

    Integer setOrder(@Param("id") Integer id,
                     @Param("ordr") Integer ordr,
                     @Param("orgId") Integer orgId);

    Integer setDefault(@Param("id") Integer id, @Param("orgId") Integer orgId);
    Integer removeDefault(@Param("orgId") Integer orgId);

    IsuPriority getPrev(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);
    IsuPriority getNext(@Param("ordr") Integer ordr, @Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);
}
