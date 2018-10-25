package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuIssue;
import org.apache.ibatis.annotations.Param;

public interface IssueDao {
    IsuIssue get(@Param("id") Integer id, @Param("orgId") Integer orgId);
}
