package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.IsuHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface IssueHistoryDao {
    List<IsuHistory> query(@Param("caseId") Integer caseId);
    void save(IsuHistory his);
}
