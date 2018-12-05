package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.CustomFieldOption;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CustomFieldOptionDao {
    List<CustomFieldOption> listByFieldId(@Param("fieldId") Integer fieldId);
    CustomFieldOption get(@Param("id") Integer id,
                          @Param("fieldId") Integer fieldId,
                          @Param("orgId") Integer orgId);

    void save(CustomFieldOption vo);
    void saveAll(@Param("fieldId") Integer fieldId, @Param("list") List<CustomFieldOption> options);

    void update(CustomFieldOption vo);
    void delete(@Param("id") Integer id);

    void setOrder(@Param("id") Integer id,
                  @Param("ordr") Integer ordr);

    Integer setDefault(@Param("id") Integer id,
                       @Param("fieldId") Integer fieldId);
    Integer removeDefault(@Param("fieldId") Integer fieldId);

    CustomFieldOption getPrev(@Param("ordr") Integer ordr,
                              @Param("fieldId") Integer fieldId);

    CustomFieldOption getNext(@Param("ordr") Integer ordr,
                              @Param("fieldId") Integer fieldId);

    Integer getMaxOrder(@Param("fieldId") Integer fieldId);
}
