package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstModule;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestModuleDao {
    List<TstModule> query(@Param("projectId") Integer projectId,
                       @Param("keywords") String keywords,
                       @Param("disabled") Boolean disabled);
    List<TstModule> listLastest(@Param("projectId") Integer projectId);

    TstModule get(@Param("id") Integer id,
               @Param("projectId") Integer projectId);

    Integer save(TstModule vo);
    Integer update(TstModule vo);

    Integer delete(@Param("id") Integer id,
                   @Param("projectId") Integer projectId);
    Integer setOrder(@Param("id") Integer id,
                     @Param("ordr") Integer ordr,
                     @Param("projectId") Integer projectId);

    Integer getMaxOrdrNumb(@Param("projectId") Integer projectId);
    TstModule getPrev(@Param("ordr") Integer ordr,
                   @Param("projectId") Integer projectId);
    TstModule getNext(@Param("ordr") Integer ordr,
                   @Param("projectId") Integer projectId);
}
