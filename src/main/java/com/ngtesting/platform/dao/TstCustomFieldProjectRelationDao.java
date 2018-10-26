package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCustomFieldProjectRelation;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface TstCustomFieldProjectRelationDao {
    List<TstCustomFieldProjectRelation> query(@Param("orgId") Integer orgId,
                                              @Param("fieldId") Integer fieldId);

    void removeAllProjectsForField(@Param("orgId") Integer orgId,
                                @Param("fieldId") Integer fieldId);

    void saveRelations(@Param("list") List<TstCustomFieldProjectRelation> selectedList);
}
