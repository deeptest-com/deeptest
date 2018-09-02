package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstEnv;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestEnvDao {
    List<TstEnv> query(@Param("projectId") Integer projectId,
                       @Param("keywords") String keywords,
                       @Param("disabled") Boolean disabled);

    List<TstEnv> listLastest(@Param("projectId") Integer projectId);

    TstEnv get(@Param("id") Integer id,
               @Param("projectId") Integer projectId);

    Integer add(TstEnv vo);
    Integer update(TstEnv vo);

    Integer setOrder(@Param("id") Integer id,
                  @Param("ordr") Integer ordr,
                  @Param("projectId") Integer projectId);

    Integer delete(@Param("id") Integer id,
                @Param("projectId") Integer projectId);

    Integer getMaxOrdrNumb(@Param("projectId") Integer projectId);
    TstEnv getPrev(@Param("ordr") Integer ordr, @Param("projectId") Integer projectId);
    TstEnv getNext(@Param("ordr") Integer ordr, @Param("projectId") Integer projectId);
}
