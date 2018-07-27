package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstEnv;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestEnvDao {
    List<TstEnv> query(@Param("projectId") Integer projectId,
                       @Param("keywords") String keywords,
                       @Param("disabled") Boolean disabled);
    TstEnv get(@Param("id") Integer id);

    Integer getMaxOrdrNumb(@Param("projectId") Integer projectId);

    void add(TstEnv vo);
    void update(TstEnv vo);

    TstEnv getPrev(@Param("ordr") Integer ordr, @Param("projectId") Integer projectId);

    TstEnv getNext(@Param("ordr") Integer ordr, @Param("projectId") Integer projectId);

    void setOrder(@Param("id") Integer id, @Param("ordr") Integer ordr);

    void delete(@Param("id") Integer id);
}
