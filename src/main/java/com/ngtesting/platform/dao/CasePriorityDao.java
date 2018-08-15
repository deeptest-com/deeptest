package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCasePriority;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CasePriorityDao {
    List<TstCasePriority> list(@Param("orgId") Integer orgId);

    TstCasePriority get(@Param("id") Integer id,
                        @Param("orgId") Integer orgId);

    Integer save(TstCasePriority vo);

    Integer update(TstCasePriority vo);

    Integer delete(@Param("id") Integer id,
                @Param("orgId") Integer orgId);

    Integer setOrder(@Param("id") Integer id,
                     @Param("ordr") Integer ordr,
                     @Param("orgId") Integer orgId);

    Integer setDefault(@Param("id") Integer id, @Param("orgId") Integer orgId);
    Integer removeDefault(@Param("orgId") Integer orgId);

    TstCasePriority getPrev(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);
    TstCasePriority getNext(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);
}
