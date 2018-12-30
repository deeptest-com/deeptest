package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseInTaskComments;
import org.apache.ibatis.annotations.Param;

public interface CaseInTaskCommentsDao {
    TstCaseInTaskComments get(@Param("id") Integer id);

    void save(TstCaseInTaskComments vo);

    void update(@Param("vo") TstCaseInTaskComments vo,
                @Param("userId") Integer userId);

    boolean delete(@Param("id") Integer id,
                   @Param("userId") Integer userId);
}
