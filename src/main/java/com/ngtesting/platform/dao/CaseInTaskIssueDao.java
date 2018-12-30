package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseInTaskIssue;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseInTaskIssueDao {
    List<TstCaseInTaskIssue> query(@Param("caseInTaskId") Integer caseInTaskId);

    void save(TstCaseInTaskIssue entity);

    void delete(@Param("id") Integer id,
                @Param("userId") Integer userId);
}
