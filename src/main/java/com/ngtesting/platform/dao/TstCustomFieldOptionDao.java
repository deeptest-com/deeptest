package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCustomFieldOption;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TstCustomFieldOptionDao {
    List<TstCustomFieldOption> listByFieldId(@Param("fieldId") Integer fieldId);
    TstCustomFieldOption get(@Param("id") Integer id);

    void save(TstCustomFieldOption vo);
    void saveAll(@Param("fieldId") Integer fieldId, @Param("list") List<TstCustomFieldOption> options);

    void update(TstCustomFieldOption vo);
    void delete(@Param("id") Integer id);

    void setOrder(@Param("id") Integer id,
                  @Param("ordr") Integer ordr);

    TstCustomFieldOption getPrev(@Param("ordr") Integer ordr,
                                 @Param("fieldId") Integer fieldId);

    TstCustomFieldOption getNext(@Param("ordr") Integer ordr,
                                 @Param("fieldId") Integer fieldId);

    Integer getMaxOrder(@Param("fieldId") Integer fieldId);

}
