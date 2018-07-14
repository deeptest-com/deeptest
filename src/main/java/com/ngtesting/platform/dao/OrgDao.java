package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstOrg;
import com.ngtesting.platform.model.TstProject;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface OrgDao {
    List<TstProject> query(@Param("keywords") String keywords, @Param("disabled") Boolean disabled);
    List<TstOrg> queryByUser(@Param("userId") Integer userId);

    void setDefault(@Param("orgId") Integer orgId, @Param("userId") Integer userId);
}
