package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCasePriority;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CasePriorityDao {
    List<TstCasePriority> list(@Param("orgId") Integer orgId);

    TstCasePriority get(@Param("id") Integer id);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);

    void save(TstCasePriority vo);

    void update(TstCasePriority vo);

    TstCasePriority getPrev(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    TstCasePriority getNext(@Param("ordr")Integer ordr, @Param("orgId") Integer orgId);

    void setOrder(@Param("id")Integer id, @Param("ordr") Integer ordr);

    void removeDefault(@Param("orgId") Integer orgId);

    void setDefault(@Param("id") Integer id, @Param("orgId") Integer orgId);

    void delete(@Param("id") Integer id);
}
