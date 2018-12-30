package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseInTaskHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseInTaskHistoryDao {

    List<TstCaseInTaskHistory> query(@Param("caseInTaskId") Integer caseInTaskId);
    void save(TstCaseInTaskHistory his);
}
