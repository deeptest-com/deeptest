package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCustomFieldOption;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CustomFieldOptionDao {
    List<TstCustomFieldOption> listByFieldId(@Param("fieldId") Integer fieldId);
    Integer getMaxOrder(@Param("fieldId") Integer fieldId);

    void save(TstCustomFieldOption vo);
    void saveAll(@Param("fieldId") Integer fieldId, @Param("list") List<TstCustomFieldOption> options);

    void update(TstCustomFieldOption vo);

    void delete(@Param("id") Integer id);
}
