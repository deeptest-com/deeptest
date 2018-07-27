package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstVer;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TestVerDao {
    List<TstVer> query(@Param("projectId") Integer projectId,
                       @Param("keywords") String keywords,
                       @Param("disabled") Boolean disabled);
    TstVer get(@Param("id") Integer id);

    Integer getMaxOrdrNumb(@Param("projectId") Integer projectId);

    void add(TstVer vo);
    void update(TstVer vo);

    TstVer getPrev(@Param("ordr") Integer ordr, @Param("projectId") Integer projectId);

    TstVer getNext(@Param("ordr") Integer ordr, @Param("projectId") Integer projectId);

    void setOrder(@Param("id") Integer id, @Param("ordr") Integer ordr);

    void delete(@Param("id") Integer id);
}
