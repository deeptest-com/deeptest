package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuPageTab;
import org.apache.ibatis.annotations.Param;

public interface IssuePageTabDao {

    void add(IsuPageTab tab);

    IsuPageTab get(@Param("tabId") Integer tabId, @Param("orgId") Integer orgId);

    Integer remove(@Param("id") Integer id, @Param("orgId") Integer orgId);
    Integer getMaxTabOrdr(Integer tabId);

}