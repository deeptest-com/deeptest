package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCustomField;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CustomFieldDao {
    List list(@Param("orgId") Integer orgId);

    TstCustomField get(@Param("id") Integer id);

    List<String> getLastUnusedColumn(@Param("orgId") Integer orgId);

    Integer getMaxOrdrNumb(@Param("orgId") Integer orgId);

    void save(TstCustomField vo);

    void update(TstCustomField vo);
}
