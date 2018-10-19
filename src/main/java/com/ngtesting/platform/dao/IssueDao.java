package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuIssue;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueDao {
    List<IsuIssue> queryByProject(@Param("projectId") Integer projectId);
    List<IsuIssue> queryBySql(@Param("sql") String sql);
}
