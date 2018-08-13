package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrg;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgDao {
    List<TstOrg> query(@Param("userId") Integer userId,
                           @Param("keywords") String keywords,
                           @Param("disabled") Boolean disabled);

    List<TstOrg> queryByUser(@Param("userId") Integer userId);

    TstOrg get(Integer orgId);

    void setDefault(@Param("orgId") Integer orgId, @Param("userId") Integer userId);

    void initOrg(@Param("orgId") Integer id, @Param("userId") Integer userId);

    void save(TstOrg vo);
    void update(TstOrg vo);

    Integer delete(@Param("id") Integer id);
}
