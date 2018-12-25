package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.CustomFieldOption;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CustomFieldOptionDao {
    List<CustomFieldOption> listByFieldId(@Param("fieldId") Integer fieldId, @Param("orgId") Integer orgId);
    CustomFieldOption get(@Param("id") Integer id,
                          @Param("fieldId") Integer fieldId,
                          @Param("orgId") Integer orgId);

    void save(CustomFieldOption vo);
    Integer update(CustomFieldOption vo);
    Integer delete(@Param("id") Integer id,
                @Param("orgId") Integer orgId);

    void setOrder(@Param("id") Integer id,
                  @Param("ordr") Integer ordr);

    Integer setDefault(@Param("id") Integer id,
                       @Param("fieldId") Integer fieldId,
                       @Param("orgId") Integer orgId);
    Integer removeDefault(@Param("fieldId") Integer fieldId,
                          @Param("orgId") Integer orgId);

    CustomFieldOption getPrev(@Param("ordr") Integer ordr,
                              @Param("fieldId") Integer fieldId,
                              @Param("orgId") Integer orgId);

    CustomFieldOption getNext(@Param("ordr") Integer ordr,
                              @Param("fieldId") Integer fieldId,
                              @Param("orgId") Integer orgId);

    Integer getMaxOrder(@Param("fieldId") Integer fieldId,
                        @Param("orgId") Integer orgId);
}
