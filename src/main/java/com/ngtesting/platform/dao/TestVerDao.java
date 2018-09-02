package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstVer;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestVerDao {
    List<TstVer> query(@Param("projectId") Integer projectId,
                       @Param("keywords") String keywords,
                       @Param("disabled") Boolean disabled);
    List<TstVer> listLastest(@Param("projectId") Integer projectId);

    TstVer get(@Param("id") Integer id,
               @Param("projectId") Integer projectId);

    Integer add(TstVer vo);
    Integer update(TstVer vo);

    Integer delete(@Param("id") Integer id,
                @Param("projectId") Integer projectId);
    Integer setOrder(@Param("id") Integer id,
                  @Param("ordr") Integer ordr,
                  @Param("projectId") Integer projectId);

    Integer getMaxOrdrNumb(@Param("projectId") Integer projectId);
    TstVer getPrev(@Param("ordr") Integer ordr,
                   @Param("projectId") Integer projectId);
    TstVer getNext(@Param("ordr") Integer ordr,
                   @Param("projectId") Integer projectId);
}
