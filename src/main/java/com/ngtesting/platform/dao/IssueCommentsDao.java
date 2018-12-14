package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuComments;
import org.apache.ibatis.annotations.Param;

public interface IssueCommentsDao {
    IsuComments get(@Param("id") Integer id);

    void update(IsuComments vo);

    void save(IsuComments vo);

    boolean delete(@Param("id") Integer id, @Param("userId") Integer userId);
}
