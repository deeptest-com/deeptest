package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueHistoryDao {
    List<IsuHistory> query(@Param("issueId") Integer issueId);
    void save(IsuHistory his);
}
