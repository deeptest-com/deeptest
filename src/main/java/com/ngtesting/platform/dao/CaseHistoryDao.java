package com.ngtesting.platform.dao;

import com.ngtesting.platform.model.TstCaseHistory;
import org.apache.ibatis.annotations.Param;

import java.util.List;

public interface CaseHistoryDao {
    List<TstCaseHistory> query(@Param("caseId") Integer caseId);
    void save(TstCaseHistory his);
}
